package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gomicroweb/model"
	"gomicroweb/conf"
)

type GinHttpHandler struct {

}

//put信息
func (g *GinHttpHandler) PutInfo(c *gin.Context) {

	var user model.User
	var res model.SimpleResponse
	err := c.ShouldBindJSON(&user)
	if err !=nil{
		fmt.Println("ShouldBindJSON:",err)
		res.Code = 400
		res.Message = "传入参数不对!"
		c.JSON(200,res)
	}
	db := conf.DB.Exec("insert into t_user( `Name`,`Age` , `Desc`) values(?,?,?)", user.Name,user.Age,user.Desc)
	err = db.Error
	if err != nil {
		fmt.Println("db:",err)
		res.Code = 500
		res.Message = "获取信息失败!"
		c.JSON(200,res)
	}
	log.Info("insert数据库信息成功!")
	res.Code = 200
	res.Message = "insert信息成功!"
	c.JSON(200,res)
}

