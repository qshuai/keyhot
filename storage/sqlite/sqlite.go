package sqlite

import (
	"github.com/jmoiron/sqlx"
	"github.com/qshuai/keyhot/model"
	"github.com/qshuai/keyhot/storage/common"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteStorage struct {
	*sqlx.DB
}

func (s *SqliteStorage) Create(word, explain string) error {
	record, err := s.GetRecordByOrigin(word)
	if err != nil {
		if err == common.NotFound {
			// create new record
			_, err = s.DB.Exec("insert into word (`origin`, `target`) VALUES (?, ?)", word, explain)
			return err
		}

		return err
	}

	// record exists
	_, err = s.DB.Exec("update word set hits = hits + 1 where id = ?", record.ID)
	return err
}

func (s *SqliteStorage) GetRecordByOrigin(word string) (*model.Word, error) {
	query, err := s.DB.Query("select id from word where origin = ?", word)
	if err != nil {
		return nil, err
	}

	var ret []*model.Word
	err = sqlx.StructScan(query, &ret)
	if err != nil {
		return nil, err
	}

	if len(ret) <= 0 {
		return nil, common.NotFound
	}

	return ret[0], nil
}

func New() (*SqliteStorage, error) {
	db, err := sqlx.Connect("sqlite3", "word.db")
	if err != nil {
		return nil, err
	}

	return &SqliteStorage{
		DB: db,
	}, nil
}
