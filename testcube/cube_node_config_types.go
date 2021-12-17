package main

import (
	"time"

	"github.com/reusee/fz"
)

type CubeCapacity int64

func (_ CubeNodeScope) CubeCapacity() (c CubeCapacity) {
	fz.RandBetween(&c, 1*1024*1024, 1*1024*1024*1024)
	return
}

type CubeShardGroups uint64

func (_ CubeNodeScope) CubeShardGroups() (c CubeShardGroups) {
	fz.RandBetween(&c, 1, 3)
	return
}

type CubeMaxPeerDownTime time.Duration

func (_ CubeNodeScope) CubeMaxPeerDownTime() (t CubeMaxPeerDownTime) {
	fz.RandBetween(&t, time.Minute, time.Minute*5)
	return
}

type CubeShardHeartbeatDuration time.Duration

func (_ CubeNodeScope) CubeShardHeartbeatDuration() (d CubeShardHeartbeatDuration) {
	fz.RandBetween(&d, time.Minute, time.Minute*5)
	return
}

type CubeStoreHeartbeatDuration time.Duration

func (_ CubeNodeScope) CubeStoreHeartbeatDuration() (d CubeStoreHeartbeatDuration) {
	fz.RandBetween(&d, time.Minute, time.Minute*5)
	return
}

type CubeShardSplitCheckDuration time.Duration

func (_ CubeNodeScope) CubeShardSplitCheckDuration() (d CubeShardSplitCheckDuration) {
	fz.RandBetween(&d, time.Minute, time.Minute*5)
	return
}

type CubeShardStateCheckDuration time.Duration

func (_ CubeNodeScope) CubeShardStateCheckDuration() (d CubeShardStateCheckDuration) {
	fz.RandBetween(&d, time.Minute, time.Minute*5)
	return
}

type CubeCompactLogCheckDuration time.Duration

func (_ CubeNodeScope) CubeCompactLogCheckDuration() (d CubeCompactLogCheckDuration) {
	fz.RandBetween(&d, time.Minute, time.Minute*5)
	return
}

type CubeDisableShardSplit bool

func (_ CubeNodeScope) CubeDisableShardSplit() (c CubeDisableShardSplit) {
	fz.RandBool(&c)
	return
}

type CubeAllowRemoveLeader bool

func (_ CubeNodeScope) CubeAllowRemoveLeader() (c CubeAllowRemoveLeader) {
	fz.RandBool(&c)
	return
}

type CubeShardCapacityBytes uint64

func (_ CubeNodeScope) CubeShardCapacityBytes() (c CubeShardCapacityBytes) {
	fz.RandBetween(&c, 4*1024*1024, 128*1024*1024)
	return
}

type CubeShardSplitCheckBytes uint64

func (_ CubeNodeScope) CubeShardSplitCheckBytes() (c CubeShardSplitCheckBytes) {
	fz.RandBetween(&c, 4*1024*1024, 128*1024*1024)
	return
}
