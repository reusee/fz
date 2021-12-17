package main

import "github.com/matrixorigin/matrixcube/raftstore"

type CubeNodeRaftStore = raftstore.Store

func (_ CubeNodeScope) RaftStore(
	config CubeConfig,
) CubeNodeRaftStore {
	return raftstore.NewStore(config)
}
