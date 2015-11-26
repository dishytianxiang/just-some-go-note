package main
import (
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-xorm/xorm"
	"log"
)
type Account struct {
	Id  int64
	Name  string `xorm:"unique"`
	Balance  float64
	Version  int `xorm:"version"`
}
var x *xorm.Engine 

func init() {
	var err error
	x,err = xorm.NewEngine("sqlite3", "./bank.db")
	if err !=nil {
		log.Fatal("fail to create engine: %v",err)
	}
	if err = x.Sync(new(Account));err !=nil {
		log.Fatal("fail to Sync database: %v",err)
	}	
}
func newAccount(name string,balance float64) error {
	_,err := x.Insert(&Account{Name: name,Balance: balance})
	return err
}
func getAccount(id int64) (*Account,error){
	a := &Account{}
	has,err := x.Id(id).Get(a)
	if err !=nil {
		return nil,err
	}else if !has {
		return nil,errors.New("Acount not fount")
	}
	return a,nil
}
