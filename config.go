package main

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var configFile = "./gowatch.toml"

type config struct {
	AppName string
	//指定ouput执行的程序路径
	Output string
	//需要追加监听的文件后缀名字，默认是'.go'，
	WatchExts []string
	//需要追加监听的目录，默认是当前文件夹，
	WatchPaths []string
	//执行时的额外参数
	CmdArgs []string
	//执行时追加的环境变量
	Envs []string
	//vendor 目录下的文件是否也监听
	VendorWatch bool
	//不需要监听的目录
	ExcludedPaths []string
	//需要编译的包或文件,优先使用-p参数
	BuildPkg string
	//在go build 时期接收的-tags参数
	BuildTags string
}

func parseConfig() *config {
	c := &config{}
	filename, _ := filepath.Abs(configFile)
	if !fileExist(filename) {
		return c
	}

	_, err := toml.DecodeFile(filename, c)
	if err != nil {
		panic(err)
	}
	return c
}

func fileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
