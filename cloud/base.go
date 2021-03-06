package cloud

const (
	StatusPending = "创建中"
	StatusLaunchFailed = "创建失败"
	StatusRunning = "运行中"
	StatusStopped = "已停止"
	StatusStarting = "开机中"
	StatusStopping = "关机中"
	StatusRebooting = "重启中"
	StatusTerminating = "销毁中"
	StatusShutdown = "停止待销毁"
	StatusUnknown = "未知"
)

type Instance struct {
	UUID         string
	Name         string
	OS           string
	CPU          int
	Mem          int64
	PublicAddrs  []string
	PrivateAddrs []string
	Status       string
	CreatedTime  string
	ExpiredTime  string
}

type ICloud interface {
	Type() string
	Name() string
	Init(string, string, string, string)
	TestConnect() error
	GetInstances() []*Instance
	StopInstance(string) error
	StartInstance(string) error
	RebootInstance(string) error
}
