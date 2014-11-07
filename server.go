package main

import (
	"fmt"
	"github.com/jamesyong/o3erp/go/config"
	frontend "github.com/jamesyong/o3erp/go/server"
	"os"
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	config.WORKING_DIRECTORY = pwd
	config.PATH_BASE_GOLANG_TEMPLATES = pwd + "/../o3erp/go/templates"
	config.PATH_BASE_GOLANG_ASSETS = pwd + "/../o3erp/go/assets"
}

func main() {
	frontend.Startup()
}
