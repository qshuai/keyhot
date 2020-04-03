package mysql

import (
	"testing"
)

func TestGetRecordByOrigin(t *testing.T) {
	storage, err := New("127.0.0.1", 3306, "root", "646689abc", "word_note")
	if err != nil {
		t.Errorf("connect to database failed: %s", err)
	}

	record, err := storage.GetRecordByOrigin("hello")
	if err != nil {
		t.Errorf("fetch record by key word error: %s", err)
	}

	t.Logf("record: %+v", record)
}
