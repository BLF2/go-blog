package test

import (
	"dao"
	"database/sql"
	"log"
	"model"
	"time"
)

func main() {

	//新增
	var userInfo model.UserInfo = model.UserInfo{
		LoginId:    sql.NullString{"hitol", true},
		NickName:   sql.NullString{String: "zhuhaitao", Valid: true},
		Password:   sql.NullString{String: "123456", Valid: true},
		CreateTime: time.Now(),
	}

	var id int64
	var err error
	id, err = dao.PublicUserInfoDao.Insert(&userInfo)

	if err != nil {
		log.Fatalln("插入用户信息失败，userInfo=", userInfo, err)
	} else {
		log.Println("插入用户数据成功，id=", id)
		userInfo.Id = id
	}
	//查询
	var userInfoQuery *model.UserInfo
	userInfoQuery, err = dao.PublicUserInfoDao.SelectById(id)
	if err != nil {

		log.Fatalln("根据id查询客户信息失败，id=", id, err)
	} else {
		if userInfoQuery == nil {

			log.Println("根据id查询客户信息为空，id=", id)
		} else {
			log.Println("根据id查询客户信息成功，userInfo=", userInfoQuery)
		}
	}

	//修改
	var userInfoUpd model.UserInfo = model.UserInfo{
		Id:       userInfo.Id,
		LoginId:  sql.NullString{"zjj", true},
		NickName: sql.NullString{"朱海涛", true},
	}
	err = dao.PublicUserInfoDao.UpdateBySelectivePk(&userInfoUpd)
	if err != nil {

		log.Fatalln("根据id更新用户信息失败，userInfo=", userInfoUpd)
	}
	//删除
	//err = dao.PublicUserInfoDao.DeleteById(id)
	//if err != nil{
	//	log.Fatalln("根据id删除用户信息失败，id=",id,err)
	//}

	userInfoAll := dao.PublicUserInfoDao.SelectAll()

	log.Println("查询所有用户信息", userInfoAll)
}
