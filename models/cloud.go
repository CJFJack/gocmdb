package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

// 云平台对象
type CloudPlatform struct {
	ID              int               `orm:"column(id)"`
	Name            string            `orm:"column(name);size(64)"`
	Type            string            `orm:"column(type);size(32)"`
	Addr            string            `orm:"column(addr);size(1024)"`
	AccessKey       string            `orm:"column(access_key);size(1024)"`
	SecretKey       string            `orm:"column(secret_key);size(1024)"`
	Region          string            `orm:"column(region);size(64)"`
	Remark          string            `orm:"column(remark);size(1024)"`
	CreatedTime     *time.Time        `orm:"column(created_time);type(datetime);auto_now_add"`
	DeletedTime     *time.Time        `orm:"column(deleted_time);type(datetime);null"`
	SyncTime        *time.Time        `orm:"column(sync_time);type(datetime);null"`
	User            *User             `orm:"column(user);rel(fk)"`
	Status          int               `orm:"column(status);description(0-启用,1-禁用)"`
	VirtualMachines []*VirtualMachine `orm:"reverse(many)"`
	Msg             string            `orm:"column(msg);null"`
}

func NewCloudPlatform() (cloudPlatform *CloudPlatform) {
	return &CloudPlatform{}
}

// 云主机对象
type VirtualMachine struct {
	ID            int            `orm:"column(id)"`
	Platform      *CloudPlatform `orm:"column(platform);rel(fk)"`
	UUID          string         `orm:"column(uuid);size(128)"`
	Name          string         `orm:"column(name);size(64)"`
	CPU           int            `orm:"column(cpu)"`
	Mem           int64          `orm:"column(mem)"`
	OS            string         `orm:"column(os);size(128)"`
	PrivateAddrs  string         `orm:"column(private_addr);size(1024)"`
	PublicAddrs   string         `orm:"column(public_addr);size(1024)"`
	Status        string         `orm:"column(status);size(32)"`
	VmCreatedTime string         `orm:"column(vm_created_time);size(100)"`
	VmExpiredTime string         `orm:"column(vm_expired_time);size(100)"`
	CreatedTime   *time.Time     `orm:"column(created_time);auto_now_add;type(datetime)"`
	UpdatedTime   *time.Time     `orm:"column(updated_time);auto_now;type(datetime)"`
	DeletedTime   *time.Time     `orm:"column(deleted_time);null;type(datetime)"`
}

func NewVirtualMachine() (cloudPlatform *VirtualMachine) {
	return &VirtualMachine{}
}

// 注册model
func init() {
	orm.RegisterModel(new(CloudPlatform), new(VirtualMachine))
}
