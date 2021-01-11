package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	One One `yaml:"one"`
}
type One struct {
	Database Database `yaml:"database"`
}

type Database struct {
	ST       int    `yaml:"type"`
	One      int `yaml:"one"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	Address  string `yaml:"address"`
}

var con []byte
const configFilePath string = "test.yaml"

func GetConfig()(config Config,err error)  {
	err =yaml.Unmarshal(con,&config)
	return config,err
}

func main()  {
	conf,err:=GetConfig()
	if err!=nil{
		log.Fatalf("GetConfig err: %v",err)
		return
	}
	fmt.Println(conf)
}

func init()  {
	var err error
	con,err = ioutil.ReadFile("test.yaml")
	if err!=nil{
		log.Fatalf("yamlFile.Get err %v",err)
	}
}