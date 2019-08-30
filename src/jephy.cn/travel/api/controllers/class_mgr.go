package controllers

import (
	"jephy.cn/travel/api/global"
	"jephy.cn/travel/api/module/class"
	"sync"
)

var (
	classOnce sync.Once
	classMgr  *ClassMgr
)

func GetClassMgr() *ClassMgr {
	classOnce.Do(func() {
		classMgr = &ClassMgr{
			dao: class.NewClassDao(global.Config().DB),
		}
	})
	return classMgr
}

type ClassMgr struct {
	dao *class.ClassDao
}

func (self *ClassMgr) Insert(info *class.ClassInfo) error {
	return self.dao.Insert(info)
}

func (self *ClassMgr) Update(info *class.ClassInfo) error {
	return self.dao.Update(info)
}

func (self *ClassMgr) Get(classId int64) (*class.ClassInfo, error) {
	return self.dao.Get(classId)
}

func (self *ClassMgr) GetParentId(parentId int64, offset, limit int64) ([]*class.ClassInfo, error) {
	return self.dao.GetParentId(parentId, offset, limit)
}
