package class

type ClassInfo struct {
	ClassId  int64  `db:"classid"  json:"classid,omitempty"`
	ParentId int64  `db:"parentid" json:"parentid,omitempty"`
	Name     string `db:"name"     json:"name,omitempty"`
	Des      string `db:"des"      json:"des,omitempty"`
	CTime    int64  `db:"ctime"    json:"ctime,omitempty"`
	MTime    int64  `db:"mtime"    json:"mtime,omitempty"`
}
