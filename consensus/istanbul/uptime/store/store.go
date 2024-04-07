package store

import (
	"github.com/Alexfordev/atlas/consensus/istanbul/uptime"
	"github.com/Alexfordev/atlas/core/rawdb"
	"github.com/ethereum/go-ethereum/ethdb"
)

type uptimeStoreImpl struct {
	db ethdb.Database
}

func New(db ethdb.Database) uptime.Store {
	return &uptimeStoreImpl{
		db: db,
	}
}

func (us *uptimeStoreImpl) ReadAccumulatedEpochUptime(epoch uint64) *uptime.Uptime {
	return rawdb.ReadAccumulatedEpochUptime(us.db, epoch)
}
func (us *uptimeStoreImpl) WriteAccumulatedEpochUptime(epoch uint64, uptime *uptime.Uptime) {
	rawdb.WriteAccumulatedEpochUptime(us.db, epoch, uptime)
}
