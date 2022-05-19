// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build mage
// +build mage

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/magefile/mage/mg"

	devtools "github.com/elastic/beats/v7/dev-tools/mage"
	"github.com/elastic/beats/v7/dev-tools/mage/target/build"
	filebeat "github.com/elastic/beats/v7/filebeat/scripts/mage"

	//mage:import
	"github.com/elastic/beats/v7/dev-tools/mage/target/common"
	//mage:import generate
	_ "github.com/elastic/beats/v7/filebeat/scripts/mage/generate"
	//mage:import
	_ "github.com/elastic/beats/v7/dev-tools/mage/target/integtest/docker"
	//mage:import
	_ "github.com/elastic/beats/v7/dev-tools/mage/target/unittest"
	//mage:import
	"github.com/elastic/beats/v7/dev-tools/mage/target/test"
)

func init() {
	common.RegisterCheckDeps(Update)
	test.RegisterDeps(IntegTest)

	devtools.BeatDescription = "Filebeat sends log files to Logstash or directly to Elasticsearch."
	devtools.BeatLicense = "Elastic License"
}

// Build builds the Beat binary.
func Build() error {
	return devtools.Build(devtools.DefaultBuildArgs())
}

// BuildSystemTestBinary builds a binary instrumented for use with Python system tests.
func BuildSystemTestBinary() error {
	return devtools.BuildSystemTestBinary()
}

// GolangCrossBuild builds the Beat binary inside of the golang-builder.
// Do not use directly, use crossBuild instead.
func GolangCrossBuild() error {
	return filebeat.GolangCrossBuild()
}

// CrossBuild cross-builds the beat for all target platforms.
func CrossBuild() error {
	return filebeat.CrossBuild()
}

// BuildGoDaemon builds the go-daemon binary (use crossBuildGoDaemon).
func BuildGoDaemon() error {
	return devtools.BuildGoDaemon()
}

// CrossBuildGoDaemon cross-builds the go-daemon binary using Docker.
func CrossBuildGoDaemon() error {
	return devtools.CrossBuildGoDaemon()
}

// AssembleDarwinUniversal merges the darwin/amd64 and darwin/arm64 into a single
// universal binary using `lipo`. It assumes the darwin/amd64 and darwin/arm64
// were built and only performs the merge.
func AssembleDarwinUniversal() error {
	return build.AssembleDarwinUniversal()
}

// Package packages the Beat for distribution.
// Use SNAPSHOT=true to build snapshots.
// Use PLATFORMS to control the target platforms.
// Use VERSION_QUALIFIER to control the version qualifier.
func Package() {
	start := time.Now()
	defer func() { fmt.Println("package ran for", time.Since(start)) }()

	if v, found := os.LookupEnv("AGENT_PACKAGING"); found && v != "" {
		devtools.UseElasticBeatXPackReducedPackaging()
	} else {
		devtools.UseElasticBeatXPackPackaging()
	}

	devtools.PackageKibanaDashboardsFromBuildDir()
	filebeat.CustomizePackaging()

	mg.Deps(Update)
	mg.Deps(CrossBuild, CrossBuildGoDaemon)
	mg.SerialDeps(devtools.Package, TestPackages)
}

// TestPackages tests the generated packages (i.e. file modes, owners, groups).
func TestPackages() error {
	return devtools.TestPackages()
}

// Fields generates the fields.yml file and a fields.go for each module, input,
// and processor.
func Fields() {
	mg.Deps(fieldsYML, moduleFieldsGo, inputFieldsGo, processorsFieldsGo)
}

func inputFieldsGo() error {
	return devtools.GenerateModuleFieldsGo("input")
}

func moduleFieldsGo() error {
	return devtools.GenerateModuleFieldsGo("module")
}

func processorsFieldsGo() error {
	return devtools.GenerateModuleFieldsGo("processors")
}

// fieldsYML generates a fields.yml based on filebeat + x-pack/filebeat/modules.
func fieldsYML() error {
	return devtools.GenerateFieldsYAML(devtools.OSSBeatDir("module"), "module", "input", "processors")
}

// Dashboards collects all the dashboards and generates index patterns.
func Dashboards() error {
	return devtools.KibanaDashboards(devtools.OSSBeatDir("module"), "module", "input")
}

// ExportDashboard exports a dashboard and writes it into the correct directory.
//
// Required environment variables:
// - MODULE: Name of the module
// - ID:     Dashboard id
func ExportDashboard() error {
	return devtools.ExportDashboard()
}

// Config generates both the short and reference configs.
func Config() {
	mg.Deps(configYML, devtools.GenerateDirModulesD)
	mg.SerialDeps(devtools.ValidateDirModulesD, devtools.ValidateDirModulesDDatasetsDisabled)
}

func configYML() error {
	return devtools.Config(devtools.AllConfigTypes, filebeat.XPackConfigFileParams(), ".")
}

// Update is an alias for executing fields, dashboards, config.
func Update() {
	mg.SerialDeps(Fields, Dashboards, Config, includeList,
		filebeat.PrepareModulePackagingXPack)
}

func includeList() error {
	options := devtools.DefaultIncludeListOptions()
	options.ImportDirs = []string{"input/*", "processors/*"}
	return devtools.GenerateIncludeListGo(options)
}

// IntegTest executes integration tests (it uses Docker to run the tests).
func IntegTest() {
	mg.SerialDeps(GoIntegTest, PythonIntegTest)
}

// GoIntegTest starts the docker containers and executes the Go integration tests.
func GoIntegTest(ctx context.Context) error {
	return devtools.GoIntegTestFromHost(ctx, devtools.DefaultGoTestIntegrationFromHostArgs())
}

// PythonIntegTest starts the docker containers and executes the Python integration tests.
func PythonIntegTest(ctx context.Context) error {
	mg.Deps(Fields, Dashboards, devtools.BuildSystemTestBinary)
	return devtools.PythonIntegTestFromHost(devtools.DefaultPythonTestIntegrationFromHostArgs())
}
