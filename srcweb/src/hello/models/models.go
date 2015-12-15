package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"time"
	"strconv"
)

const (
	_DB_NAME       = "data/beeblog.db"
	_SQITE3_DRIVER = "sqlite3"
)
type User struct {
	Id 				int64
	Email			string
	Passwd			string
	CreatedTime		time.Time `orm:"index"`
}
type Category struct {
	Id              int64
	Title           string
	CreatedTime     time.Time `orm:"index"`
	views           int64     `orm:"index"`
	TopicTime       time.Time 
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
	orm.RegisterModel(new(Category), new(Topic),new(User))
	orm.RegisterDriver(_SQITE3_DRIVER, orm.DR_Sqlite)
	orm.RegisterDataBase("default", _SQITE3_DRIVER, _DB_NAME, 10)
}
func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name,CreatedTime: time.Now(),TopicTime: time.Now()}
	qs := o.QueryTable("category")
	err := qs.Filter("title",name).One(cate)
	if err == nil {
		return err
	}
	_,err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}
func DelCategory(id string) error {
	cid,err := strconv.ParseInt(id,10,64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_,err = o.Delete(cate)
	return err
}
func GetAllCategories() ([]*Category,error){
	o := orm.NewOrm()
	cates := make([]*Category,0)
	qs := o.QueryTable("category")
	_,err := qs.All(&cates)
	return cates,err
}
