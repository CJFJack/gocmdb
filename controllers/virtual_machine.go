package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"gocmdb/base/controllers/auth"
	"gocmdb/cloud"
	"gocmdb/services"
)

type VirtualMachineController struct {
	auth.LayoutController
}

// 查询云主机信息
func (c *VirtualMachineController) Query() {
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

		result["tableData"], result["tableTotal"] = services.VirtualMachineService.Query("", limit, offset)
		result["tableColumns"] = []map[string]string{
			{"title": "平台", "key": "Platform"},
			{"title": "名称", "key": "Name"},
			{"title": "UUID", "key": "UUID"},
			{"title": "配置", "key": "Config"},
			{"title": "操作系统", "key": "OS"},
			{"title": "IP地址", "key": "IPAddr"},
			{"title": "时间", "key": "Time"},
			{"title": "状态", "key": "Status"},
		}
	}
}

// 启动云主机
func (c *VirtualMachineController) Start() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "启动成功",
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
		id := int(rawData["ID"].(float64))
		if vm := services.VirtualMachineService.GetByPk(id); vm == nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("云主机不存在")
			return
		} else {
			if sdk, ok := cloud.DefaultManager.Cloud(vm.Platform.Type); !ok {
				result["code"] = 400
				result["msg"] = "云平台未注册"
				return
			} else {
				sdk.Init(vm.Platform.Addr, vm.Platform.Region, vm.Platform.AccessKey, vm.Platform.SecretKey)
				if err := sdk.StartInstance(vm.UUID); err != nil {
					result["code"] = 400
					result["msg"] = fmt.Sprintf("%s", err)
					return
				}
			}
		}
	}
}

// 关闭云主机
func (c *VirtualMachineController) Stop() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "关闭成功",
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
		id := int(rawData["ID"].(float64))
		if vm := services.VirtualMachineService.GetByPk(id); vm == nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("云主机不存在")
			return
		} else {
			if sdk, ok := cloud.DefaultManager.Cloud(vm.Platform.Type); !ok {
				result["code"] = 400
				result["msg"] = "云平台未注册"
				return
			} else {
				sdk.Init(vm.Platform.Addr, vm.Platform.Region, vm.Platform.AccessKey, vm.Platform.SecretKey)
				if err := sdk.StopInstance(vm.UUID); err != nil {
					result["code"] = 400
					result["msg"] = fmt.Sprintf("%s", err)
					return
				}
			}
		}
	}
}

// 启动云主机
func (c *VirtualMachineController) Reboot() {
	if c.Ctx.Input.IsPost() {
		result := map[string]interface{}{
			"code": 0,
			"msg":  "重启成功",
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
		id := int(rawData["ID"].(float64))
		if vm := services.VirtualMachineService.GetByPk(id); vm == nil {
			result["code"] = 500
			result["msg"] = fmt.Sprintf("云主机不存在")
			return
		} else {
			if sdk, ok := cloud.DefaultManager.Cloud(vm.Platform.Type); !ok {
				beego.Error("云平台未注册")
			} else {
				sdk.Init(vm.Platform.Addr, vm.Platform.Region, vm.Platform.AccessKey, vm.Platform.SecretKey)
				if err := sdk.StopInstance(vm.UUID); err != nil {
					result["code"] = 400
					result["msg"] = fmt.Sprintf("%s", err)
					return
				}
			}
		}
	}
}
