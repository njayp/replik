package env

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	os.Setenv("REPLIK_PORT", "1")
	e := SetConfig()
	if e.Port != 1 {
		t.Error(e.Port)
	}
}

func TestEnvDefault(t *testing.T) {
	e := SetConfig()
	if e.Port != 9090 {
		t.Error(e.Port)
	}
}
