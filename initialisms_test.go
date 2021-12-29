package tbox

import (
	"testing"
)

type param struct {
	Input  string
	Except string
}

func TestName(t *testing.T) {
	var list = []param{
		{"target_id", "targetID"},
		{"a_id", "aID"},
		{"k_app_id", "kAPPID"},
		{"app_id", "appID"},
	}

	for _, val := range list {
		should := lintName(val.Input)
		if should != val.Except {
			t.Logf("input %s except:%s,but got:%s", val.Input, val.Except, should)
		}
	}
}
