package handler

import (
	context "context"
	log "github.com/sirupsen/logrus"
	"gomicrosrv/conf"
	pb "gomicrosrv/proto"
)

type UserServiceHandler struct {

}

func NewUserRpcHandler() *UserServiceHandler {
	return &UserServiceHandler{}
}

func (u *UserServiceHandler)GetUserById(ctx context.Context, in *pb.IdInfo, res *pb.UserRes)  error{//GetUserById(ctx context.Context, in pb.IdInfo) (res *pb.UserRes, err error){
	res = &pb.UserRes{}
	db := conf.DB.Raw("select ID,Name,Age,t_user.Desc from t_user where id =?", in.Id)
	err := db.Error
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("[sqlService]获取信息失败!")
		res.Code = 500
		res.Message = "获取信息失败!"
		res.Data = &pb.SqlInfo{}
		return err
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
		return err
	}

	log.Info("获取数据库信息成功!")

	res.Data = &data
	res.Code = 200
	res.Message = "获取信息成功!"
	return nil
}