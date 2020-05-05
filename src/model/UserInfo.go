package model

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

type UserInfo struct {
	Id int64 `db:"id"`
	LoginId sql.NullString `db:"login_id"`
	NickName sql.NullString `db:"nick_name"`
	Password sql.NullString `db:"password"`
	CreateTime time.Time  `db:"create_time"`
	CreateUser sql.NullString `db:"create_user"`
}

func GetUserInfoAllColumns() string {

	return "id," + GetUserInfoColumnsExcludeId()
}

func GetUserInfoColumnsExcludeId() string  {

	return "login_id,nick_name,password,create_time,create_user"
}

func GetUserInfoValuesExcludeIdForInsert() string {

	return "?,?,?,?,?"
}

func (userInfo *UserInfo)GetUserInfoUpdSelectiveFields() (string,[]interface{},error){

	var set string = ""
	var values []interface{} = []interface{}{}
	if userInfo.LoginId.Valid {

		set += "login_id = ?,"
		values = append(values,userInfo.LoginId.String)
	}

	if userInfo.NickName.Valid{

		set += "nick_name = ?,"
		values = append(values,userInfo.NickName.String)
	}

	if userInfo.Password.Valid{

		set += "password = ?"
		values = append(values, userInfo.Password.String)
	}

	if set == "" {
		var err error = errors.New("可空字段全为空")
		return "",[]interface{}{},err
	}

	if strings.HasSuffix(set,","){

		set = set[:len(set) - 1]
	}

	return set,values,nil
}
