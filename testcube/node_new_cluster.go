package main

import (
	"github.com/matrixorigin/matrixcube/config"
	"github.com/matrixorigin/matrixcube/raftstore"
)

type NewCubeCluster func(
	nodeConfigs []*config.Config,
) (
	stores []raftstore.Store,
	err error,
)

func (_ CubeScope) NewCubeCluster() NewCubeCluster {
	return func(
		nodeConfigs []*config.Config,
	) (
		stores []raftstore.Store,
		err error,
	) {
		defer he(&err)

		return
	}
}
