package controllers

import (
	"fmt"
	"gocmdb/base/controllers/auth"
	"gocmdb/forms"
	"gocmdb/models"
	"gocmdb/services"
)

type NodeController struct {
	auth.LayoutController
}

// 查询Node节点信息
func (c *NodeController) Query() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code":         0,
			"msg":          "ok",
			"tableData":    []*map[string]interface{}{},
			"tableColumns": []*map[string]interface{}{},
			"tableTotal":   0,
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

		result["tableData"], result["tableTotal"] = services.NodeService.Query("", limit, offset)
		result["tableColumns"] = []map[string]string{
			{"title": "UUID", "key": "UUID"},
			{"title": "主机名", "key": "Hostname"},
			{"title": "地址", "key": "Addr"},
			{"title": "创建时间", "key": "CreatedAt"},
			{"title": "更新时间", "key": "UpdatedAt"},
		}
	}
}

// 新增Node节点
func (c *NodeController) Add() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "ok",
		}
		defer func() {
			c.Data["json"] = result
			c.ServeJSON()
		}()

		model := models.NewNode()
		err := c.ParseJson(model)
		if err != nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("解析Json数据失败：%s", err)
			return
		}
		err = services.NodeService.Add(model)
		if err != nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("插入数据库失败：%s", err)
			return
		}
	}
}

// 删除Node节点
func (c *NodeController) Delete() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "ok",
		}
		defer func() {
			c.Data["json"] = result
			c.ServeJSON()
		}()

		model := models.NewNode()
		err := c.ParseJson(model)
		if err != nil {
			result["code"] = 500
			result["msg"] = err
			return
		}
		err = services.NodeService.Delete(model.ID)
		if err != nil {
			result["code"] = 500
			result["msg"] = err
			return
		}
	}
}

// 修改Node信息
func (c *NodeController) Modify() {
	result := map[string]interface{}{
		"code": 0,
		"msg":  "ok",
	}
	defer func() {
		c.Data["json"] = result
		c.ServeJSON()
	}()

	model := models.NewNode()

	err := c.ParseJson(model)
	if err != nil {
		result["code"] = 500
		result["msg"] = err
		return
	}

	err = services.NodeService.Modify(model)
	if err != nil {
		result["code"] = 500
		result["msg"] = fmt.Sprintf("%s", err)
		return
	}
}

type JobController struct {
	auth.LayoutController
}

// 查询Job信息
func (c *JobController) Query() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code":         0,
			"msg":          "ok",
			"tableData":    []*map[string]interface{}{},
			"tableColumns": []*map[string]interface{}{},
			"tableTotal":   0,
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

		result["tableData"], result["tableTotal"] = services.JobService.Query("", limit, offset)
		result["tableColumns"] = []map[string]string{
			{"title": "节点", "key": "Node"},
			{"title": "Key", "key": "Key"},
			{"title": "备注", "key": "Remark"},
			{"title": "创建时间", "key": "CreatedAt"},
			{"title": "更新时间", "key": "UpdatedAt"},
		}
		result["nodeSelectOptions"], _ = services.NodeService.Query("", 0, 0)
	}
}

// 新增Job
func (c *JobController) Add() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "ok",
		}
		defer func() {
			c.Data["json"] = result
			c.ServeJSON()
		}()

		//model := models.NewJob()
		form := &forms.JobAddForm{}
		err := c.ParseJson(form)
		if err != nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("解析Json数据失败：%s", err)
			return
		}
		err = services.JobService.Add(form)
		if err != nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("插入数据库失败：%s", err)
			return
		}
	}
}

// 删除Job
func (c *JobController) Delete() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "ok",
		}
		defer func() {
			c.Data["json"] = result
			c.ServeJSON()
		}()

		model := models.NewJob()
		err := c.ParseJson(model)
		if err != nil {
			result["code"] = 500
			result["msg"] = err
			return
		}
		err = services.JobService.Delete(model.ID)
		if err != nil {
			result["code"] = 500
			result["msg"] = err
			return
		}
	}
}

// 修改Job
func (c *JobController) Modify() {
	result := map[string]interface{}{
		"code": 0,
		"msg":  "ok",
	}
	defer func() {
		c.Data["json"] = result
		c.ServeJSON()
	}()

	form := &forms.JobModifyForm{}

	err := c.ParseJson(form)
	if err != nil {
		result["code"] = 500
		result["msg"] = err
		return
	}

	err = services.JobService.Modify(form)
	if err != nil {
		result["code"] = 500
		result["msg"] = fmt.Sprintf("%s", err)
		return
	}
}

type TargetController struct {
	auth.LayoutController
}

// 查询Target信息
func (c *TargetController) Query() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code":         0,
			"msg":          "ok",
			"tableData":    []*map[string]interface{}{},
			"tableColumns": []*map[string]interface{}{},
			"tableTotal":   0,
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

		result["tableData"], result["tableTotal"] = services.TargetService.Query("", limit, offset)
		result["tableColumns"] = []map[string]string{
			{"title": "节点", "key": "Node"},
			{"title": "Job", "key": "Job"},
			{"title": "名称", "key": "Name"},
			{"title": "备注", "key": "Remark"},
			{"title": "地址", "key": "Addr"},
			{"title": "创建时间", "key": "CreatedAt"},
			{"title": "更新时间", "key": "UpdatedAt"},
		}
		result["jobSelectOptions"], _ = services.JobService.Query("", 0, 0)
	}
}

// 新增Target
func (c *TargetController) Add() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "ok",
		}
		defer func() {
			c.Data["json"] = result
			c.ServeJSON()
		}()

		//model := models.NewJob()
		form := &forms.TargetAddForm{}
		err := c.ParseJson(form)
		if err != nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("解析Json数据失败：%s", err)
			return
		}
		err = services.TargetService.Add(form)
		if err != nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("插入数据库失败：%s", err)
			return
		}
	}
}

// 删除Target
func (c *TargetController) Delete() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "ok",
		}
		defer func() {
			c.Data["json"] = result
			c.ServeJSON()
		}()

		model := models.NewTarget()
		err := c.ParseJson(model)
		if err != nil {
			result["code"] = 500
			result["msg"] = err
			return
		}
		err = services.TargetService.Delete(model.ID)
		if err != nil {
			result["code"] = 500
			result["msg"] = err
			return
		}
	}
}

// 修改Target
func (c *TargetController) Modify() {
	result := map[string]interface{}{
		"code": 0,
		"msg":  "ok",
	}
	defer func() {
		c.Data["json"] = result
		c.ServeJSON()
	}()

	form := &forms.TargetModifyForm{}

	err := c.ParseJson(form)
	if err != nil {
		result["code"] = 500
		result["msg"] = err
		return
	}

	err = services.TargetService.Modify(form)
	if err != nil {
		result["code"] = 500
		result["msg"] = fmt.Sprintf("%s", err)
		return
	}
}
