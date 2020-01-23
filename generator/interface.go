package generator

// Generator - common interface for id generators
type Generator interface {
	GetId() uint64
}
