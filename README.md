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

Go to $GOPATH/src/github.com/jamesyong/o3erp/backend/framework/base/ofbiz-component.xml, and
update path-to-key-store to correct path value.

In the terminal, type
```
	cd $GOPATH/src/github.com/jamesyong/o3erp-ext-ecommerce/backend
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

Open another terminal and type
```
    cd $GOPATH/src/github.com/jamesyong/o3erp-ext-ecommerce
	go run server.go
```

To access the ecommerce store:
Open a browser and navigate to http://localhost:9080/ecommerce

To access the management modules:
Open a browser, navigate to https://localhost:8443, and type admin/ofbiz to login.

Switch to Postgres
==================
Derby is the default database used in O3 ERP. To switch to Postgres, do the following:

From the terminal, type
```
    cd $GOPATH/src/github.com/jamesyong/o3erp-ext-ecommerce
    ant download-PG-JDBC
    cd backend/hot-deploy/entity-data/config
    sudo nano -w entityengine.xml
```
Look for 
```
<delegator name="default" entity-model-reader="main" entity-group-reader="main" entity-eca-reader="main" distributed-cache-clear-enabled="false">
        <group-map group-name="org.ofbiz" datasource-name="localderby"/>
        <group-map group-name="org.ofbiz.olap" datasource-name="localderbyolap"/>
        <group-map group-name="org.ofbiz.tenant" datasource-name="localderbytenant"/>
        <!-- <group-map group-name="org.ofbiz" datasource-name="localmysql"/>
        <group-map group-name="org.ofbiz.olap" datasource-name="localmysqlolap"/>
        <group-map group-name="org.ofbiz.tenant" datasource-name="localmysqltenant"/>  -->
        <!-- <group-map group-name="org.ofbiz" datasource-name="localpostnew"/>
        <group-map group-name="org.ofbiz.olap" datasource-name="localpostolap"/>
```

Comment out
```
		<group-map group-name="org.ofbiz" datasource-name="localderby"/>
        <group-map group-name="org.ofbiz.olap" datasource-name="localderbyolap"/>
```

Uncomment
```
		<group-map group-name="org.ofbiz" datasource-name="localpostnew"/>
        <group-map group-name="org.ofbiz.olap" datasource-name="localpostolap"/>
```

Edit values for the following keys to match the database that you will create / have created for O3 ERP
```
	jdbc-uri, jdbc-username, jdbc-password
```
at
```
<datasource name="localpostnew"
        helper-class="org.ofbiz.entity.datasource.GenericHelperDAO"
        schema-name="public"
        field-type-name="postnew"
        check-on-start="true"
        add-missing-on-start="true"
        use-fk-initially-deferred="false"
        alias-view-columns="false"
        join-style="ansi"
        result-fetch-size="50"
        use-binary-type-for-blob="true"
        use-order-by-nulls="true">
		...
        <inline-jdbc
            jdbc-driver="org.postgresql.Driver"
            jdbc-uri="jdbc:postgresql://127.0.0.1/ofbiz"
            jdbc-username="ofbiz"
            jdbc-password="ofbiz"
            isolation-level="ReadCommitted"
            pool-minsize="2"
            pool-maxsize="250"
            time-between-eviction-runs-millis="600000"/>
        ...
    </datasource>
    <datasource name="localpostolap"
            helper-class="org.ofbiz.entity.datasource.GenericHelperDAO"
            schema-name="public"
            field-type-name="postnew"
            check-on-start="true"
            add-missing-on-start="true"
            use-fk-initially-deferred="false"
            alias-view-columns="false"
            join-style="ansi"
            result-fetch-size="50"
            use-binary-type-for-blob="true"
            use-order-by-nulls="true">
		...
        <inline-jdbc
                jdbc-driver="org.postgresql.Driver"
                jdbc-uri="jdbc:postgresql://127.0.0.1/ofbizolap"
                jdbc-username="ofbiz"
                jdbc-password="ofbiz"
                isolation-level="ReadCommitted"
                pool-minsize="2"
                pool-maxsize="250"
                time-between-eviction-runs-millis="600000"/>
		...
    </datasource>
```
Extend/Override Properties
==========================
Properties files located in config and i18n folders under
$GOPATH/src/github.com/jamesyong/o3erp/backend/ can be overrided and/or extended.

See the following files for examples:
$GOPATH/src/github.com/jamesyong/o3erp-ext-ecommerce/backend/hot-deploy/demo-data/config/ext.general.properties
$GOPATH/src/github.com/jamesyong/o3erp-ext-ecommerce/backend/hot-deploy/demo-data/i18n/ext.AccountingUiLabels.properties



