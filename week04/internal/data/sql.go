package data

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"week04/internal/conf"

	_ "github.com/go-sql-driver/mysql"
)

const sqlFormat = "%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=1000ms"

var ErrNoUser = errors.New("GetUserInfo: No such user")

type DBModel struct {
	dbEngine *sql.DB
}

func NewDBModel(info *conf.Data) *DBModel {

	connUrl := fmt.Sprintf(
		sqlFormat,
		info.Db.UserName,
		info.Db.Password,
		info.Db.Host,
		info.Db.DatabaseName,
		info.Db.Charset,
	)

	db, err := sql.Open(info.Db.DBType, connUrl)
	if err != nil {
		log.Fatal("open db failed:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("The ping of %s has failed failed: %s\n", connUrl, err)
	}

	return &DBModel{dbEngine: db}
}

func (d *DBModel) GetUserName(account, query string) (name string, err error) {

	account = strings.TrimSpace(account)
	if account == "" {
		return "", errors.New("GetUserInfo: account is blank")
	}

	err = d.dbEngine.QueryRow(query, account).Scan(&name)
	switch err {
	case sql.ErrNoRows:
		return "", ErrNoUser
	default:
		return
	}
}
