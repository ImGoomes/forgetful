package storage

import (
	"os"
	"testing"
	"time"

	"github.com/imgoomes/forgetful/internal/model"
)

// setHome redirects storePath to a temp directory for the duration of the test.
func setHome(t *testing.T) {
	t.Helper()
	t.Setenv("HOME", t.TempDir())
}

func TestLoad_EmptyStore(t *testing.T) {
	setHome(t)

	store, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if store.NextID != 1 {
		t.Errorf("expected NextID=1, got %d", store.NextID)
	}
	if len(store.Entries) != 0 {
		t.Errorf("expected no entries, got %d", len(store.Entries))
	}
}

func TestSaveAndLoad_RoundTrip(t *testing.T) {
	setHome(t)

	original := &model.Store{
		NextID: 3,
		Entries: []*model.Entry{
			{ID: 1, Command: "git status", Tag: "git", Description: "Show status", CreatedAt: time.Now()},
			{ID: 2, Command: "docker ps -a", Tag: "docker", CreatedAt: time.Now()},
		},
	}

	if err := Save(original); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded, err := Load()
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	if loaded.NextID != original.NextID {
		t.Errorf("NextID: want %d, got %d", original.NextID, loaded.NextID)
	}
	if len(loaded.Entries) != len(original.Entries) {
		t.Fatalf("entry count: want %d, got %d", len(original.Entries), len(loaded.Entries))
	}
	for i, e := range loaded.Entries {
		orig := original.Entries[i]
		if e.ID != orig.ID || e.Command != orig.Command || e.Tag != orig.Tag || e.Description != orig.Description {
			t.Errorf("entry %d mismatch: got %+v, want %+v", i, e, orig)
		}
	}
}

func TestLoad_InvalidJSON(t *testing.T) {
	setHome(t)

	path, err := storePath()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(path[:len(path)-len("/commands.json")], 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("not valid json"), 0o644); err != nil {
		t.Fatal(err)
	}

	_, err = Load()
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}
