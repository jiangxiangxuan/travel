package controllers

import (
	"jephy.cn/travel/api/global"
	"jephy.cn/travel/api/module/article"
	"sync"
)

var (
	articleOnce sync.Once
	articleMgr  *ArticleMgr
)

func GetArticleMgr() *ArticleMgr {
	articleOnce.Do(func() {
		articleMgr = &ArticleMgr{
			dao: article.NewArticleDao(global.Config().DB),
		}
	})
	return articleMgr
}

type ArticleMgr struct {
	dao *article.ArticleDao
}

func (self *ArticleMgr) Insert(info *article.ArticleInfo) error {
	return self.dao.Insert(info)
}

func (self *ArticleMgr) Update(info *article.ArticleInfo) error {
	return self.dao.Update(info)
}

func (self *ArticleMgr) Get(id int64) (*article.ArticleInfo, error) {
	return self.dao.Get(id)
}

func (self *ArticleMgr) GetArticlesByClassId(classId int64, offset, limit int64) ([]*article.ArticleInfo, error) {
	return self.dao.GetArticlesByClassId(classId, offset, limit)
}
