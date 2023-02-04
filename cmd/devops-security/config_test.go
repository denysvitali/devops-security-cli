package main

import (
	"fmt"
	"testing"
)

func TestParseConfig(t *testing.T) {
	cfg, err := ParseConfig()
	if err != nil {
		t.Fatalf("unable to parse config: %v", err)
	}

	fmt.Printf("config=%+v\n", cfg)
}
