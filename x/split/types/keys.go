package types

const (
	// ModuleName defines the module name
	ModuleName = "split"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_split"
)

var (
	ParamsKey = []byte("p_split")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
