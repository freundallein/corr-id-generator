package generator

import (
	"time"
)

//ComplexGenerator - generate unique uint64 ids
type ComplexGenerator struct {
	machineId uint8
	counter   chan uint16
}

func startCounter(results chan uint16) {
	var i uint64
	i = 0
	for {
		results <- uint16(i % (1<<14 - 1))
		i++
	}
}

// New - generat
func New(id uint8) *ComplexGenerator {
	counter := make(chan uint16)
	go startCounter(counter)
	return &ComplexGenerator{machineId: id, counter: counter}
}

//GetId - return unique uint64 id
func (gen *ComplexGenerator) GetId() uint64 {
	var id uint64
	epoch := uint64(time.Now().UTC().UnixNano()) / 1000000
	id = epoch << (64 - 42)                      // 42 bits for epoch
	id |= uint64(gen.machineId) << (64 - 42 - 8) // 8 bit for machine id
	id |= uint64(<-gen.counter)                  // 14 bit for counter
	return id
}
