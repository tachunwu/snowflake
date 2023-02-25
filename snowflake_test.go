package snowflake

import "testing"

func TestSnowflake(t *testing.T) {
	workerID := uint16(1)
	sf := NewSnowflake(workerID)

	var ids [100]int64
	var maxID int64 = 0
	for i := 0; i < 100; i++ {
		id := sf.NextID()
		ids[i] = int64(id)
		if id < uint64(maxID) {
			t.Errorf("ID %d is smaller than previous ID %d", id, maxID)
		}
		maxID = int64(id)
	}

	for i := 0; i < 99; i++ {
		if ids[i] == ids[i+1] {
			t.Errorf("ID %d and ID %d are the same", ids[i], ids[i+1])
		}
	}
}

func BenchmarkSnowflake(b *testing.B) {
	workerID := uint16(1)
	sf := NewSnowflake(workerID)

	for n := 0; n < b.N; n++ {
		sf.NextID()
	}
}
