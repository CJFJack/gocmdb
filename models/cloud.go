package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 云平台对象
type CloudPlatform struct {
	ID          int        `orm:"column(id)"`
	Name        string     `orm:"column(name);size(64)"`
	Type        string     `orm:"column(type);size(32)"`
	Addr        string     `orm:"column(addr);size(1024)"`
	AccessKey   string     `orm:"column(access_key);size(1024)"`
	SecretKey   string     `orm:"column(secret_key);size(1024)"`
	Region      string     `orm:"column(region);size(64)"`
	Remark      string     `orm:"column(remark);size(1024)"`
	CreatedTime *time.Time `orm:"column(created_time);type(datetime);auto_now_add"`
	DeletedTime *time.Time `orm:"column(deleted_time);type(datetime);null"`
	SyncTime    *time.Time `orm:"column(sync_time);type(datetime);null"`
	User        *User      `orm:"column(user);rel(fk)"`
	Status      int        `orm:"column(status)"`
}

func NewCloudPlatform() (cloudPlatform *CloudPlatform) {
	return &CloudPlatform{}
}

// 注册model
func init() {
	orm.RegisterModel(new(CloudPlatform))
}
