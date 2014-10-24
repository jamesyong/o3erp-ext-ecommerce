o3erp-ext-ecommerce
===================

Ecommerce Demo for O3 ERP


Quick Start Guide
=================
Install Go 1.3<br>
Install Thrift (latest)<br>
Install Java 7<br>

Open a terminal and type<br> 
```
	go get github.com/jamesyong/o3erp-ext-ecommerce
```

Go to $gopath/src/github.com/jamesyong/o3erp/backend/framework/base/ofbiz-component.xml, and
update path-to-key-store to correct path value.

In the terminal, change directory to $gopath/src/github.com/jamesyong/o3erp-ext-ecommerce/backend, and type
```
	ant load-demo
```

At this point, you can run the unit tests with
```
	ant run-tests
```

To start the backend, type
```
	ant start
```

Open another terminal, change directory to $gopath/src/github.com/jamesyong/o3erp-ext-ecommerce, and type
```
	go run server.go
```

To access the ecommerce store:
Open a browser and navigate to http://localhost:9080/ecommerce

To access the management modules:
Open a browser, navigate to https://localhost:8443, and type admin/ofbiz to login.
