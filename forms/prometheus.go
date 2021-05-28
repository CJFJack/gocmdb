package forms

type JobAddForm struct {
	Key    string `form:"key"`
	Remark string `form:"remark"`
	Node   int    `form:"node"`
}

type JobModifyForm struct {
	ID     int    `form:"id"`
	Key    string `form:"key"`
	Remark string `form:"remark"`
	Node   int    `form:"node"`
}

type TargetAddForm struct {
	Name   string `form:"name"`
	Remark string `form:"remark"`
	Addr   string `form:"addr"`
	Job    int    `form:"job"`
}

type TargetModifyForm struct {
	ID     int    `form:"id"`
	Name   string `form:"name"`
	Remark string `form:"remark"`
	Addr   string `form:"addr"`
	Job    int    `form:"job"`
}
