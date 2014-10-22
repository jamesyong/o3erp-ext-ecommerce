o3erp-ext-ecommerce
===================

Ecommerce Demo for O3 ERP


Quick Start Guide
=================
Install Go 1.3
Install Thrift
Install Java 7

Open a terminal and type
	go get github.com/jamesyong/o3erp-ext-ecommerce

Go to $gopath/src/github.com/jamesyong/o3erp/backend/framework/base/ofbiz-component.xml, and
update path-to-key-store to correct path value.

In the terminal, change directory to $gopath/src/github.com/jamesyong/o3erp-ext-ecommerce/backend, and type
	ant load-demo
	ant start

Open another terminal, change directory to $gopath/src/github.com/jamesyong/o3erp-ext-ecommerce, and type
	go run server.go


Open a browser, navigate to https://localhost:8443, and type admin/ofbiz to login.
