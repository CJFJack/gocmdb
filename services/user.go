package services

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"gocmdb/models"
	"gocmdb/utils"
)

type userService struct {
}

// 查询用户
func (s *userService) Query(q string, limit, offset int) ([]*models.User, int64) {
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

	querySet = querySet.SetCond(cond)
	total, _ := querySet.Count()
	querySet.Limit(limit).Offset(offset).All(&users, "ID", "StaffID", "Name", "NickName", "Password", "Gender", "Tel", "Addr", "Email", "Department", "Status", "CreatedAt", "UpdatedAt")
	return users, total
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
func (s *userService) Add(model *models.User) error {
	ormer := orm.NewOrm()
	model.Password = utils.GeneratePassword(model.Password)
	_, _, err := ormer.ReadOrCreate(model, "Name")
	return err
}

// 修改用户信息
func (s *userService) Modify(model *models.User) error {
	if user := s.GetByPk(model.ID); user != nil {
		user.NickName = model.NickName
		user.Gender = model.Gender
		user.Status = model.Status
		if utils.GeneratePassword(model.Password) != user.Password {
			user.Password = utils.GeneratePassword(model.Password)
		}
		user.Status = model.Status
		ormer := orm.NewOrm()
		_, err := ormer.Update(user, "NickName", "Status", "Gender", "Password")
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

// 删除用户数据
func (s *userService) Delete(pk int) error {
	if user := s.GetByPk(pk); user != nil {
		ormer := orm.NewOrm()
		_, err := ormer.Delete(user)
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

// 用户性别映射
func (s *userService) GenderTextMap() map[string]string {
	return map[string]string{
		"api.conf": "男",
		"0":        "女",
	}
}

// 用户状态映射
func (s *userService) StatusTextMap() map[string]string {
	return map[string]string{
		"0":        "正常",
		"api.conf": "锁定",
		"2":        "离职",
	}
}

// 用户操作实例
var UserService = new(userService)
