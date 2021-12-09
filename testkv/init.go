package testkv

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

func init() {
	var seed int64
	ce(binary.Read(crand.Reader, binary.BigEndian, &seed))
	rand.Seed(seed)
}
