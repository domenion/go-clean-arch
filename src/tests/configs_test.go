package main

import (
	"go-clean-arch/src/internal/configs"
	"testing"
)

func TestGetConfig(t *testing.T) {
	cfg, err := configs.GetConfig()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	if cfg == nil {
		t.Errorf("cfg is nil")
		return
	}
}
