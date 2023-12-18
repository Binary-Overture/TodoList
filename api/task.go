package api

import (
	"Todolist/pkg/util"
	"Todolist/service"
	"Todolist/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateTaskHandler @Tags TASK
// @Summary 创建任务
// @Produce json
// @Accept json
// @Header 200 {string} Authorization "必备"
// @Param data body service.CreateTaskService true  "title"
// @Success 200 {object} serializer.ResponseTask "{"success":true,"data":{},"msg":"ok"}"
// @Failure 500 {json} {"status":500,"data":{},"Msg":{},"Error":"error"}
// @Router /task [post]
func CreateTaskHandler(c *gin.Context) {
	var req types.CreateTaskReq
	if err := c.ShouldBind(&req); err == nil {
		//参数校验
		l := service.GetTaskSrv()
		resp, err := l.CreateTask(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse(err))
			return
		}
		c.JSON(http.StatusOK, resp)
	} else {
		util.LogrusObj.Infoln(err)
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
	}
}
