package controllers

import (
	"fmt"
	"gocmdb/base/controllers/auth"
	"gocmdb/cloud"
	_ "gocmdb/cloud/plugin"
	"gocmdb/models"
	"gocmdb/services"
)

type CloudPlatformController struct {
	auth.LayoutController
}

// 查询云平台信息
func (c *CloudPlatformController) Query() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code":          0,
			"msg":           "ok",
			"genderTextMap": map[string]string{},
			"tableData": []*map[string]interface{}{},
			"tableColumns": []*map[string]interface{}{},
			"tableTotal": 0,
			"typeOptions": map[string]string{},
			"statusOptions": []*map[string]string{},
		}
		defer func() {
			c.Data["json"] = result
			c.ServeJSON()
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

		result["tableData"], result["tableTotal"] = services.CloudService.Query("", limit, offset, true)
		result["tableColumns"] = []map[string]string{
			{"title": "名称", "key": "Name"},
			{"title": "类型", "key": "Type"},
			{"title": "区域", "key": "Region"},
			{"title": "备注", "key": "Remark"},
			{"title": "创建时间", "key": "CreatedTime"},
			{"title": "最近同步时间", "key": "SyncTime"},
			{"title": "状态", "key": "Status"},
		}
		result["typeOptions"] = cloud.DefaultManager.TypeOptions()
		result["statusOptions"] = services.CloudService.StatusTextMap()
	}
}

// 新增云平台
func (c *CloudPlatformController) Add() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "ok",
		}
		defer func() {
			c.Data["json"] = result
			c.ServeJSON()
		}()

		model := models.NewCloudPlatform()
		err := c.ParseJson(model)
		if err != nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("解析Json数据失败：%s", err)
			return
		}
		username := c.GetRequestUser()
		user := services.UserService.GetByName(username)
		err = services.CloudService.Add(model, user)
		if err != nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("插入数据库失败：%s", err)
			return
		}
	}
}

// 修改云平台信息
func (c *CloudPlatformController) Modify() {
	result := map[string]interface{}{
		"code": 0,
		"msg":  "ok",
	}
	defer func() {
		c.Data["json"] = result
		c.ServeJSON()
	}()

	model := models.NewCloudPlatform()

	err := c.ParseJson(model)
	if err != nil {
		result["code"] = 500
		result["msg"] = err
		return
	}

	err = services.CloudService.Modify(model)
	if err != nil {
		result["code"] = 500
		result["msg"] = fmt.Sprintf("%s", err)
		return
	}
}

// 删除云平台
func (c *CloudPlatformController) Delete() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "ok",
		}
		defer func() {
			c.Data["json"] = result
			c.ServeJSON()
		}()

		model := models.NewCloudPlatform()
		err := c.ParseJson(model)
		if err != nil {
			result["code"] = 500
			result["msg"] = err
			return
		}
		err = services.CloudService.Delete(model.ID)
		if err != nil {
			result["code"] = 500
			result["msg"] = err
			return
		}
	}
}
