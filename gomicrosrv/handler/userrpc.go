package handler

import (
	context "context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gomicrosrv/conf"
	pb "gomicrosrv/proto"
)

type MyUserServiceHandler struct {

}

func NewUserRpcHandler() *MyUserServiceHandler {
	return &MyUserServiceHandler{}
}

func (u *MyUserServiceHandler)GetUserById(ctx context.Context, in *pb.IdInfo, res *pb.UserRes) error {//GetUserById(ctx context.Context, in pb.IdInfo) (res *pb.UserRes, err error){
	db := conf.DB.Raw("select ID,Name,Age,t_user.Desc from t_user where id =?", in.Id)
	err := db.Error
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("[sqlService]获取信息失败!")
		res.Code = 500
		res.Message = "获取信息失败!"
		res.Data = &pb.SqlInfo{}
		return nil
	}

	var data pb.SqlInfo
	r := db.Row()
	err = r.Scan(&data.Id,&data.Name,&data.Age,&data.Desc)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("[sqlService] 传递信息失败!")

		res.Code = 500
		res.Message = "获取信息失败!"
		res.Data = &pb.SqlInfo{}
		return nil
	}

	log.Info("获取数据库信息成功!")
	fmt.Println("data:",data)
	res.Data = &data
	res.Code = 200
	res.Message = "获取信息成功!"
	return nil
}