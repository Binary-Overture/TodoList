package types

type CreateTaskReq struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Status  int    `form:"status" json:"status"`
	Content string `form:"content" json:"content" binding:"max=1000"`
}
