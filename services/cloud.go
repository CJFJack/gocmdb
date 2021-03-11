package services

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"gocmdb/models"
	"time"
)

type cloudService struct {
}

// 查询云平台
func (s *cloudService) Query(q string, limit, offset int) ([]*models.CloudPlatform, int64) {
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
	querySet.Limit(limit).Offset(offset).All(&cloudPlatform, "ID", "Name", "Type", "Addr", "Region", "Remark", "CreatedTime", "SyncTime", "User", "Status")

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

// 新增云平台
func (s *cloudService) Add(model *models.CloudPlatform, user *models.User) error {
	ormer := orm.NewOrm()
	model.User = user
	_, _, err := ormer.ReadOrCreate(model, "Name")
	return err
}

// 修改云平台信息
func (s *cloudService) Modify(model *models.CloudPlatform) error {
	if cloudPlatform := s.GetByPk(model.ID); cloudPlatform != nil {
		cloudPlatform.Name = model.Name
		cloudPlatform.Type = model.Type
		if model.AccessKey != "" {
			cloudPlatform.AccessKey = model.AccessKey
		}
		if model.SecretKey != "" {
			cloudPlatform.SecretKey = model.SecretKey
		}
		cloudPlatform.Addr = model.Addr
		cloudPlatform.Region = model.Region
		cloudPlatform.Remark = model.Remark
		cloudPlatform.Status = model.Status
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

// 用户性别映射
func (s *cloudService) GenderTextMap() map[string]string {
	return map[string]string{
		"1": "男",
		"0": "女",
	}
}

// 用户状态映射
func (s *cloudService) StatusTextMap() map[string]string {
	return map[string]string{
		"0": "正常",
		"1": "锁定",
		"2": "离职",
	}
}

// 用户操作实例
var CloudService = new(cloudService)
