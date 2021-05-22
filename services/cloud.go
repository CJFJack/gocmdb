package services

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"gocmdb/cloud"
	"gocmdb/models"
	"time"
)

type cloudService struct {
}

// 查询云平台
func (s *cloudService) Query(q string, limit, offset int, hidePass bool) ([]*models.CloudPlatform, int64) {
	var cloudPlatform []*models.CloudPlatform
	querySet := orm.NewOrm().QueryTable(&models.CloudPlatform{})

	cond := orm.NewCondition()
	cond = cond.And("deleted_time__isnull", true)

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("addr__icontains", q)
		query = query.Or("remark__icontains", q)
		query = query.Or("region__icontains", q)
		cond.AndCond(query)
	}

	querySet = querySet.SetCond(cond)
	total, _ := querySet.Count()
	if hidePass {
		querySet.Limit(limit).Offset(offset).All(&cloudPlatform, "ID", "Name", "Type", "Addr", "Region", "Remark", "CreatedTime", "SyncTime", "User", "Status")
	} else {
		querySet.Limit(limit).Offset(offset).All(&cloudPlatform, "ID", "Name", "Type", "Addr", "Region", "Remark", "CreatedTime", "SyncTime", "User", "Status", "AccessKey", "SecretKey")
	}


	return cloudPlatform, total
}

// 根据 id 查询云平台
func (s *cloudService) GetByPk(pk int) *models.CloudPlatform {
	cloudPlatform := &models.CloudPlatform{ID: pk}
	ormer := orm.NewOrm()
	if err := ormer.Read(cloudPlatform); err == nil {
		return cloudPlatform
	}
	return nil
}

// 检验参数配置
func (s *cloudService) Valid(model *models.CloudPlatform) error {
	beego.Info(model.Type)
 	if sdk, ok := cloud.DefaultManager.Cloud(model.Type); !ok {
 		return fmt.Errorf("类型错误")
	} else {
		sdk.Init(model.Addr, model.Region, model.AccessKey, model.SecretKey)
		if err := sdk.TestConnect(); err != nil {
			return fmt.Errorf("配置参数错误，请检查AccessKey/SecretKey/地域、地址是否正确")
		}
	}
	return nil
}

// 新增云平台
func (s *cloudService) Add(model *models.CloudPlatform, user *models.User) error {
	err := s.Valid(model)
	if err != nil {
		return err
	}
	ormer := orm.NewOrm()
	model.User = user
	beego.Info(fmt.Sprintf("%#v", model))
	_, _, err = ormer.ReadOrCreate(model, "Name")
	return err
}

// 修改云平台信息
func (s *cloudService) Modify(model *models.CloudPlatform) error {
	if cloudPlatform := s.GetByPk(model.ID); cloudPlatform != nil {
		cloudPlatform.Name = model.Name
		cloudPlatform.Type = model.Type
		if model.AccessKey != "" {
			cloudPlatform.AccessKey = model.AccessKey
		} else {
			model.AccessKey = cloudPlatform.AccessKey
		}
		if model.SecretKey != "" {
			cloudPlatform.SecretKey = model.SecretKey
		} else {
			model.SecretKey = cloudPlatform.SecretKey
		}
		cloudPlatform.Addr = model.Addr
		cloudPlatform.Region = model.Region
		cloudPlatform.Remark = model.Remark
		cloudPlatform.Status = model.Status
		if err := s.Valid(model); err != nil {
			return err
		}
		ormer := orm.NewOrm()
		_, err := ormer.Update(cloudPlatform, "Name", "Type", "AccessKey", "Addr", "SecretKey", "Region", "Remark", "Status")
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

// 逻辑删除云平台数据
func (s *cloudService) Delete(pk int) error {
	if cloudPlatform := s.GetByPk(pk); cloudPlatform != nil {
		ormer := orm.NewOrm()
		now := time.Now()
		cloudPlatform.DeletedTime = &now
		_, err := ormer.Update(cloudPlatform, "DeletedTime")
		// 设置关联云主机的删除时间
		ormer.QueryTable(&models.VirtualMachine{}).Filter("Platform__exact", cloudPlatform).Update(orm.Params{"DeletedTime": &now})
		return err
	} else {
		return fmt.Errorf("云平台不存在")
	}
}

// 通过用户名获取用户指针
func (s *cloudService) GetByName(name string) *models.User {
	ormer := orm.NewOrm()
	user := &models.User{Name: name}
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

// 状态映射
func (s *cloudService) StatusTextMap() map[string]string {
	return map[string]string{
		"0": "启用",
		"1": "禁用",
	}
}

// 更新同步信息
func(s *cloudService) SyncInfo(platform *models.CloudPlatform, now *time.Time, msg string) error {
	platform.SyncTime = now
	platform.Msg = msg
	_, err := orm.NewOrm().Update(platform, "SyncTime", "Msg")
	return err
}

// 用户操作实例
var CloudService = new(cloudService)
