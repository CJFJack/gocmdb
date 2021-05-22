package tencent

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"gocmdb/cloud"
)

type TencentCloud struct {
	addr       string
	region     string
	accessKey  string
	secretKey  string
	credential *common.Credential
	profile    *profile.ClientProfile
}

func (c *TencentCloud) Type() string {
	return "tencent"
}

func (c *TencentCloud) Name() string {
	return "腾讯云"
}

func (c *TencentCloud) Init(addr, region, accessKey, secretKey string) {
	c.addr = addr
	c.region = region
	c.accessKey = accessKey
	c.secretKey = secretKey
	c.credential = common.NewCredential(c.accessKey, c.secretKey)
	c.profile = profile.NewClientProfile()
	c.profile.HttpProfile.Endpoint = c.addr
}

func (c *TencentCloud) TestConnect() error {
	client, _ := cvm.NewClient(c.credential, c.region, c.profile)
	request := cvm.NewDescribeRegionsRequest()
	_, err := client.DescribeRegions(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

func (c *TencentCloud) GetInstances() []*cloud.Instance {
	var (
		offset int64 = 0
		limit  int64 = 100
		total  int64 = 1
		rt     []*cloud.Instance
	)

	for offset < total {
		var instances []*cloud.Instance
		total, instances = c.getInstanceByOffsetLimit(offset, limit)
		if offset == 0 {
			rt = make([]*cloud.Instance, 0, total)
		}
		rt = append(rt, instances...)
		offset += limit
	}

	return rt
}

func (c *TencentCloud) transformStatus(status string) string {
	sMap := map[string]string{
		"PENDING":       cloud.StatusPending,
		"LAUNCH_FAILED": cloud.StatusLaunchFailed,
		"RUNNING":       cloud.StatusRunning,
		"STOPPED":       cloud.StatusStopped,
		"STARTING":      cloud.StatusStarting,
		"STOPPING":      cloud.StatusStopping,
		"REBOOTING":     cloud.StatusRebooting,
		"SHUTDOWN":      cloud.StatusShutdown,
		"TERMINATING":   cloud.StatusTerminating,
	}
	if rt, ok := sMap[status]; ok {
		return rt
	}
	return cloud.StatusUnknown
}

func (c *TencentCloud) getInstanceByOffsetLimit(offset, limit int64) (int64, []*cloud.Instance) {
	client, err := cvm.NewClient(c.credential, c.region, c.profile)

	request := cvm.NewDescribeInstancesRequest()
	request.Offset = common.Int64Ptr(offset)
	request.Limit = common.Int64Ptr(limit)

	response, err := client.DescribeInstances(request)
	if err != nil {
		return 0, nil
	}

	total := *response.Response.TotalCount

	if total > 0 {
		instances := response.Response.InstanceSet
		rt := make([]*cloud.Instance, len(instances))

		for index, instance := range instances {
			PublicAddrs := make([]string, len(instance.PublicIpAddresses))
			PrivateAddrs := make([]string, len(instance.PrivateIpAddresses))
			for index, addr := range instance.PublicIpAddresses {
				PublicAddrs[index] = *addr
			}
			for index, addr := range instance.PrivateIpAddresses {
				PrivateAddrs[index] = *addr
			}

			rt[index] = &cloud.Instance{
				UUID:         *instance.InstanceId,
				Name:         *instance.InstanceName,
				OS:           *instance.OsName,
				CPU:          int(*(instance.CPU)),
				Mem:          *instance.Memory * 1024,
				PublicAddrs:  PublicAddrs,
				PrivateAddrs: PrivateAddrs,
				Status:       c.transformStatus(*instance.InstanceState),
				CreatedTime:  *instance.CreatedTime,
				ExpiredTime:  *instance.ExpiredTime,
			}
		}
	}
	return 0, nil
}

func (c *TencentCloud) StopInstance(uuid string) error {
	client, _ := cvm.NewClient(c.credential, c.region, c.profile)
	request := cvm.NewStopInstancesRequest()
	request.InstanceIds = common.StringPtrs([]string{uuid})

	_, err := client.StopInstances(request)
	return err
}

func (c *TencentCloud) StartInstance(uuid string) error {
	client, _ := cvm.NewClient(c.credential, c.region, c.profile)
	request := cvm.NewStartInstancesRequest()
	request.InstanceIds = common.StringPtrs([]string{uuid})

	_, err := client.StartInstances(request)
	beego.Info("err", err)
	return err
}

func (c *TencentCloud) RebootInstance(uuid string) error {
	client, _ := cvm.NewClient(c.credential, c.region, c.profile)
	request := cvm.NewRebootInstancesRequest()
	request.InstanceIds = common.StringPtrs([]string{uuid})

	_, err := client.RebootInstances(request)
	return err
}

func init() {
	cloud.DefaultManager.Register(new(TencentCloud))
}
