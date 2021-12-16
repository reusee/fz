package main

import (
	"sync"

	"github.com/matrixorigin/matrixcube/config"
	"github.com/matrixorigin/matrixcube/raftstore"
)

type CubeCluster struct {
	Nodes []*CubeClusterNode
}

type CubeClusterNode struct {
	RaftStore raftstore.Store
}

type StartCubeCluster func(
	nodeConfigs []*config.Config,
) (
	cluster *CubeCluster,
	err error,
)

func (_ CubeScope) StartCubeCluster() StartCubeCluster {
	return func(
		nodeConfigs []*config.Config,
	) (
		cluster *CubeCluster,
		err error,
	) {
		defer he(&err)

		cluster = new(CubeCluster)
		wg := new(sync.WaitGroup)

		for _, config := range nodeConfigs {
			store := raftstore.NewStore(config)
			wg.Add(1)
			go func() {
				defer wg.Done()
				store.Start()
			}()
			cluster.Nodes = append(cluster.Nodes, &CubeClusterNode{
				RaftStore: store,
			})
		}

		wg.Wait()

		return
	}
}

type StopCubeCluster func(*CubeCluster) error

func (_ CubeScope) StopCubeCluster() StopCubeCluster {
	return func(cluster *CubeCluster) (err error) {
		defer he(&err)

		for _, node := range cluster.Nodes {
			node.RaftStore.Stop()
		}

		return
	}
}
