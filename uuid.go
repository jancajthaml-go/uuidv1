// Information about the algorithm is available on Wikipedia
//
// https://en.wikipedia.org/wiki/Universally_unique_identifier
//
package uuid

import (
	"math/rand"
	"sync"
	"time"
	"encoding/hex"
	"encoding/binary"
	"fmt"
	"net"
)

const epochOffset = uint64(122192928000000000)

var (
	timeMu   sync.Mutex
	lastTime uint64
	clockSeq uint16
	nodeID   string
)

func getHardwareInterface() []byte {
	var err error
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, ifs := range interfaces {
		if len(ifs.HardwareAddr) >= 6 {
			return ifs.HardwareAddr
		}
	}
	return nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
	seed := make([]byte, 2)
	binary.LittleEndian.PutUint16(seed, uint16(rand.Uint32()))
	seq := int(seed[0])<<8 | int(seed[1])
	clockSeq = uint16(seq&0x3fff)
	lastTime = 0
	iface := getHardwareInterface()
	if len(iface) != 0 {
		nodeID = hex.EncodeToString(iface)
	} else {
		node := make([]byte, 8)
		binary.LittleEndian.PutUint64(node, rand.Uint64())
		nodeID = fmt.Sprintf("%012x", node[0:6:6])
	}
}

// Generate uuid version 1 variant DCE 1.1, ISO/IEC 11578:1996
func Generate() ([]byte, error) {
	defer timeMu.Unlock()
	timeMu.Lock()
	t := time.Now()
	now := uint64(t.UnixNano()/100) + epochOffset
	if now <= lastTime {
		clockSeq = ((clockSeq + 1) & 0x3fff)
	}
	lastTime = now
	r := fmt.Sprintf("%08x", now & 0xffffffff)
	r += "-"
	r += fmt.Sprintf("%04x", (now >> 32) & 0xffff)
	r += "-"
	r += fmt.Sprintf("%04x", ((now >> 48) & 0x0fff) | 0x1000)
	r += "-"
	r += fmt.Sprintf("%04x", clockSeq | 0x8000)
	r += "-"
	r += nodeID
	return []byte(r)[0:36:36], nil
}
