package article

type ArticleInfo struct {
	Id      int64  `db:"id"        json:"id,omitempty"`
	ClassId int64  `db:"classid"   json:"classid,omitempty"`
	Title   string `db:"title"     json:"title,omitempty"`
	Content string `db:"content"   json:"content,omitempty"`
	CTime   int64  `db:"ctime"     json:"ctime,omitempty"`
	MTime   int64  `db:"mtime"     json:"mtime,omitempty"`
}
