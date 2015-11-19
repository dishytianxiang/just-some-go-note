package main

import (
	"log"
	"github.com/Unknwon/goconfig"
)

func main() {
	cfg,err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件:%s",err)
	}
	value,err := cfg.GetValue(goconfig.DEFAULT_SECTION,"key_default")
	if err != nil {
		log.Fatalf("无法获取值:%s",err)
	}
	log.Printf("%s > %s:%s",goconfig.DEFAULT_SECTION,"key_default",value)
	isInsert := cfg.SetValue(goconfig.DEFAULT_SECTION,"key_default","这是新的值")
	log.Printf("设在键值%s: %v","key_default",isInsert)
	comment := cfg.GetSectionComments("super")	
	log.Printf("分区%s的注释:%s","supper",comment)
	_ = cfg.SetSectionComments("super","# 这是一个新的注释")
	_,err = cfg.Int("must","int")
	err = goconfig.SaveConfigFile(cfg,"conf_save.ini")
	if err != nil {
		log.Fatalf("无法获取值:%s",err)
	}
	
}
