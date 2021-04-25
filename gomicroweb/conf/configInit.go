package conf

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlConf struct {
	Dsn string `json:"dsn"`
	MaxIdleConnection int`json:"maxIdleConnection"`
	MaxOpenConnection int `json:"maxOpenConnection"`
	MaxLifeTime int `json:"maxLifeTime"`
}
var (
	defaultPath="api"
	m sync.RWMutex
	//Tk Token
	mc MysqlConf
	DB *gorm.DB
)

func Init() {
	m.Lock()
	defer m.Unlock()

	err := config.Load(file.NewSource(
		file.WithPath("./conf/application.yml"),
	))

	if err := config.Get(defaultPath, "mysql").Scan(&mc); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("读取配置信息失败!")
		return
	}

	//初始化数据库
	DB, err = gorm.Open("mysql", mc.Dsn)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("打开数据库失败!")
		return
	}

	DB.DB().SetMaxIdleConns(mc.MaxIdleConnection)
	DB.DB().SetMaxOpenConns(mc.MaxOpenConnection)
	DB.DB().SetConnMaxLifetime(time.Duration(mc.MaxLifeTime) * time.Second)

	if err := DB.DB().Ping(); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("数据库连接失败!")
		return
	}

	log.Info("连接数据库成功!")
}