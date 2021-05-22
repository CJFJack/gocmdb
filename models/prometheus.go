package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Node struct {
	ID        int        `orm:"column(id);"`
	UUID      string     `orm:"column(uuid);varchar(64);unique"`
	Hostname  string     `orm:"varchar(64)"`
	Addr      string     `orm:"varchar(512)"`
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`
	Jobs      []*Job     `orm:"reverse(many)"`
}

type Job struct {
	ID        int        `orm:"column(id);"`
	Key       string     `orm:"varchar(64)"`
	Remark    string     `orm:"varchar(512)"`
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`
	Node      *Node      `orm:"rel(fk)"`
	Targets   []*Target  `orm:"reverse(many)"`
}

type Target struct {
	ID        int        `orm:"column(id);"`
	Name      string     `orm:"varchar(64)"`
	Remark    string     `orm:"varchar(512)"`
	Addr      string     `orm:"varchar(126)"`
	CreatedAt *time.Time `orm:"auto_now_add"`
	UpdatedAt *time.Time `orm:"auto_now"`
	DeletedAt *time.Time `orm:"null"`
	Job       *Job       `orm:"rel(fk)"`
}

func init()  {
	orm.RegisterModel(new(Node), new(Job), new(Target))
}
