package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"time"
)

const (
	_DB_NAME       = "data/beeblog.db"
	_SQITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	CreatedTime     time.Time `orm:"index"`
	views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopiclastUserId int64
}
type Topic struct {
	Id               int64
	Uid              int64
	Title            string
	Content          string `orm:"size(5000)"`
	Attachment       string
	CreatedTime      time.Time `orm:"index"`
	views            int64     `orm:"index"`
	Author           string
	ReplyTime        time.Time `orm:"index"`
	ReplayCount      int64
	RepleyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQITE3_DRIVER, orm.DR_Sqlite)
	orm.RegisterDataBase("default", _SQITE3_DRIVER, _DB_NAME, 10)
}
