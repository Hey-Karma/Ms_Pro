package model

type TaskWorkTime struct {
	Id         int64  `json:"id"`
	TaskCode   string `json:"task_code"`
	MemberCode string `json:"member_code"`
	CreateTime string `json:"create_time"`
	Content    string `json:"content"`
	BeginTime  string `json:"begin_time"`
	Num        int    `json:"num"`
	Code       string `json:"code"`
	Member     Member `json:"member"`
}
