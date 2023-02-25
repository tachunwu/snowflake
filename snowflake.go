package snowflake

import (
	"sync"
	"time"
)

const (
	epoch      int64 = 1577808000000 // 2020-01-01 00:00:00 UTC
	workerBits uint8 = 5             // Machine ID
	seqBits    uint8 = 12            // Serial Number
)

type Snowflake struct {
	mu        sync.Mutex
	timestamp int64  // now
	workerID  uint16 // Machine ID
	sequence  uint16
}

func NewSnowflake(workerID uint16) *Snowflake {
	return &Snowflake{
		timestamp: 0,
		workerID:  workerID,
		sequence:  0,
	}
}

func (sf *Snowflake) NextID() uint64 {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	now := time.Now().UnixNano() / 1000000 // 转换为毫秒
	if sf.timestamp == now {
		sf.sequence = (sf.sequence + 1) & (1<<seqBits - 1)
		if sf.sequence == 0 {
			now = sf.nextMillis(now)
		}
	} else {
		sf.sequence = 0
	}

	sf.timestamp = now
	return uint64((now-epoch)<<(workerBits+seqBits) | int64(sf.workerID)<<(seqBits) | int64(sf.sequence))
}

func (sf *Snowflake) nextMillis(last int64) int64 {
	now := time.Now().UnixNano() / 1000000
	for now <= last {
		now = time.Now().UnixNano() / 1000000
	}
	return now
}
