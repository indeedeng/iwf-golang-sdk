package iwf


type PersistenceLoadingType string

const (
	// LOAD_ALL_WITHOUT_LOCKING is the default policy, it will load all attributes without locking
	LOAD_ALL_WITHOUT_LOCKING PersistenceLoadingType = "LOAD_ALL_WITHOUT_LOCKING"

	// LOAD_ALL_WITH_EXCLUSIVE_LOCK will load all attributes but lock them exclusively for one request(execute/decide) at a time
	// this is to prevent racing condition of different states overwriting the attributes
	// Note that the lock will be released after execute/decide is completed, so it will allow commands to
	// execute in parallel.
	// LOAD_ALL_WITH_EXCLUSIVE_LOCK PersistenceLoadingType = "LOAD_ALL_WITH_EXCLUSIVE_LOCK"

	// LOAD_PARTIAL_WITHOUT_LOCKING is same as LOAD_ALL_WITHOUT_LOCKING but only load part of the attributes(need to specified in policy)
	// LOAD_PARTIAL_WITHOUT_LOCKING PersistenceLoadingType = "LOAD_PARTIAL_WITHOUT_LOCKING"

	// LOAD_PARTIAL_WITH_EXCLUSIVE_LOCK is same as LOAD_ALL_WITH_EXCLUSIVE_LOCK but only load part of the attributes(need to specified in policy)
	// LOAD_PARTIAL_WITH_EXCLUSIVE_LOCK PersistenceLoadingType = "LOAD_PARTIAL_WITH_EXCLUSIVE_LOCK"
)

type PersistenceLoadingPolicy struct {
	PersistenceLoadingType PersistenceLoadingType
	PartialLoadingKeys     []string
}

