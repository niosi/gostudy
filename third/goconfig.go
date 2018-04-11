package main

import "github.com/Unknwon/goconfig"
import "log"

func main() {
	cfg,err := goconfig.LoadConfigFile("static/conf.ini");
	if err!=nil{
		log.Fatalf("read conf fail")
		return
	}
	str, _ := cfg.GetValue("Demo", "key1")
	println(str)
}
