module gomicrosrv

go 1.16
replace (
	github.com/micro/go-micro => github.com/micro/go-micro v1.16.0
)
require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/sirupsen/logrus v1.8.1
)
