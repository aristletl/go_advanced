/*我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
	该error是需要抛给上层调用者的，dao作为数据操作层，不应对错误去进行处理，应返回给调用者统一处理
 */

package main

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	pErr "github.com/pkg/errors"
)

type Dao struct {
	db *sql.DB
}

func NewDao(connUrl string) *Dao {
	db, err := sql.Open("mysql", connUrl)
	if err != nil {
		log.Fatal("open db failed:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("The ping of %s has failed failed: %s\n", connUrl, err)
	}
	return &Dao{db: db}
}

func (d *Dao) GetUserInfo(account, query string) (name string, err error) {

	account = strings.TrimSpace(account)
	if account == "" {
		return "", errors.New("dao.GetUserInfo: account is blank")
	}

	e := d.db.QueryRow(query, account).Scan(&name)
	err = pErr.Wrap(e, "dao.GetUserInfo: ")
	return
}

func main() {
	connUrl := "root:123456@tcp(localhost:3306)/go_advanced?charset=utf8"
	account := "test"
	query := `select * from user where account =?`
	
	dao := NewDao(connUrl)
	_, err := dao.GetUserInfo(account, query)
	if err != nil {
		switch errors.Unwrap(err) {
		case sql.ErrNoRows:
			log.Printf("user is not exit with account: %s, %+v\n", account, err)
		default:
			log.Println(err)
		}
	}
}
