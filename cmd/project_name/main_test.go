package main

import (
	"testing"

	"go.uber.org/fx"
)

func TestOptions(t *testing.T) {
	opts := Options

	if err := fx.ValidateApp(opts...); err != nil {
		t.Fatalf("fx.Options binding failed: %v", err)
	}
}
