package class

import (
	"fmt"
	"jephy.cn/lib/db"
	"strings"
	"time"
)

const (
	TABLE_NAME = "class"
)

var fields = []string{
	"`classid`",
	"`parentid`",
	"`name`",
	"`des`",
	"`ctime`",
	"`mtime`",
}

type ClassDao struct {
	db db.DB
}

func NewClassDao(db db.DB) *ClassDao {
	dao := &ClassDao{db: db}
	return dao
}

func (self *ClassDao) Insert(info *ClassInfo) error {
	now := time.Now().Unix()
	sql := "INSERT INTO `%s`(`parentid`,`name`,`des`, `ctime`, `mtime`)" +
		" VALUES (%d, '%s', '%s', %d, %d)"
	sql = fmt.Sprintf(sql, TABLE_NAME, info.ParentId, db.Escape(info.Name), db.Escape(info.Des), now, now)

	return self.db.Exec(sql)
}

func (self *ClassDao) Update(info *ClassInfo) error {
	now := time.Now().Unix()
	sql := "UPDATE %s set parentid=%d, name='%s', des='%s', mtime=%d where classid=%d"
	sql = fmt.Sprintf(sql, TABLE_NAME, info.ParentId, db.Escape(info.Name), db.Escape(info.Des), now, info.ClassId)

	return self.db.Exec(sql)
}

func (self *ClassDao) Get(classId int64) (*ClassInfo, error) {
	info := new(ClassInfo)

	sql := "select " + strings.Join(fields, ",") + " from `%s` where `classid`=%d"
	sql = fmt.Sprintf(sql, TABLE_NAME, classId)

	err := self.db.Query(info, sql)
	if err != nil {
		if err == db.ErrNoRows {
			//return nil, nil
		}
		return nil, err
	}
	return info, nil
}

func (self *ClassDao) GetParentId(parentId int64, offset, limit int64) ([]*ClassInfo, error) {
	infos := []*ClassInfo{}

	sql := "select " + strings.Join(fields, ",") + " from `%s` where `parentid`=%d order by ctime desc limit %d,%d"
	sql = fmt.Sprintf(sql, TABLE_NAME, parentId, offset, limit)

	err := self.db.Query(&infos, sql)
	if err != nil {
		if err == db.ErrNoRows {
			return infos, nil
		}
		return nil, err
	}
	return infos, nil
}
