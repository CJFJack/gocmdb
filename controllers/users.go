package controllers

import (
	"github.com/astaxie/beego"
	"gocmdb/base/controllers/auth"
	"gocmdb/forms"
	"gocmdb/services"
	"net/http"
)

type UsersController struct {
	auth.LayoutController
}

// 查询用户信息
func (c *UsersController) Query() {
	result := map[string]interface{}{
		"code":         0,
		"msg":          "ok",
		"username":     "",
		"tableData":    map[string]interface{}{},
		"tableColumns": map[string]interface{}{},
	}
	q := c.GetString("q")
	result["tableData"] = services.UserService.Query(q)
	result["tableColumns"] = []map[string]string{
		{"title": "员工ID", "key": "StaffID"},
		{"title": "用户名", "key": "Name"},
		{"title": "昵称", "key": "NickName"},
		{"title": "性别", "key": "Gender"},
		{"title": "电话", "key": "Tel"},
		{"title": "邮件", "key": "Email"},
		{"title": "部门", "key": "Department"},
		{"title": "状态", "key": "Status"},
		{"title": "创建时间", "key": "CreatedAt"},
		{"title": "更新时间", "key": "UpdatedAt"},
	}
	c.Data["json"] = result
	c.ServeJSON()
}

// 新增用户
func (c *UsersController) Add() {
	if c.Ctx.Input.IsPost() {
		form := &forms.UserAddForm{}
		if err := c.ParseForm(form); err == nil {
			services.UserService.Add(form)
			c.Redirect(beego.URLFor("UsersController.Query"), http.StatusFound)
		}
	}
	c.TplName = "user/add.html"
}

// 修改用户信息
func (c *UsersController) Modify() {
	//c.Abort("NotPermission")
	//return

	form := &forms.UserModifyForm{}
	// GET 获取参数
	// POST 修改用户
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			services.UserService.Modify(form)
			flash := beego.NewFlash()
			flash.Set("notice", "修改用户信息成功")
			flash.Store(&c.AuthorizationController.BaseController.Controller)
			c.Redirect(beego.URLFor("UsersController.Query"), http.StatusFound)
		}
	} else if pk, err := c.GetInt("pk"); err == nil {
		if user := services.UserService.GetByPk(pk); user != nil {
			form.ID = user.ID
			form.Name = user.Name
		}
	}
	c.Data["form"] = form
	c.Data["title"] = "用户编辑"
	c.TplName = "user/modify.html"
}

// 删除用户
func (c *UsersController) Delete() {
	if pk, err := c.GetInt("pk"); err == nil && c.LoginUser.ID != pk {
		services.UserService.Delete(pk)
	}
	c.Redirect(beego.URLFor("UsersController.Query"), http.StatusFound)
}
