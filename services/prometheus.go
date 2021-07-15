package services

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"gocmdb/forms"
	"gocmdb/models"
	"time"
)

type nodeService struct {
}

// 查询Node节点
func (s *nodeService) Query(q string, limit, offset int) ([]*models.Node, int64) {
	var nodes []*models.Node
	querySet := orm.NewOrm().QueryTable(&models.Node{})

	cond := orm.NewCondition()
	cond = cond.And("deleted_at__isnull", true)

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("uuid__icontains", q)
		query = query.Or("hostname__icontains", q)
		query = query.Or("addr__icontains", q)
		cond.AndCond(query)
	}

	querySet = querySet.SetCond(cond)
	total, _ := querySet.Count()
	querySet.Limit(limit).Offset(offset).All(&nodes, "ID", "UUID", "Hostname", "Addr", "CreatedAt", "UpdatedAt")

	return nodes, total
}

// 根据 id 查询Node节点
func (s *nodeService) GetByPk(pk int) *models.Node {
	node := &models.Node{ID: pk}
	ormer := orm.NewOrm()
	if err := ormer.Read(node); err == nil {
		return node
	}
	return nil
}

// 根据 UUID 查询Node节点
func (s *nodeService) GetByUUID(UUID string) *models.Node {
	node := &models.Node{UUID: UUID}
	ormer := orm.NewOrm()
	if err := ormer.Read(node, "UUID"); err == nil {
		return node
	}
	return nil
}

// 注册Node节点
func (s *nodeService) Register(model *models.Node) error {
	ormer := orm.NewOrm()
	if node := s.GetByUUID(model.UUID); node != nil {
		node.Hostname = model.Hostname
		node.Addr = model.Addr
		node.DeletedAt = nil
		_, err := ormer.Update(node)
		return err
	} else {
		_, err := ormer.Insert(model)
		return err
	}
}

// 逻辑删除Node节点数据
func (s *nodeService) Delete(pk int) error {
	if node := s.GetByPk(pk); node != nil {
		ormer := orm.NewOrm()
		now := time.Now()
		node.DeletedAt = &now
		_, err := ormer.Update(node, "DeletedAt")
		return err
	} else {
		return fmt.Errorf("Node节点不存在")
	}
}

// 修改Node信息
func (s *nodeService) Modify(model *models.Node) error {
	if node := s.GetByPk(model.ID); node != nil {
		node.Hostname = model.Hostname
		node.Addr = model.Addr
		ormer := orm.NewOrm()
		_, err := ormer.Update(node, "Hostname", "Addr")
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

// 更新同步信息
func (s *nodeService) SyncInfo(platform *models.CloudPlatform, now *time.Time, msg string) error {
	platform.SyncTime = now
	platform.Msg = msg
	_, err := orm.NewOrm().Update(platform, "SyncTime", "Msg")
	return err
}

type jobService struct {
}

// 根据 id 查询Job
func (s *jobService) GetByPk(pk int) *models.Job {
	job := &models.Job{ID: pk}
	ormer := orm.NewOrm()
	if err := ormer.Read(job); err == nil {
		// 初始化关联对象的所有属性
		ormer.LoadRelated(job, "Node")
		return job
	}
	return nil
}

// 根据 uuid 查询Target
func (s *jobService) GetByUUID(uuid string) []*models.Job {
	var jobs []*models.Job
	ormer := orm.NewOrm()
	querySet := ormer.QueryTable(&models.Job{})
	querySet.RelatedSel().Filter("deleted_at__isnull", true).Filter("node__uuid", uuid).All(&jobs)
	for _, job := range jobs {
		ormer.LoadRelated(job, "Targets")
	}
	return jobs
}

// 查询Job
func (s *jobService) Query(q string, limit, offset int) ([]*models.Job, int64) {
	var jobs []*models.Job
	querySet := orm.NewOrm().QueryTable(&models.Job{})

	cond := orm.NewCondition()
	cond = cond.And("deleted_at__isnull", true)

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("key__icontains", q)
		query = query.Or("remark__icontains", q)
		query = query.Or("node__hostname__icontains", q)
		query = query.Or("node__addr__icontains", q)
		cond.AndCond(query)
	}

	querySet = querySet.SetCond(cond)
	total, _ := querySet.Count()
	querySet.RelatedSel().Limit(limit).Offset(offset).All(&jobs, "ID", "Node", "Key", "Remark", "CreatedAt", "UpdatedAt")
	return jobs, total
}

// 新增Job
func (s *jobService) Add(form *forms.JobAddForm) error {
	job := &models.Job{
		Key:    form.Key,
		Remark: form.Remark,
		Node:   NodeService.GetByPk(form.Node),
	}
	_, err := orm.NewOrm().Insert(job)
	return err
}

// 修改Job
func (s *jobService) Modify(form *forms.JobModifyForm) error {
	if job := s.GetByPk(form.ID); job != nil {
		job.Key = form.Key
		job.Remark = form.Remark
		job.Node = NodeService.GetByPk(form.Node)
		ormer := orm.NewOrm()
		_, err := ormer.Update(job, "Key", "Remark", "Node")
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

// 逻辑删除Job数据
func (s *jobService) Delete(pk int) error {
	if job := s.GetByPk(pk); job != nil {
		ormer := orm.NewOrm()
		now := time.Now()
		job.DeletedAt = &now
		_, err := ormer.Update(job, "DeletedAt")
		return err
	} else {
		return fmt.Errorf("Job不存在")
	}
}

type targetService struct {
}

// 根据 id 查询Target
func (s *targetService) GetByPk(pk int) *models.Target {
	target := &models.Target{ID: pk}
	ormer := orm.NewOrm()
	if err := ormer.Read(target); err == nil {
		// 初始化关联对象的所有属性
		ormer.LoadRelated(target, "Job")
		return target
	}
	return nil
}

// 查询Target
func (s *targetService) Query(q string, limit, offset int) ([]*models.Target, int64) {
	var targets []*models.Target
	querySet := orm.NewOrm().QueryTable(&models.Target{})

	cond := orm.NewCondition()
	cond = cond.And("deleted_at__isnull", true)

	if q != "" {
		query := orm.NewCondition()
		query = query.Or("name__icontains", q)
		query = query.Or("remark__icontains", q)
		query = query.Or("addr__icontains", q)
		query = query.Or("job__key__icontains", q)
		query = query.Or("job__remark__icontains", q)
		cond.AndCond(query)
	}

	querySet = querySet.SetCond(cond)
	total, _ := querySet.Count()
	querySet.RelatedSel().Limit(limit).Offset(offset).All(&targets, "ID", "Job", "Name", "Remark", "Addr", "CreatedAt", "UpdatedAt")
	return targets, total
}

// 新增Target
func (s *targetService) Add(form *forms.TargetAddForm) error {
	target := &models.Target{
		Name:   form.Name,
		Remark: form.Remark,
		Addr:   form.Addr,
		Job:    JobService.GetByPk(form.Job),
	}
	_, err := orm.NewOrm().Insert(target)
	return err
}

// 修改Target
func (s *targetService) Modify(form *forms.TargetModifyForm) error {
	if target := s.GetByPk(form.ID); target != nil {
		target.Name = form.Name
		target.Remark = form.Remark
		target.Addr = form.Addr
		target.Job = JobService.GetByPk(form.Job)
		ormer := orm.NewOrm()
		_, err := ormer.Update(target, "Name", "Remark", "Addr", "Job")
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

// 逻辑删除Target数据
func (s *targetService) Delete(pk int) error {
	if target := s.GetByPk(pk); target != nil {
		ormer := orm.NewOrm()
		now := time.Now()
		target.DeletedAt = &now
		_, err := ormer.Update(target, "DeletedAt")
		return err
	} else {
		return fmt.Errorf("Target不存在")
	}
}

// 用户操作实例
var (
	NodeService   = new(nodeService)
	JobService    = new(jobService)
	TargetService = new(targetService)
)
