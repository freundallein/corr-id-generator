package generator

import (
	"reflect"
	"testing"
)

func TestStartCounter(t *testing.T) {

}

func TestNew(t *testing.T) {
	observed := New(uint8(1))
	observedType := reflect.TypeOf(observed)
	expectedType := reflect.TypeOf(&ComplexGenerator{})
	if observedType != expectedType {
		t.Error("Expected", expectedType, "got", observedType)
	}
	if observed.machineId != 1 {
		t.Error("Expected", 1, "got", observed.machineId)
	}
	expectedChan := make(chan uint16)
	if reflect.TypeOf(observed.counter) != reflect.TypeOf(expectedChan) {
		t.Error("Expected", expectedChan, "got", reflect.TypeOf(observed.counter))
	}
}

func TestGetId(t *testing.T) {
	N := 100000
	gen := New(uint8(1))
	curr := gen.GetId()
	for i := 0; i < N; i++ {
		val := gen.GetId()
		if val == curr {
			t.Error("id is not unique")
		}
		curr = val
	}

}
