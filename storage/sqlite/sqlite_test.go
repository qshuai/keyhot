package sqlite

import (
	"testing"
)

func TestSqliteStorage_Create(t *testing.T) {
	storage, err := New()
	if err != nil {
		t.Errorf("create db instance error: %s", err)
	}

	err = storage.Create("hi", "你好")
	if err != nil {
		t.Errorf("insert new record error: %s", err)
	}
}
