package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckNQuad(t *testing.T) {
	if err := checkNQuad("", "name", "", Str("Alice")); err == nil {
		t.Fatal(err)
	}
	if err := checkNQuad("alice", "", "", Str("Alice")); err == nil {
		t.Fatal(err)
	}
	if err := checkNQuad("alice", "name", "", nil); err == nil {
		t.Fatal(err)
	}
	if err := checkNQuad("alice", "name", "id", Str("Alice")); err == nil {
		t.Fatal(err)
	}
}

func TestSetMutation(t *testing.T) {
	req := NewRequest()

	if err := req.SetMutation("alice", "name", "", Str("Alice"), ""); err != nil {
		t.Fatal(err)
	}
	if err := req.SetMutation("alice", "falls.in", "", Str("rabbithole"), ""); err != nil {
		t.Fatal(err)
	}
	if err := req.DelMutation("alice", "falls.in", "", Str("rabbithole"), ""); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(req.gr.Mutation.Set), 2, "Set should have 2 entries")
	assert.Equal(t, len(req.gr.Mutation.Del), 1, "Del should have 1 entry")
}
