package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"gocmdb/cloud"
	_ "gocmdb/cloud/plugin"
	"gocmdb/services"
	"time"
)

func main() {
	for now := range time.Tick(10 * time.Second) {
		beego.Info(now)
		platforms, _ := services.CloudService.Query("", 0, 0, false)
		for _, platform := range platforms {
			msg := ""
			if platform.Status == 1 {
				continue
			}
			if sdk, ok := cloud.DefaultManager.Cloud(platform.Type); !ok {
				beego.Error("云平台未注册")
			} else {
				sdk.Init(platform.Addr, platform.Region, platform.AccessKey, platform.SecretKey)
				if err := sdk.TestConnect(); err != nil {
					msg = fmt.Sprintf("测试连接失败：%s", err)
					beego.Error(msg)
				} else {
					//for _, instance := range sdk.GetInstances() {
					//	services.VirtualMachineService.SyncInstances(instance, platform)
					//}
					instance := cloud.Instance{
						UUID:         "ins-xxxxxxxxx",
						Name:         "test",
						OS:           "Centos 7",
						CPU:          1,
						Mem:          2048,
						PublicAddrs:  []string{"8.45.157.155",},
						PrivateAddrs: []string{"172.16.1.1",},
						Status:       "RUNNING",
						CreatedTime:  "2021-04-02 15:22:00",
						ExpiredTime:  "2099-04-02 15:22:00",
					}

					services.VirtualMachineService.SyncInstances(&instance, platform)
					services.VirtualMachineService.SyncInstanceStatus(now, platform)
				}
			}
			services.CloudService.SyncInfo(platform, &now, msg)
		}
	}
}
