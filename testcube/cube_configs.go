package main

import (
	"math/rand"
	"time"
)

type CubeCapacity int64

func (_ CubeScope) CubeCapacity() CubeCapacity {
	return 1 * 1024 * 1024 * 1024
}

type CubeShardGroups uint64

func (_ CubeScope) CubeShardGroups() CubeShardGroups {
	return CubeShardGroups(1 + rand.Intn(3))
}

type CubeMaxPeerDownTime time.Duration

func (_ CubeScope) CubeMaxPeerDownTime() CubeMaxPeerDownTime {
	return CubeMaxPeerDownTime(time.Minute*1 + time.Minute*time.Duration(rand.Intn(5)))
}

type CubeShardHeartbeatDuration time.Duration

func (_ CubeScope) CubeShardHeartbeatDuration() CubeShardHeartbeatDuration {
	return CubeShardHeartbeatDuration(time.Minute*1 + time.Minute*time.Duration(rand.Intn(5)))
}

type CubeStoreHeartbeatDuration time.Duration

func (_ CubeScope) CubeStoreHeartbeatDuration() CubeStoreHeartbeatDuration {
	return CubeStoreHeartbeatDuration(time.Minute*1 + time.Minute*time.Duration(rand.Intn(5)))
}

type CubeShardSplitCheckDuration time.Duration

func (_ CubeScope) CubeShardSplitCheckDuration() CubeShardSplitCheckDuration {
	return CubeShardSplitCheckDuration(time.Minute*1 + time.Minute*time.Duration(rand.Intn(5)))
}

type CubeShardStateCheckDuration time.Duration

func (_ CubeScope) CubeShardStateCheckDuration() CubeShardStateCheckDuration {
	return CubeShardStateCheckDuration(time.Minute*1 + time.Minute*time.Duration(rand.Intn(5)))
}

type CubeCompactLogCheckDuration time.Duration

func (_ CubeScope) CubeCompactLogCheckDuration() CubeCompactLogCheckDuration {
	return CubeCompactLogCheckDuration(time.Minute*1 + time.Minute*time.Duration(rand.Intn(5)))
}

type CubeDisableShardSplit bool

func (_ CubeScope) CubeDisableShardSplit() CubeDisableShardSplit {
	return rand.Intn(2) == 0
}

type CubeAllowRemoveLeader bool

func (_ CubeScope) CubeAllowRemoveLeader() CubeAllowRemoveLeader {
	return rand.Intn(2) == 0
}

type CubeShardCapacityBytes uint64

func (_ CubeScope) CubeShardCapacityBytes() CubeShardCapacityBytes {
	return 128 * 1024 * 1024
}

type CubeShardSplitCheckBytes uint64

func (_ CubeScope) CubeShardSplitCheckBytes() CubeShardSplitCheckBytes {
	return 128 * 1024 * 1024
}
