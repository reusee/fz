package main

import "github.com/matrixorigin/matrixcube/server"

type CubeNodeApplication = *server.Application

func (_ CubeNodeScope) App(
	store CubeNodeRaftStore,
) CubeNodeApplication {
	return server.NewApplication(server.Cfg{
		Store: store,
	})
}
