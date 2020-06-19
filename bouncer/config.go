package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/elastic/beats/v7/libbeat/paths"
	"github.com/elastic/beats/v7/libbeat/publisher/pipeline"
)

// flagsConfig is used for parsing all available CLI flags
type flagsConfig struct {
	ConfigFiles       []string
	StrictPermissions bool
	Path              pathSettings
}

type settings struct {
	Inputs   []*common.Config
	Path     pathSettings
	Logging  logp.Config
	Registry registrySettings // XXX: copied from filebeat
	Limits   limitsSettings
	Location string                 // time zone info
	Output   common.ConfigNamespace `config:"output"`
	Pipeline pipeline.Config        `config:",inline"`
}

type registrySettings struct {
	Path          string        `config:"path"`
	Permissions   os.FileMode   `config:"file_permissions"`
	FlushTimeout  time.Duration `config:"flush"`
	CleanInterval time.Duration `config:"cleanup_interval"`
	MigrateFile   string        `config:"migrate_file"`
}

// configure global resource limits to be shared with input managers
type limitsSettings struct {
	// heartbeat monitors scheduled concurrent active operations limit
	Monitors int64
}

// pathSettings mimics how paths are configured in Beats.
// NOTE: As path setup and config file reloading is interleaved and managed
//       between multiple packages in libbeat we 'duplicate' the behavior
//       to not rely too much on libbeat globals setup.
type pathSettings struct {
	Home   string
	Config string
	Data   string
	Logs   string
}

func (c *flagsConfig) parseArgs(args []string) error {
	basename := filepath.Base(args[0])
	defaultConfigFile := basename + ".yml"

	flags := flag.NewFlagSet(basename, flag.ContinueOnError)
	c.registerFlags(flags)
	err := flags.Parse(args[1:])
	if err != nil {
		return err
	}

	if len(c.ConfigFiles) == 0 {
		c.ConfigFiles = []string{defaultConfigFile}
	}
	return nil
}

func (c *flagsConfig) registerFlags(flags *flag.FlagSet) {
	common.StringArrVarFlag(flags, &c.ConfigFiles, "c", "configuration files")
	flags.BoolVar(&c.StrictPermissions, "strict.perms", true, "Strict permission checking on config files")
	c.Path.registerFlags(flags)
}

func (p *pathSettings) registerFlags(flags *flag.FlagSet) {
	flags.StringVar(&p.Config, "path.config", "", "Configurations directory to look for config files")
	flags.StringVar(&p.Home, "path.home", "", "Home path")
	flags.StringVar(&p.Data, "path.data", "", "Data path")
	flags.StringVar(&p.Logs, "path.logs", "", "Logs path")
}

func (s pathSettings) Unify(cwd string) pathSettings {
	if s.Home == "" {
		s.Home = cwd
	}
	if s.Config == "" {
		s.Config = s.Home
	}
	if s.Data == "" {
		s.Data = filepath.Join(s.Home, "data")
	}
	if s.Logs == "" {
		s.Logs = filepath.Join(s.Home, "logs")
	}

	return s
}

func initPaths(ps pathSettings) (pathSettings, error) {
	workingDir := ps.Home
	if workingDir == "" {
		var err error
		workingDir, err = os.Getwd()
		if err != nil {
			return ps, fmt.Errorf("Failed to read working directory: %w", err)
		}
	}

	ps = ps.Unify(workingDir)

	// configure libbeat globals to help inputs accessing them
	paths.InitPaths(&paths.Path{
		Home:   ps.Home,
		Config: ps.Config,
		Data:   ps.Data,
		Logs:   ps.Logs,
	})

	return ps, nil
}
