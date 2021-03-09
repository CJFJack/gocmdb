package models

import (
	"github.com/astaxie/beego/orm"
	"gocmdb/utils"
	"time"
)

// User 用户对象
type User struct {
	ID         int        `orm:"column(id)"`
	StaffID    string     `orm:"column(staff_id);size(32)"`
	Name       string     `orm:"size(64)"`
	NickName   string     `orm:"size(64)"`
	Password   string     `orm:"size(1024)" form:"Password"`
	Gender     int        `orm:"description(1-男,0-女)"`
	Tel        string     `orm:"size(32)"`
	Addr       string     `orm:"size(128)"`
	Email      string     `orm:"size(64)"`
	Department string     `orm:"size(128)"`
	Status     int        `orm:"description(0-正常,1-锁定，2-离职)"`
	CreatedAt  *time.Time `orm:"auto_now_add"`
	UpdatedAt  *time.Time `orm:"auto_now"`
	DeletedAt  *time.Time `orm:"null"`
	Deleted    int
}

func NewUser() (user *User) {
	return &User{}
}

// 注册model
func init() {
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

// 验证用户密码是否正确
func (u *User) ValidPassword(password string) bool {
	return utils.CheckPassword(u.Password, password)
}
