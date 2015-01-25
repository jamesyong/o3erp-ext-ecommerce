package main

import (
	"fmt"
	"github.com/jamesyong/o3erp/go/config"
	frontend "github.com/jamesyong/o3erp/go/server"
	"github.com/jamesyong/o3erp/go/templating"
	"os"
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	config.WORKING_DIRECTORY = pwd
	config.PATH_BASE_GOLANG_TEMPLATES = pwd + "/../go/templates"
	config.PATH_BASE_GOLANG_ASSETS = pwd + "/../go/assets"
	config.PATH_BASE_GOLANG_CERT = pwd + "/../../" + os.Getenv("ext.dir") + "/../go"
}

func main() {
	templating.Options.IsDevelopment = true
	frontend.Startup()
}
