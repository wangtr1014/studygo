package main

import (
	"errors"
	"io/ioutil"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Port     int    `ini:"port"`
}

type RedisConfig struct {
	Address string `ini:"address"`
	Port    int    `ini:"port"`
	Db      int    `ini:"db"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	//传进来的data应该是指针类型，因为要在函数中对其赋值，值类型传递的是副本
	//传进来的需要是结构体指针
	t := reflect.TypeOf(data)

	if t.Kind() != reflect.Ptr {
		err = errors.New("data paramter kind should be a pointer")
		return
	}
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data paramter kind should be a struct pointer")
		return
	}
	//读文件
	//一行一行的读
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	var c string
	c = string(contents)
	lines := strings.Split(c, "\n")
	for _, v := range lines {
		v = strings.TrimSpace(v)
		//如果是注释 忽略
		//判断是不是[]
		if strings.HasPrefix(v, "#") || strings.HasPrefix(v, ";") {
			continue
		}
		if strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]") {
			sectionName := strings.TrimSpace(v[1 : len(v)-1])
			if len(sectionName) == 0 {
				err = errors.New("section name should not be empty")
				return
			}
		}
		//根据sectionName确定结构体

		//判断key value
	}
	return
}

func main() {
	var cfg Config
	loadIni("./test.ini", &cfg)
}
