package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"net/http"
)

type Post struct {
	name  string //用户名
	email string //密码
}

func selectValue(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()
	name := r.Form["name"]

	db, err := sql.Open("mysql", "root:sheyu@/user?charset=utf8")
	if err != nil {
		return errors.Wrap(err, "Fail to connect to mysql!")
	}
	defer db.Close()

	var post Post
	post.name = name[0]

	var row *sql.Row
	row = db.QueryRow("select name from user where name=?", post.name) //检索数据
	err = row.Scan(&post.name)                                         //遍历！！！
	switch {
	case err == sql.ErrNoRows:
		return errors.Wrap(err, "no result")
	case err != nil:
		return errors.Wrap(err, "get result error")
	default:
		return nil
	}
}
