package body

type PushrecordParam struct {
	Genre   byte   `json:"genre"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Thumb   string `json:"thumb"`
	Adminid int64  `json:"adminid`
	Admin   string `json:"admin"`
}

type SendParam struct {
	Content string `json:"content"`
}
