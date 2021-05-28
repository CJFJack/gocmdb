package controllers

import (
	"fmt"
	"gocmdb/base/controllers/auth"
	"gocmdb/models"
	"gocmdb/services"
)

type UsersController struct {
	auth.LayoutController
}

// 查询用户信息
func (c *UsersController) Query() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code":          0,
			"msg":           "ok",
			"genderTextMap": map[string]string{},
			"tableData":     []*map[string]interface{}{},
			"tableColumns":  []*map[string]interface{}{},
			"tableTotal":    0,
		}
		defer func() {
			c.Data["json"] = result
		}()

		jsonData := struct {
			pagination map[string]interface{}
		}{}
		rawData, err := c.ParsePostForm(jsonData)
		if err != nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("解析Json数据失败：%s", err)
			return
		}
		limit := int(rawData["pagination"].(map[string]interface{})["pageSize"].(float64))
		offset := int(rawData["pagination"].(map[string]interface{})["currentPage"].(float64)-1) * limit

		result["tableData"], result["tableTotal"] = services.UserService.Query("", limit, offset)
		result["tableColumns"] = []map[string]string{
			{"title": "员工ID", "key": "StaffID"},
			{"title": "用户名", "key": "Name"},
			{"title": "昵称", "key": "NickName"},
			{"title": "性别", "key": "Gender"},
			{"title": "电话", "key": "Tel"},
			{"title": "地址", "key": "Addr"},
			{"title": "邮件", "key": "Email"},
			{"title": "部门", "key": "Department"},
			{"title": "状态", "key": "Status"},
			{"title": "创建时间", "key": "CreatedAt"},
			{"title": "更新时间", "key": "UpdatedAt"},
		}
		result["genderTextMap"] = services.UserService.GenderTextMap()
		result["statusTextMap"] = services.UserService.StatusTextMap()
	}
}

// 新增用户
func (c *UsersController) Add() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "ok",
		}
		defer func() {
			c.Data["json"] = result
		}()

		model := models.NewUser()
		err := c.ParseJson(model)
		if err != nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("解析Json数据失败：%s", err)
			return
		}
		err = services.UserService.Add(model)
		if err != nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("插入数据库失败：%s", err)
			return
		}
	}
}

// 修改用户信息
func (c *UsersController) Modify() {
	result := map[string]interface{}{
		"code": 0,
		"msg":  "ok",
	}
	defer func() {
		c.Data["json"] = result
	}()

	model := models.NewUser()
	err := c.ParseJson(model)
	if err != nil {
		result["code"] = 500
		result["msg"] = err
		return
	}
	err = services.UserService.Modify(model)
	if err != nil {
		result["code"] = 500
		result["msg"] = err
		return
	}
}

// 删除用户
func (c *UsersController) Delete() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "ok",
		}
		defer func() {
			c.Data["json"] = result
		}()

		model := models.NewUser()
		err := c.ParseJson(model)
		if err != nil {
			result["code"] = 500
			result["msg"] = err
			return
		}
		err = services.UserService.Delete(model.ID)
		if err != nil {
			result["code"] = 500
			result["msg"] = err
			return
		}
	}
}
