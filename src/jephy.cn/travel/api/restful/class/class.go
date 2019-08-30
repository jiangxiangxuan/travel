package class

import (
	"github.com/gin-gonic/gin"
	"jephy.cn/lib/log"
	. "jephy.cn/travel/api/controllers"
	. "jephy.cn/travel/api/errors"
	"jephy.cn/travel/api/module/class"
	"strconv"
)

func RegisterURLs(r gin.IRoutes) {
	handler := &classHandler{}
	r.POST("/api/class/add", handler.Add)
	r.POST("/api/class/update", handler.Update)
	r.GET("/api/class/get", handler.Get)
	r.GET("/api/class/getbyparentid", handler.GetByParentId)
}

type classHandler struct {
}

func (self *classHandler) Add(c *gin.Context) {
	info := &class.ClassInfo{}

	info.ParentId, _ = strconv.ParseInt(c.PostForm("parentid"), 10, 64)
	info.Name = c.PostForm("name")
	info.Des = c.PostForm("des")

	err := GetClassMgr().Insert(info)
	if err != nil {
		log.Error("[classHandler Get] err=[%v]", err)
		Fail.Json(c)
		return
	}

	OK.Json(c)
}

func (self *classHandler) Update(c *gin.Context) {
	info := &class.ClassInfo{}

	info.ClassId, _ = strconv.ParseInt(c.PostForm("classid"), 10, 64)
	info.ParentId, _ = strconv.ParseInt(c.PostForm("parentid"), 10, 64)
	info.Name = c.PostForm("name")
	info.Des = c.PostForm("des")

	err := GetClassMgr().Update(info)
	if err != nil {
		log.Error("[classHandler Get] err=[%v]", err)
		Fail.Json(c)
		return
	}

	OK.Json(c)
}

func (self *classHandler) Get(c *gin.Context) {
	classId, _ := strconv.ParseInt(c.Query("classid"), 10, 64)

	info, err := GetClassMgr().Get(classId)
	if err != nil {
		log.Error("[classHandler Get] err=[%v]", err)
		Fail.Json(c)
		return
	}

	Result(c, info)
}

func (self *classHandler) GetByParentId(c *gin.Context) {
	parentId, _ := strconv.ParseInt(c.Query("parentid"), 10, 64)
	offset, _ := strconv.ParseInt(c.Query("offset"), 10, 64)
	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)

	infos, err := GetClassMgr().GetParentId(parentId, offset, limit)
	if err != nil {
		log.Error("[classHandler GetByParentId] err=[%v]", err)
		Fail.Json(c)
		return
	}

	Result(c, infos)
}
