package main

import (
	"encoding/json"
	"fmt"
	"net"
	"path/filepath"
	"strings"
	"testing"
	"time"

	crdbpebble "github.com/cockroachdb/pebble"
	"github.com/matrixorigin/matrixcube/aware"
	prophetconfig "github.com/matrixorigin/matrixcube/components/prophet/config"
	"github.com/matrixorigin/matrixcube/components/prophet/util/typeutil"
	"github.com/matrixorigin/matrixcube/config"
	"github.com/matrixorigin/matrixcube/metric"
	"github.com/matrixorigin/matrixcube/pb/meta"
	"github.com/matrixorigin/matrixcube/storage"
	"github.com/matrixorigin/matrixcube/storage/executor/simple"
	"github.com/matrixorigin/matrixcube/storage/kv"
	"github.com/matrixorigin/matrixcube/storage/kv/pebble"
	"github.com/matrixorigin/matrixcube/vfs"
	"github.com/reusee/e4"
	"github.com/reusee/fz"
	"go.uber.org/zap"
)

func TestNewCubeCluster(t *testing.T) {
	defer he(nil, e4.TestingFatal(t))

	NewTestScope().Call(func(
		start StartCubeCluster,
		stop StopCubeCluster,
		cleanup fz.Cleanup,
		tempDir fz.TempDir,
	) {
		defer cleanup()

		port := int64(54321)
		nextPort := func() (ret string) {
			ret = fmt.Sprintf("%d", port)
			port++
			return
		}

		numNodes := 3

		var etcdClientEndpoints []string
		var etcdPeerEndpoints []string
		var prophetRPCAddrs []string
		for i := 0; i < numNodes; i++ {
			etcdClientEndpoints = append(etcdClientEndpoints, "http://"+net.JoinHostPort("localhost", nextPort()))
			etcdPeerEndpoints = append(etcdPeerEndpoints, "http://"+net.JoinHostPort("localhost", nextPort()))
			prophetRPCAddrs = append(prophetRPCAddrs, net.JoinHostPort("127.0.0.1", nextPort()))
		}

		leaderReady := make(chan struct{})

		var configs []*config.Config
		for i := 0; i < numNodes; i++ {
			i := i

			loggerConfigStr := `{
        "level": "fatal",
        "encoding": "json"
      }`
			var loggerConfig zap.Config
			ce(json.NewDecoder(strings.NewReader(loggerConfigStr)).Decode(&loggerConfig))
			logger, err := loggerConfig.Build()
			ce(err)
			defer logger.Sync()
			fs := vfs.Default

			configs = append(configs, &config.Config{
				RaftAddr:   net.JoinHostPort("127.0.0.1", nextPort()),
				ClientAddr: net.JoinHostPort("127.0.0.1", nextPort()),
				DataPath:   filepath.Join(string(tempDir), fmt.Sprintf("data-%d", i)),
				DeployPath: "",
				Version:    "42",
				GitHash:    "",
				Labels: [][]string{
					{"node", fmt.Sprintf("%d", i)},
				},
				Capacity:           1 * 1024 * 1024 * 1024,
				UseMemoryAsStorage: false,
				ShardGroups:        1,

				Replication: config.ReplicationConfig{
					MaxPeerDownTime:         typeutil.NewDuration(time.Minute * 30),
					ShardHeartbeatDuration:  typeutil.NewDuration(time.Second * 5),
					StoreHeartbeatDuration:  typeutil.NewDuration(time.Second * 5),
					ShardSplitCheckDuration: typeutil.NewDuration(time.Second * 5),
					ShardStateCheckDuration: typeutil.NewDuration(time.Second),
					CompactLogCheckDuration: typeutil.NewDuration(time.Second),
					DisableShardSplit:       false,
					AllowRemoveLeader:       false,
					ShardCapacityBytes:      128 * 1024 * 1024,
					ShardSplitCheckBytes:    128 * 1024 * 1024,
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
					DataDir:    filepath.Join(string(tempDir), fmt.Sprintf("prophet-%d", i)),
					RPCAddr:    prophetRPCAddrs[i],
					RPCTimeout: typeutil.NewDuration(time.Second * 32),

					StorageNode:  true,
					ExternalEtcd: []string{},
					EmbedEtcd: prophetconfig.EmbedEtcdConfig{
						Join: func() string {
							if i == 0 {
								return ""
							}
							return etcdPeerEndpoints[0]
						}(),
						ClientUrls:              etcdClientEndpoints[i],
						PeerUrls:                etcdPeerEndpoints[i],
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

				Storage: func() config.StorageConfig {
					kvStorage, err := pebble.NewStorage(
						fs.PathJoin(string(tempDir), fmt.Sprintf("storage-%d", i)),
						logger,
						&crdbpebble.Options{},
					)
					ce(err)
					base := kv.NewBaseStorage(kvStorage, fs)
					dataStorage := kv.NewKVDataStorage(base, simple.NewSimpleKVExecutor(kvStorage))
					return config.StorageConfig{
						DataStorageFactory: func(group uint64) storage.DataStorage {
							return dataStorage
						},
						ForeachDataStorageFunc: func(fn func(storage.DataStorage)) {
							fn(dataStorage)
						},
					}
				}(),

				Customize: config.CustomizeConfig{
					CustomShardStateAwareFactory: func() aware.ShardStateAware {
						return &cubeShardStateAware{
							becomeLeader: func(_ meta.Shard) {
								close(leaderReady)
							},
						}
					},
				},

				Logger: logger,

				Metric: metric.Cfg{
					Addr:     "",
					Interval: 0,
				},

				FS: fs,
			})

		}

		cluster, err := start(configs)
		ce(err)

		<-leaderReady

		defer stop(cluster)

	})
}

type cubeShardStateAware struct {
	created         func(meta.Shard)
	updated         func(meta.Shard)
	splited         func(meta.Shard)
	destroyed       func(meta.Shard)
	becomeLeader    func(meta.Shard)
	becomeFollower  func(meta.Shard)
	snapshotApplied func(meta.Shard)
}

var _ aware.ShardStateAware = new(cubeShardStateAware)

func (c *cubeShardStateAware) Created(shard meta.Shard) {
	if c.created != nil {
		c.created(shard)
	}
}

func (c *cubeShardStateAware) Updated(shard meta.Shard) {
	if c.updated != nil {
		c.updated(shard)
	}
}

func (c *cubeShardStateAware) Splited(shard meta.Shard) {
	if c.splited != nil {
		c.splited(shard)
	}
}

func (c *cubeShardStateAware) Destroyed(shard meta.Shard) {
	if c.destroyed != nil {
		c.destroyed(shard)
	}
}

func (c *cubeShardStateAware) BecomeLeader(shard meta.Shard) {
	if c.becomeLeader != nil {
		c.becomeLeader(shard)
	}
}

func (c *cubeShardStateAware) BecomeFollower(shard meta.Shard) {
	if c.becomeFollower != nil {
		c.becomeFollower(shard)
	}
}

func (c *cubeShardStateAware) SnapshotApplied(shard meta.Shard) {
	if c.snapshotApplied != nil {
		c.snapshotApplied(shard)
	}
}