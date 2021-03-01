package services

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"gocmdb/forms"
	"gocmdb/models"
	"gocmdb/utils"
	"strconv"
	"time"
)

type userService struct {
}

// 根据user id 查询用户信息
func (s *userService) GetByPk(pk int) *models.User {
	user := &models.User{ID: pk}
	ormer := orm.NewOrm()
	if err := ormer.Read(user); err == nil {
		return user
	}
	return nil
}

// 新增用户
func (s *userService) Add(rawData map[string]interface{}) error {
	ormer := orm.NewOrm()
	status, _ := strconv.Atoi(rawData["Status"].(string))
	gender, _ := rawData["Gender"].(int)
	if user := s.GetByName(rawData["Name"].(string)); user != nil {
		user.StaffID = rawData["StaffID"].(string)
		user.Deleted = rawData["Deleted"].(int)
		user.Status = status
		user.NickName, _ = rawData["NickName"].(string)
		user.Password = utils.GeneratePassword(rawData["Password"].(string))
		user.Gender = gender
		user.Tel = rawData["Tel"].(string)
		user.Email = rawData["Email"].(string)
		user.Department = rawData["Department"].(string)
		_, err := ormer.Update(user, "Deleted", "Status", "StaffID", "NickName", "Password", "Gender", "Tel", "Email", "Department")
		return err
	}
	user := &models.User{
		StaffID:    rawData["StaffID"].(string),
		Name:       rawData["Name"].(string),
		NickName:   rawData["NickName"].(string),
		Password:   utils.GeneratePassword(rawData["Password"].(string)),
		Gender:     gender,
		Tel:        rawData["Tel"].(string),
		Email:      rawData["Email"].(string),
		Department: rawData["Department"].(string),
		Status:     status,
	}
	_, _, err := ormer.ReadOrCreate(user, "Name")
	return err
}

// 修改用户信息
func (s *userService) Modify(form *forms.UserModifyForm) {
	if user := s.GetByPk(form.ID); user != nil {
		user.Name = form.Name
		ormer := orm.NewOrm()
		ormer.Update(user, "Name")
	}
}

// 删除用户数据
func (s *userService) Delete(pk int) error {
	if user := s.GetByPk(pk); user != nil {
		now := time.Now()
		user.DeletedAt = &now
		user.Deleted = 1
		ormer := orm.NewOrm()
		_, err := ormer.Update(user, "DeletedAt", "Deleted")
		return err
	} else {
		return fmt.Errorf("用户不存在")
	}
}

// 修改用户密码
func (s *userService) ModifyPassword(pk int, password string) {
	if user := s.GetByPk(pk); user != nil {
		user.Password = utils.GeneratePassword(password)
		orm := orm.NewOrm()
		orm.Update(user, "Password")
	}

}

// 通过用户名获取用户指针
func (s *userService) GetByName(name string) *models.User {
	ormer := orm.NewOrm()
	user := &models.User{Name: name}
	if err := ormer.Read(user, "Name"); err == nil {
		return user
	}
	return nil
}

// 查询用户
func (s *userService) Query(q string) []*models.User {
	var users []*models.User
	querySet := orm.NewOrm().QueryTable(&models.User{})
	cond := orm.NewCondition()
	if q != "" {
		cond = cond.Or("name__icontains", q)
		cond = cond.Or("nickname__icontains", q)
		cond = cond.Or("tel__icontains", q)
		cond = cond.Or("addr__icontains", q)
		cond = cond.Or("email__icontains", q)
		cond = cond.Or("department__icontains", q)
	}
	cond = cond.AndNot("deleted__exact", 1)
	querySet = querySet.SetCond(cond)
	querySet.All(&users)
	return users
}

// 用户性别映射
func (s *userService) GenderTextMap() map[string]string {
	return map[string]string{
		"1": "男",
		"0": "女",
	}
}

// 用户状态映射
func (s *userService) StatusTextMap() map[string]string {
	return map[string]string{
		"0": "正常",
		"1": "锁定",
		"2": "离职",
	}
}

// 用户操作实例
var UserService = new(userService)
