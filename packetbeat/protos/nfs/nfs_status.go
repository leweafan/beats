// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package nfs

var nfsStatus = map[int]string{
	0:     "NFS_OK",
	1:     "NFSERR_PERM",
	2:     "NFSERR_NOENT",
	5:     "NFSERR_IO",
	6:     "NFSERR_NXIO",
	13:    "NFSERR_ACCESS",
	17:    "NFSERR_EXIST",
	18:    "NFSERR_XDEV",
	19:    "NFSERR_NODEV",
	20:    "NFSERR_NOTDIR",
	21:    "NFSERR_ISDIR",
	22:    "NFSERR_INVAL",
	27:    "NFSERR_FBIG",
	28:    "NFSERR_NOSPC",
	30:    "NFSERR_ROFS",
	31:    "NFSERR_MLINK",
	63:    "NFSERR_NAMETOOLONG",
	66:    "NFSERR_NOTEMPTY",
	69:    "NFSERR_DQUOT",
	70:    "NFSERR_STALE",
	71:    "NFSERR_REMOTE",
	99:    "NFSERR_WFLUSH",
	10001: "NFSERR_BADHANDLE",
	10002: "NFSERR_NOT_SYNC",
	10003: "NFSERR_BAD_COOKIE",
	10004: "NFSERR_NOTSUPP",
	10005: "NFSERR_TOOSMALL",
	10006: "NFSERR_SERVERFAULT",
	10007: "NFSERR_BADTYPE",
	10008: "NFSERR_DELAY",
	10009: "NFSERR_SAME",
	10010: "NFSERR_DENIED",
	10011: "NFSERR_EXPIRED",
	10012: "NFSERR_LOCKED",
	10013: "NFSERR_GRACE",
	10014: "NFSERR_FHEXPIRED",
	10015: "NFSERR_SHARE_DENIED",
	10016: "NFSERR_WRONGSEC",
	10017: "NFSERR_CLID_INUSE",
	10018: "NFSERR_RESOURCE",
	10019: "NFSERR_MOVED",
	10020: "NFSERR_NOFILEHANDLE",
	10021: "NFSERR_MINOR_VERS_MISMATCH",
	10022: "NFSERR_STALE_CLIENTID",
	10023: "NFSERR_STALE_STATEID",
	10024: "NFSERR_OLD_STATEID",
	10025: "NFSERR_BAD_STATEID",
	10026: "NFSERR_BAD_SEQID",
	10027: "NFSERR_NOT_SAME",
	10028: "NFSERR_LOCK_RANGE",
	10029: "NFSERR_SYMLINK",
	10030: "NFSERR_RESTOREFH",
	10031: "NFSERR_LEASE_MOVED",
	10032: "NFSERR_ATTRNOTSUPP",
	10033: "NFSERR_NO_GRACE",
	10034: "NFSERR_RECLAIM_BAD",
	10035: "NFSERR_RECLAIM_CONFLICT",
	10036: "NFSERR_BADXDR",
	10037: "NFSERR_LOCKS_HELD",
	10038: "NFSERR_OPENMODE",
	10039: "NFSERR_BADOWNER",
	10040: "NFSERR_BADCHAR",
	10041: "NFSERR_BADNAME",
	10042: "NFSERR_BAD_RANGE",
	10043: "NFSERR_LOCK_NOTSUPP",
	10044: "NFSERR_OP_ILLEGAL",
	10045: "NFSERR_DEADLOCK",
	10046: "NFSERR_FILE_OPEN",
	10047: "NFSERR_ADMIN_REVOKED",
	10048: "NFSERR_CB_PATH_DOWN",
	10049: "NFSERR_BADIOMODE",
	10050: "NFSERR_BADLAYOUT",
	10051: "NFSERR_BAD_SESSION_DIGEST",
	10052: "NFSERR_BADSESSION",
	10053: "NFSERR_BADSLOT",
	10054: "NFSERR_COMPLETE_ALREADY",
	10055: "NFSERR_CONN_NOT_BOUND_TO_SESSION",
	10056: "NFSERR_DELEG_ALREADY_WANTED",
	10057: "NFSERR_BACK_CHAN_BUSY",
	10058: "NFSERR_LAYOUTTRYLATER",
	10059: "NFSERR_LAYOUTUNAVAILABLE",
	10060: "NFSERR_NOMATCHING_LAYOUT",
	10061: "NFSERR_RECALLCONFLICT",
	10062: "NFSERR_UNKNOWN_LAYOUTTYPE",
	10063: "NFSERR_SEQ_MISORDERED",
	10064: "NFSERR_SEQUENCE_POS",
	10065: "NFSERR_REQ_TOO_BIG",
	10066: "NFSERR_REP_TOO_BIG",
	10067: "NFSERR_REP_TOO_BIG_TO_CACHE",
	10068: "NFSERR_RETRY_UNCACHED_REP",
	10069: "NFSERR_UNSAFE_COMPOUND",
	10070: "NFSERR_TOO_MANY_OPS",
	10071: "NFSERR_OP_NOT_IN_SESSION",
	10072: "NFSERR_HASH_ALG_UNSUPP",
	10073: "NFSERR_CONN_BINDING_NOT_ENFORCED",
	10074: "NFSERR_CLIENTID_BUSY",
	10075: "NFSERR_PNFS_IO_HOLE",
	10076: "NFSERR_SEQ_FALSE_RETRY",
	10077: "NFSERR_BAD_HIGH_SLOT",
	10078: "NFSERR_DEADSESSION",
	10079: "NFSERR_ENCR_ALG_UNSUPP",
	10080: "NFSERR_PNFS_NO_LAYOUT",
	10081: "NFSERR_NOT_ONLY_OP",
	10082: "NFSERR_WRONG_CRED",
	10083: "NFSERR_WRONG_TYPE",
	10084: "NFSERR_DIRDELEG_UNAVAIL",
	10085: "NFSERR_REJECT_DELEG",
	10086: "NFSERR_RETURNCONFLICT",
}
