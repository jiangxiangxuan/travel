package article

import (
	"github.com/gin-gonic/gin"
	"jephy.cn/lib/log"
	. "jephy.cn/travel/api/controllers"
	. "jephy.cn/travel/api/errors"
	"jephy.cn/travel/api/module/article"
	"strconv"
)

func RegisterURLs(r gin.IRoutes) {
	handler := &articleHandler{}
	r.POST("/api/article/add", handler.Add)
	r.POST("/api/article/update", handler.Update)
	r.GET("/api/article/get", handler.Get)
	r.GET("/api/article/getbyclassid", handler.GetByClassId)
}

type articleHandler struct {
}

func (self *articleHandler) Add(c *gin.Context) {
	info := &article.ArticleInfo{}

	info.ClassId, _ = strconv.ParseInt(c.PostForm("classid"), 10, 64)
	info.Title = c.PostForm("title")
	info.Content = c.PostForm("content")

	err := GetArticleMgr().Insert(info)
	if err != nil {
		log.Error("[articleHandler Insert] err=[%v]", err)
		Fail.Json(c)
		return
	}

	OK.Json(c)
}

func (self *articleHandler) Update(c *gin.Context) {
	info := &article.ArticleInfo{}

	info.Id, _ = strconv.ParseInt(c.PostForm("id"), 10, 64)
	info.ClassId, _ = strconv.ParseInt(c.PostForm("classid"), 10, 64)
	info.Title = c.PostForm("title")
	info.Content = c.PostForm("content")

	err := GetArticleMgr().Update(info)
	if err != nil {
		log.Error("[articleHandler Update] err=[%v]", err)
		Fail.Json(c)
		return
	}

	OK.Json(c)
}

func (self *articleHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)

	info, err := GetArticleMgr().Get(id)
	if err != nil {
		log.Error("[articleHandler Get] err=[%v]", err)
		Fail.Json(c)
		return
	}

	Result(c, info)
}

func (self *articleHandler) GetByClassId(c *gin.Context) {
	classId, _ := strconv.ParseInt(c.Query("classid"), 10, 64)
	offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)

	infos, err := GetArticleMgr().GetArticlesByClassId(classId, offset, limit)
	if err != nil {
		log.Error("[articleHandler GetByClassId] err=[%v]", err)
		Fail.Json(c)
		return
	}

	Result(c, infos)
}
