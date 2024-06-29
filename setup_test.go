package wasm

import (
	"testing"

	"github.com/coredns/caddy"
)

func TestFileParse(t *testing.T) {
	c := caddy.NewTestController("dns", `wasm`)
	if err := setup(c); err != nil {
		t.Fatalf("Expected no errors, but got: %v", err)
	}

	c = caddy.NewTestController("dns", `wasm more`)
	if err := setup(c); err == nil {
		t.Fatalf("Expected errors, but got: %v", err)
	}
}
