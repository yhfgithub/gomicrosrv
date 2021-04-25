package conf

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/source/file"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
	"fmt"
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
	err error
)

func Init() {
	m.Lock()
	defer m.Unlock()

	err := config.Load(file.NewSource(
		file.WithPath("./conf/application.yml"),
	))
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("Load 配置信息失败!")
		return
	}
	if err := config.Get(defaultPath, "mysql").Scan(&mc); err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("读取配置信息失败!")
		return
	}

	//开启监控
	 go configWatch()

	//连接数据库
	mysqlOpen()

}

//监控配置文件
func configWatch()  {
	for {
		w, err := config.Watch(defaultPath, "mysql")
		fmt.Println("Watch2")
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("监控配置信息失败!")
		}

		// wait for next value
		_, err = w.Next()
		fmt.Println("Watch3")
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("wait for next value 失败!")
		}

		err = config.Load(file.NewSource(
			file.WithPath("./conf/application.yml"),
		))
		if err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("Load 配置信息失败!")
			return
		}
		if err := config.Get(defaultPath, "mysql").Scan(&mc); err != nil {
			log.WithFields(log.Fields{
				"err": err,
			}).Error("读取配置信息失败!")
			return
		}
		//连接数据库
		mysqlOpen()
	}
}


//初始化数据库
func mysqlOpen()  {
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
