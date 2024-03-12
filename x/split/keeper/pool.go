package keeper

import (
	"context"

	"github.com/99adarsh/orbit/x/split/types"
	"github.com/cosmos/gogoproto/proto"
	gogotypes "github.com/cosmos/gogoproto/types"
)

// return the next pool id
func (k Keeper) GetNextPoolId(ctx context.Context) uint64 {
	store := k.storeService.OpenKVStore(ctx)
	b, err := store.Get(types.KeyNextPoolId)
	if err != nil {
		panic(err)
	}
	nextPoolId := gogotypes.UInt64Value{}
	err = proto.Unmarshal(b, &nextPoolId)
	if err != nil {
		panic(err)
	}
	return nextPoolId.Value
}

// store the next pool id
func (k Keeper) SetNextPoolId(ctx context.Context, poolId uint64) {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := proto.Marshal(&gogotypes.UInt64Value{Value: poolId})
	if err != nil {
		panic(err)
	}

	err = store.Set(types.KeyNextPoolId, bz)
	if err != nil {
		panic(err)
	}
}

// store the pool and return the id of the pool(map poolId -> Pool)
func (k Keeper) GetPool(ctx context.Context, poolId uint64) (types.Pool, bool) {
	store := k.storeService.OpenKVStore(ctx)
	poolStoreKey := types.GetPoolKey(poolId)

	poolBytes, found := store.Get(poolStoreKey)
	if found != nil {
		return types.Pool{}, false
	}
	pool := types.Pool{}
	err := proto.Unmarshal(poolBytes, &pool)
	if err != nil {
		panic(err)
	}

	return pool, true

}

func (k Keeper) SetPool(ctx context.Context, pool types.Pool) {
	store := k.storeService.OpenKVStore(ctx)
	poolStoreKey := types.GetPoolKey(pool.Id)

	bz, err := proto.Marshal(&pool)
	if err != nil {
		panic(err)
	}

	err = store.Set(poolStoreKey, bz)
	if err != nil {
		panic(err)
	}
}
