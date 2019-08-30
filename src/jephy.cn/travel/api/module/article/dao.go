package article

import (
	"fmt"
	"jephy.cn/lib/db"
	"strings"
	"time"
)

const (
	TABLE_NAME = "article"
)

var fields = []string{
	"`id`",
	"`classid`",
	"`title`",
	//"`content`",
	"`ctime`",
	"`mtime`",
}

type ArticleDao struct {
	db db.DB
}

func NewArticleDao(db db.DB) *ArticleDao {
	dao := &ArticleDao{db: db}
	return dao
}

func (self *ArticleDao) Insert(info *ArticleInfo) error {
	now := time.Now().Unix()
	sql := "INSERT INTO `%s`(`classid`,`title`,`content`, `ctime`, `mtime`)" +
		" VALUES (%d, '%s', '%s', %d, %d)"
	sql = fmt.Sprintf(sql, TABLE_NAME, info.ClassId, db.Escape(info.Title), db.Escape(info.Content), now, now)

	return self.db.Exec(sql)
}

func (self *ArticleDao) Update(info *ArticleInfo) error {
	now := time.Now().Unix()
	sql := "UPDATE %s set title='%s', content='%s', mtime=%d where id=%d"
	sql = fmt.Sprintf(sql, TABLE_NAME, info.Title, db.Escape(info.Content), now, info.Id)

	return self.db.Exec(sql)
}

func (self *ArticleDao) Get(id int64) (*ArticleInfo, error) {
	info := new(ArticleInfo)

	sql := "select " + strings.Join(fields, ",") + ",`content` from `%s` where `id`=%d"
	sql = fmt.Sprintf(sql, TABLE_NAME, id)

	err := self.db.Query(info, sql)
	if err != nil {
		if err == db.ErrNoRows {
			//return nil, nil
		}
		return nil, err
	}
	return info, nil
}

func (self *ArticleDao) GetArticlesByClassId(classId int64, offset, limit int64) ([]*ArticleInfo, error) {
	infos := []*ArticleInfo{}

	sql := "select " + strings.Join(fields, ",") + " from `%s` where `classid`=%d order by ctime desc limit %d,%d"
	sql = fmt.Sprintf(sql, TABLE_NAME, classId, offset, limit)

	err := self.db.Query(&infos, sql)
	if err != nil {
		if err == db.ErrNoRows {
			return infos, nil
		}
		return nil, err
	}
	return infos, nil
}
