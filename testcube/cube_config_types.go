package main

import (
	"time"

	"github.com/reusee/fz"
)

type CubeCapacity int64

func (_ CubeScope) CubeCapacity() (c CubeCapacity) {
	fz.RandBetween(&c, 1*1024*1024, 1*1024*1024*1024)
	return
}

type CubeShardGroups uint64

func (_ CubeScope) CubeShardGroups() (c CubeShardGroups) {
	fz.RandBetween(&c, 1, 3)
	return
}

type CubeMaxPeerDownTime time.Duration

func (_ CubeScope) CubeMaxPeerDownTime() (t CubeMaxPeerDownTime) {
	fz.RandBetween(&t, time.Minute, time.Minute*5)
	return
}

type CubeShardHeartbeatDuration time.Duration

func (_ CubeScope) CubeShardHeartbeatDuration() (d CubeShardHeartbeatDuration) {
	fz.RandBetween(&d, time.Minute, time.Minute*5)
	return
}

type CubeStoreHeartbeatDuration time.Duration

func (_ CubeScope) CubeStoreHeartbeatDuration() (d CubeStoreHeartbeatDuration) {
	fz.RandBetween(&d, time.Minute, time.Minute*5)
	return
}

type CubeShardSplitCheckDuration time.Duration

func (_ CubeScope) CubeShardSplitCheckDuration() (d CubeShardSplitCheckDuration) {
	fz.RandBetween(&d, time.Minute, time.Minute*5)
	return
}

type CubeShardStateCheckDuration time.Duration

func (_ CubeScope) CubeShardStateCheckDuration() (d CubeShardStateCheckDuration) {
	fz.RandBetween(&d, time.Minute, time.Minute*5)
	return
}

type CubeCompactLogCheckDuration time.Duration

func (_ CubeScope) CubeCompactLogCheckDuration() (d CubeCompactLogCheckDuration) {
	fz.RandBetween(&d, time.Minute, time.Minute*5)
	return
}

type CubeDisableShardSplit bool

func (_ CubeScope) CubeDisableShardSplit() (c CubeDisableShardSplit) {
	fz.RandBool(&c)
	return
}

type CubeAllowRemoveLeader bool

func (_ CubeScope) CubeAllowRemoveLeader() (c CubeAllowRemoveLeader) {
	fz.RandBool(&c)
	return
}

type CubeShardCapacityBytes uint64

func (_ CubeScope) CubeShardCapacityBytes() (c CubeShardCapacityBytes) {
	fz.RandBetween(&c, 4*1024*1024, 128*1024*1024)
	return
}

type CubeShardSplitCheckBytes uint64

func (_ CubeScope) CubeShardSplitCheckBytes() (c CubeShardSplitCheckBytes) {
	fz.RandBetween(&c, 4*1024*1024, 128*1024*1024)
	return
}
