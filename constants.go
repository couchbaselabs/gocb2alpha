package gocb

import (
	"github.com/couchbaselabs/gocbcore8alpha"
)

const (
	goCbVersionStr = "v2.0.0-dev"

	persistenceTimeoutFloor = 1500
)

// IndexType provides information on the type of indexer used for an index.
type IndexType string

const (
	// IndexTypeN1ql indicates that GSI was used to build the index.
	IndexTypeN1ql = IndexType("gsi")

	// IndexTypeView indicates that views were used to build the index.
	IndexTypeView = IndexType("views")
)

// ServiceType specifies a particular Couchbase service type.
type ServiceType gocbcore.ServiceType

const (
	// MemdService represents a memcached service.
	MemdService = ServiceType(gocbcore.MemdService)

	// MgmtService represents a management service (typically ns_server).
	MgmtService = ServiceType(gocbcore.MgmtService)

	// CapiService represents a CouchAPI service (typically for views).
	CapiService = ServiceType(gocbcore.CapiService)

	// N1qlService represents a N1QL service (typically for query).
	N1qlService = ServiceType(gocbcore.N1qlService)

	// FtsService represents a full-text-search service.
	FtsService = ServiceType(gocbcore.FtsService)

	// CbasService represents an analytics service.
	CbasService = ServiceType(gocbcore.CbasService)
)

// QueryProfileType specifies the profiling mode to use during a query.
type QueryProfileType string

const (
	// QueryProfileNone disables query profiling
	QueryProfileNone = QueryProfileType("off")

	// QueryProfilePhases includes phase profiling information in the query response
	QueryProfilePhases = QueryProfileType("phases")

	// QueryProfileTimings includes timing profiling information in the query response
	QueryProfileTimings = QueryProfileType("timings")
)

// SubdocFlag provides special handling flags for sub-document operations
type SubdocFlag gocbcore.SubdocFlag

const (
	// SubdocFlagNone indicates no special behaviours
	SubdocFlagNone = SubdocFlag(gocbcore.SubdocFlagNone)

	// SubdocFlagCreatePath indicates you wish to recursively create the tree of paths
	// if it does not already exist within the document.
	SubdocFlagCreatePath = SubdocFlag(gocbcore.SubdocFlagMkDirP)

	// SubdocFlagXattr indicates your path refers to an extended attribute rather than the document.
	SubdocFlagXattr = SubdocFlag(gocbcore.SubdocFlagXattrPath)

	// SubdocFlagUseMacros indicates that you wish macro substitution to occur on the value
	SubdocFlagUseMacros = SubdocFlag(gocbcore.SubdocFlagExpandMacros)
)

// SubdocDocFlag specifies document-level flags for a sub-document operation.
type SubdocDocFlag gocbcore.SubdocDocFlag

const (
	// SubdocDocFlagNone indicates no special behaviours
	SubdocDocFlagNone = SubdocDocFlag(gocbcore.SubdocDocFlagNone)

	// SubdocDocFlagMkDoc indicates that the document should be created if it does not already exist.
	SubdocDocFlagMkDoc = SubdocDocFlag(gocbcore.SubdocDocFlagMkDoc)

	// SubdocDocFlagReplaceDoc indices that this operation should be a replace rather than upsert.
	SubdocDocFlagReplaceDoc = SubdocDocFlag(gocbcore.SubdocDocFlagReplaceDoc)

	// SubdocDocFlagAccessDeleted indicates that you wish to receive soft-deleted documents.
	SubdocDocFlagAccessDeleted = SubdocDocFlag(gocbcore.SubdocDocFlagAccessDeleted)
)

// DurabilityLevel specifies the level of synchronous replication to use.
type DurabilityLevel uint8

const (
	// DurabilityLevelMajority specifies that a mutation must be replicated (held in memory) to a majority of nodes.
	DurabilityLevelMajority = DurabilityLevel(1)

	// DurabilityLevelMajorityAndPersistActive specifies that a mutation must be replicated (held in memory) to a
	// majority of nodes and also persisted (written to disk) on the active node.
	DurabilityLevelMajorityAndPersistActive = DurabilityLevel(2)

	// DurabilityLevelPersistToMajority specifies that a mutation must be persisted (written to disk) to a majority
	// of nodes.
	DurabilityLevelPersistToMajority = DurabilityLevel(3)
)
