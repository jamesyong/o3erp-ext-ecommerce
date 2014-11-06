package main

import (
	"fmt"
	frontend "github.com/jamesyong/o3erp/frontend"
	"github.com/jamesyong/o3erp/frontend/config"
	"os"
)

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	config.WORKING_DIRECTORY = pwd
	config.PATH_BASE_FRONTEND_TEMPLATES = pwd + "/../o3erp/frontend/templates"
	config.PATH_BASE_FRONTEND_ASSETS = pwd + "/../o3erp/frontend/assets"
}

func main() {
	frontend.Startup()
}
