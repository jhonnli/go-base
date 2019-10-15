package model

import (
	"time"
)

type Users struct {
	UserId     int       `db:"user_id"`    //用户ID
	Name       string    `db:"name"`       //用户名
	Age        int       `db:"age"`        //年龄
	CreateTime time.Time `db:"createTime"` //创建时间
}
