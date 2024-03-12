package types

const (
	// ModuleName defines the module name
	ModuleName = "split"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_split"

	// poolprefix defines the prefix of the pool key usd for the storage of pool
	PoolPrefixKey = "splitter_pool/"
)

var (
	ParamsKey = []byte{0x01}
	// KeyNextPoolId defines key to store the next Pool ID to be used.
	KeyNextPoolId = []byte{0x02}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetPoolKey(poolId uint64) []byte {
	poolKeyStr := PoolPrefixKey + string(poolId)
	return []byte(poolKeyStr)
}
