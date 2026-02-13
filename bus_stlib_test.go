//go:build !wasm

package bus

import "testing"

func TestBusStlib(t *testing.T) {
	testBus(t)
}
