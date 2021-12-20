package main

import (
	"time"

	"github.com/matrixorigin/matrixcube/config"
	"github.com/reusee/fz"
)

type RandomizeCubeConfig func(config *config.Config)

func (_ Def) RandomizeCubeConfig() RandomizeCubeConfig {
	return func(config *config.Config) {

		fz.RandBetween(&config.Capacity, 1*1024*1024, 1*1024*1024*1024)
		fz.RandBetween(&config.ShardGroups, 3, 7)
		fz.RandBetween(&config.Replication.MaxPeerDownTime.Duration,
			time.Minute, time.Minute*5)
		fz.RandBetween(&config.Replication.ShardHeartbeatDuration.Duration,
			time.Minute, time.Minute*5)
		fz.RandBetween(&config.Replication.StoreHeartbeatDuration.Duration,
			time.Minute, time.Minute*5)
		fz.RandBetween(&config.Replication.ShardSplitCheckDuration.Duration,
			time.Minute, time.Minute*5)
		fz.RandBetween(&config.Replication.ShardStateCheckDuration.Duration,
			time.Minute, time.Minute*5)
		fz.RandBetween(&config.Replication.CompactLogCheckDuration.Duration,
			time.Minute, time.Minute*5)
		fz.RandBool(&config.Replication.DisableShardSplit)
		fz.RandBool(&config.Replication.AllowRemoveLeader)
		fz.RandBetween(&config.Replication.ShardCapacityBytes, 4*1024*1024, 128*1024*1024)
		fz.RandBetween(&config.Replication.ShardSplitCheckBytes, 4*1024*1024, 128*1024*1024)

	}
}
