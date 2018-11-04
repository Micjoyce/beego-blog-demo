package models

import (
	"os"
	"path"
	"time"

	"github.com/astaxie/beego/orm"
	// some comment
	_ "github.com/mattn/go-sqlite3"
	"github.com/unknwon/com"
)

// DBNAME dbnme
// SQLITE3DRIVER driver
const (
	DBNAME        = "data/beeblog.db"
	SQLITE3DRIVER = "sqlite3"
)

// Category define cetetory model
type Category struct {
	ID              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserID int64
}

// Topic define topic model
type Topic struct {
	ID              int64
	UID             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastuserID int64
}

func RegisterDB() {
	if !com.IsExist(DBNAME) {
		os.MkdirAll(path.Dir(DBNAME), os.ModePerm)
		os.Create(DBNAME)
	}
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(SQLITE3DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", SQLITE3DRIVER, DBNAME, 10)
}
