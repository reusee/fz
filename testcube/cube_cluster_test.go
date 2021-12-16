package main

import (
	"fmt"
	"net"
	"path/filepath"
	"testing"
	"time"

	crdbpebble "github.com/cockroachdb/pebble"
	prophetconfig "github.com/matrixorigin/matrixcube/components/prophet/config"
	"github.com/matrixorigin/matrixcube/components/prophet/util/typeutil"
	"github.com/matrixorigin/matrixcube/config"
	"github.com/matrixorigin/matrixcube/raftstore"
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

	NewCubeScope(
		NewConfigScope(
			NewGlobalScope(),
		),
	).Call(func(
		newCluster NewCubeCluster,
		closeCluster CloseCubeCluster,
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

		var configs []*config.Config
		for i := 0; i < numNodes; i++ {
			i := i

			logger, err := zap.NewDevelopment()
			ce(err)
			fs := vfs.Default

			configs = append(configs, &config.Config{
				Logger:     logger,
				FS:         fs,
				DataPath:   filepath.Join(string(tempDir), "foo"),
				RaftAddr:   net.JoinHostPort("127.0.0.1", nextPort()),
				ClientAddr: net.JoinHostPort("127.0.0.1", nextPort()),
				Labels: [][]string{
					{"c", fmt.Sprintf("%d", i)},
				},

				Prophet: prophetconfig.Config{
					Replication: prophetconfig.ReplicationConfig{
						MaxReplicas: 3,
					},
					Schedule: prophetconfig.ScheduleConfig{
						EnableJointConsensus: true,
					},
					Name:        fmt.Sprintf("node-%d", i),
					RPCAddr:     prophetRPCAddrs[i],
					StorageNode: true,
					EmbedEtcd: prophetconfig.EmbedEtcdConfig{
						Join: func() string {
							if i == 0 {
								return ""
							}
							return etcdPeerEndpoints[0]
						}(),
						TickInterval:     typeutil.NewDuration(time.Millisecond * 30),
						ElectionInterval: typeutil.NewDuration(time.Millisecond * 150),
						ClientUrls:       etcdClientEndpoints[i],
						PeerUrls:         etcdPeerEndpoints[i],
					},
					DisableStrictReconfigCheck: true,
				},

				Replication: config.ReplicationConfig{
					ShardHeartbeatDuration:  typeutil.NewDuration(time.Millisecond * 100),
					StoreHeartbeatDuration:  typeutil.NewDuration(time.Second),
					ShardSplitCheckDuration: typeutil.NewDuration(time.Millisecond * 100),
				},

				Raft: config.RaftConfig{
					TickInterval: typeutil.NewDuration(time.Millisecond * 100),
				},

				Worker: config.WorkerConfig{
					RaftEventWorkers:       1,
					ApplyWorkerCount:       1,
					SendRaftMsgWorkerCount: 1,
				},

				Storage: func() config.StorageConfig {
					kvStorage, err := pebble.NewStorage(fs.PathJoin(string(tempDir), "data"), logger, &crdbpebble.Options{})
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

				Test: config.TestConfig{},
			})

		}

		cluster, err := newCluster(configs)
		ce(err)
		defer closeCluster(cluster)

	})
}

func TestCubeTestCluster(t *testing.T) {
	defer he(nil, e4.TestingFatal(t))

	NewCubeScope(
		NewConfigScope(
			NewGlobalScope(),
		),
	).Call(func() {

		cluster := raftstore.NewTestClusterStore(t,
			raftstore.WithTestClusterNodeCount(3),
			raftstore.DiskTestCluster,
		)
		cluster.Start()
		defer cluster.Stop()

		cluster.WaitShardByCountPerNode(1, time.Second*10)
		cluster.WaitLeadersByCount(1, time.Second*10)
		cluster.CheckShardCount(1)

	})
}
