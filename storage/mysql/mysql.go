package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/qshuai/keyhot/model"
	"github.com/qshuai/keyhot/storage/common"
)

type MysqlStorage struct {
	username string // 用户名
	passwd   string // 密码
	host     string // 主机地址
	port     int    // 端口

	*sqlx.DB // db 资源
}

func New(host string, port int, username, passwd, database string) (*MysqlStorage, error) {
	db, err := connect(host, port, username, passwd, database)
	if err != nil {
		return nil, err
	}

	return &MysqlStorage{
		username: username,
		passwd:   passwd,
		host:     host,
		port:     port,
		DB:       db,
	}, nil
}

func (m *MysqlStorage) Create(word, explain string) error {
	record, err := m.GetRecordByOrigin(word)
	if err != nil {
		if err == common.NotFound {
			// create new record
			_, err = m.DB.Exec("insert into word (`origin`, `target`) VALUES (?, ?)", word, explain)
			return err
		}

		return err
	}

	// record exists
	_, err = m.DB.Exec("update word set hits = hits + 1 where id = ?", record.ID)
	return err
}

func (m *MysqlStorage) GetRecordByOrigin(word string) (*model.Word, error) {
	query, err := m.DB.Query("select id from word where origin = ?", word)
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

func connect(host string, port int, username, passwd, database string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci", username, passwd, host, port, database))
	if err != nil {
		return nil, err
	}

	return db, nil
}
