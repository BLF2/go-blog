package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"go-blog/model"
	"log"
	"strconv"
)

type UserInfoDao struct {

}

var PublicUserInfoDao *UserInfoDao

func (*UserInfoDao)Insert(userInfo *model.UserInfo) (int64,error){


	result := DbCon.MustExec(fmt.Sprintf("insert into user_info (%s) values(%s)",
		model.GetUserInfoColumnsExcludeId(), model.GetUserInfoValuesExcludeIdForInsert()),
		userInfo.LoginId.String,userInfo.NickName.String,userInfo.Password.String,userInfo.CreateTime,userInfo.CreateUser.String)

	var id int64
	var err error

	id,err = result.LastInsertId()

	if err != nil{

		log.Println("插入数据出错,userInfo=",userInfo,err)
	}

	return id,err

}

func (*UserInfoDao)DeleteById(id int64) error {

	result := DbCon.MustExec("delete from user_info where id = ?",id)

	rows,err := result.RowsAffected()

	if err != nil{

		log.Println("根据id删除用户信息失败，id=",id,err)
	}else{

		if rows != 1{
			err = errors.New("删除失败，数据库影响行数=" + strconv.FormatInt(rows,10))
		}
	}

	return err
}

func (*UserInfoDao)UpdateBySelectivePk(userInfo *model.UserInfo) error {

	//result := DbCon.MustExec("update user_info set login_id = ?,nick_name = ?," +
	//	"password=? where id = ?",userInfo.LoginId,userInfo.NickName,userInfo.Password,userInfo.Id)

	set,values,err := userInfo.GetUserInfoUpdSelectiveFields()

	if err != nil{

		return err
	}

	result := DbCon.MustExec(fmt.Sprintf("update user_info set %s where id = ?",set), append(values, userInfo.Id)...)
	rows , err := result.RowsAffected()

	if err != nil{

		log.Println("根据id更新用户信息失败，userInfo=",userInfo,err)
	}else{
		if rows != 1{
			err = errors.New("根据id更新用户信息时影响行数不为1，rows=" + strconv.FormatInt(rows,10))
		}
	}

	return err
}

func (*UserInfoDao)SelectAll() *[]model.UserInfo {

	var userInfoAll *[]model.UserInfo = &[]model.UserInfo{}
	err := DbCon.Select(userInfoAll,fmt.Sprintf("select %s from user_info", model.GetUserInfoAllColumns()))

	if err != nil{

		if err == sql.ErrNoRows{

			userInfoAll = nil
		}else{

			log.Fatal("UserInfoDao-查询所有用户失败",err)
		}
	}

	return userInfoAll
}

func (*UserInfoDao)SelectById(id int64) (*model.UserInfo,error)  {

	var userInfo *model.UserInfo = &model.UserInfo{}
	err := DbCon.Get(userInfo,fmt.Sprintf("select %s from user_info where id = ?", model.GetUserInfoAllColumns()),id)

	if err != nil{

		if err == sql.ErrNoRows{
			err = nil
			userInfo = nil
		}else{
			log.Println("根据id查询用户信息失败，id=",id,err)
		}
	}

	return userInfo,err
}

func (*UserInfoDao)SelectByLoginId(loginId string) (*model.UserInfo,error){

	var userInfo *model.UserInfo = &model.UserInfo{}

	err := DbCon.Get(userInfo,fmt.Sprintf("select %s from user_info where login_id = ?",
		model.GetUserInfoAllColumns()),loginId)

	if err == sql.ErrNoRows{

		log.Println("根据loginId查询数据为空,loginId=",loginId)
		userInfo = nil
	}else {

		log.Println("根据loginId查询数据失败,loginId=",loginId,err)
	}

	return userInfo,err
}