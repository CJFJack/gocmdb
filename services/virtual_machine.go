package services

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"gocmdb/cloud"
	"gocmdb/models"
	"strings"
	"time"
)

type virtualMachineService struct {
}

// 查询云平台
func (s *virtualMachineService) Query(q string, limit, offset int) ([]*models.VirtualMachine, int64) {
	var virtualMachine []*models.VirtualMachine
	querySet := orm.NewOrm().QueryTable(&models.VirtualMachine{}).RelatedSel()

	cond := orm.NewCondition()
	cond = cond.And("deleted_time__isnull", true)

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("private_addr__icontains", q)
		query = query.Or("public_addr__icontains", q)
		cond.AndCond(query)
	}

	querySet = querySet.SetCond(cond)
	total, _ := querySet.Count()
	querySet.Limit(limit).Offset(offset).All(&virtualMachine)

	return virtualMachine, total
}

// 根据 id 查询云主机
func (s *virtualMachineService) GetByPk(pk int) *models.VirtualMachine {
	vm := &models.VirtualMachine{}
	ormer := orm.NewOrm()
	if err := ormer.QueryTable(vm).RelatedSel().Filter("ID", pk).Filter("DeletedTime__isnull", true).One(vm); err == nil {
		return vm
	}
	return nil
}

// 新增云平台
func (s *virtualMachineService) Add(model *models.CloudPlatform, user *models.User) error {
	ormer := orm.NewOrm()
	model.User = user
	_, _, err := ormer.ReadOrCreate(model, "Name")
	return err
}

// 同步保存云主机
func (s *virtualMachineService) SyncInstances(instance *cloud.Instance, platform *models.CloudPlatform) {
	// 存在则更新，不存在则创建
	fmt.Println(instance)
	ormer := orm.NewOrm()
	vm := models.VirtualMachine{UUID: instance.UUID, Platform: platform}
	if _, _, err := ormer.ReadOrCreate(&vm, "UUID", "Platform"); err != nil {
		return
	}
	vm.Name = instance.Name
	vm.OS = instance.OS
	vm.CPU = instance.CPU
	vm.Mem = instance.Mem
	vm.Status = instance.Status
	vm.VmCreatedTime = instance.CreatedTime
	vm.VmExpiredTime = instance.ExpiredTime
	vm.PublicAddrs = strings.Join(instance.PublicAddrs, ",")
	vm.PrivateAddrs = strings.Join(instance.PrivateAddrs, ",")
	ormer.Update(&vm)

}

// 更新云主机删除时间
func (s *virtualMachineService) SyncInstanceStatus(now time.Time, platform *models.CloudPlatform) {
	ormer := orm.NewOrm()
	ormer.QueryTable(&models.VirtualMachine{}).Filter("Platform__exact", platform).Filter("UpdatedTime__lt", now).Update(orm.Params{"DeletedTime": now})
	ormer.QueryTable(&models.VirtualMachine{}).Filter("Platform__exact", platform).Filter("UpdatedTime__gt", now).Update(orm.Params{"DeletedTime": nil})

}

// 用户操作实例
var VirtualMachineService = new(virtualMachineService)
