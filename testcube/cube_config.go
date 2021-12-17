package main

import (
	"fmt"
	"time"

	prophetconfig "github.com/matrixorigin/matrixcube/components/prophet/config"
	"github.com/matrixorigin/matrixcube/components/prophet/util/typeutil"
	"github.com/matrixorigin/matrixcube/config"
	"github.com/matrixorigin/matrixcube/metric"
)

type NewCubeConfig = func(num int) *config.Config

func (_ CubeScope) NewCubeConfig(
	capacity CubeCapacity,
	shardGroups CubeShardGroups,
	maxPeerDownTime CubeMaxPeerDownTime,
	shardHeartbeatDuration CubeShardHeartbeatDuration,
	storeHeartbeatDuration CubeStoreHeartbeatDuration,
	shardSplitCheckDuration CubeShardSplitCheckDuration,
	shardStateCheckDuration CubeShardStateCheckDuration,
	compactLogCheckDuration CubeCompactLogCheckDuration,
	disableShardSplit CubeDisableShardSplit,
	allowRemoveLeader CubeAllowRemoveLeader,
	shardCapacityBytes CubeShardCapacityBytes,
	shardSplitCheckBytes CubeShardSplitCheckBytes,
) NewCubeConfig {

	return func(i int) *config.Config {
		return &config.Config{
			DeployPath: "",
			Version:    "42",
			GitHash:    "",
			Labels: [][]string{
				{"node", fmt.Sprintf("%d", i)},
			},
			Capacity:           typeutil.ByteSize(capacity),
			UseMemoryAsStorage: false,
			ShardGroups:        uint64(shardGroups),

			Replication: config.ReplicationConfig{
				MaxPeerDownTime:         typeutil.NewDuration(time.Duration(maxPeerDownTime)),
				ShardHeartbeatDuration:  typeutil.NewDuration(time.Duration(shardHeartbeatDuration)),
				StoreHeartbeatDuration:  typeutil.NewDuration(time.Duration(storeHeartbeatDuration)),
				ShardSplitCheckDuration: typeutil.NewDuration(time.Duration(shardSplitCheckDuration)),
				ShardStateCheckDuration: typeutil.NewDuration(time.Duration(shardStateCheckDuration)),
				CompactLogCheckDuration: typeutil.NewDuration(time.Duration(compactLogCheckDuration)),
				DisableShardSplit:       bool(disableShardSplit),
				AllowRemoveLeader:       bool(allowRemoveLeader),
				ShardCapacityBytes:      typeutil.ByteSize(shardCapacityBytes),
				ShardSplitCheckBytes:    typeutil.ByteSize(shardSplitCheckBytes),
			},

			Raft: config.RaftConfig{
				TickInterval:         typeutil.NewDuration(time.Millisecond * 100),
				HeartbeatTicks:       10,
				ElectionTimeoutTicks: 50,
				MaxSizePerMsg:        8 * 1024 * 1024,
				MaxInflightMsgs:      256,
				MaxEntryBytes:        1 * 1024 * 1024,
				SendRaftBatchSize:    128,

				RaftLog: config.RaftLogConfig{
					DisableSync:         false,
					CompactThreshold:    256,
					MaxAllowTransferLag: 4,
					ForceCompactCount:   2048,
					ForceCompactBytes:   128 * 1024 * 1024,
				},
			},

			Worker: config.WorkerConfig{
				RaftEventWorkers:       128,
				ApplyWorkerCount:       128,
				SendRaftMsgWorkerCount: 128,
			},

			Prophet: prophetconfig.Config{
				Name:       fmt.Sprintf("prophet-%d", i),
				RPCTimeout: typeutil.NewDuration(time.Second * 32),

				StorageNode:  true,
				ExternalEtcd: []string{},
				EmbedEtcd: prophetconfig.EmbedEtcdConfig{
					AdvertiseClientUrls:     "",
					AdvertisePeerUrls:       "",
					InitialCluster:          "",
					InitialClusterState:     "",
					TickInterval:            typeutil.NewDuration(time.Millisecond * 30),
					ElectionInterval:        typeutil.NewDuration(time.Millisecond * 150),
					PreVote:                 true,
					AutoCompactionMode:      "periodic",
					AutoCompactionRetention: "1h",
					QuotaBackendBytes:       1 * 1024 * 1024 * 1024,
				},

				LeaderLease: 8,

				Schedule: prophetconfig.ScheduleConfig{
					MaxSnapshotCount:              3,
					MaxPendingPeerCount:           16,
					MaxMergeResourceSize:          128 * 1024 * 1024,
					MaxMergeResourceKeys:          16,
					SplitMergeInterval:            typeutil.NewDuration(time.Minute),
					EnableOneWayMerge:             true,
					EnableCrossTableMerge:         true,
					PatrolResourceInterval:        typeutil.NewDuration(time.Minute),
					MaxContainerDownTime:          typeutil.NewDuration(time.Minute),
					LeaderScheduleLimit:           4,
					LeaderSchedulePolicy:          "count",
					ResourceScheduleLimit:         2048,
					ReplicaScheduleLimit:          64,
					MergeScheduleLimit:            128,
					HotResourceScheduleLimit:      128,
					HotResourceCacheHitsThreshold: 128,
					TolerantSizeRatio:             0.8,
					LowSpaceRatio:                 0.8,
					HighSpaceRatio:                0.2,
					EnableJointConsensus:          true,
					// ... TODO

				},

				Replication: prophetconfig.ReplicationConfig{
					MaxReplicas:          3,
					EnablePlacementRules: true,
				},
			},

			Metric: metric.Cfg{
				Addr:     "",
				Interval: 0,
			},
		}
	}
}
