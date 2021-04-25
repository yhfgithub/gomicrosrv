module gomicroweb

go 1.14

replace (
	github.com/micro/go-micro => github.com/micro/go-micro v1.16.0
	gomicrosrv => github.com/yhfgithub/gomicrosrv v0.0.0-20210310103949-e891803b96e0

)

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.3
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/sirupsen/logrus v1.8.1
	gomicrosrv v0.0.0-00010101000000-000000000000
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)
