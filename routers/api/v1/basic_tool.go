package v1

import (
	"chatai/pkg/app"
	"chatai/pkg/e"
	"chatai/service/basic_tool_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BasicFunc struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

func GetBasicFunc(c *gin.Context) {
	appG := app.Gin{C: c}
	data, err := basic_tool_service.GetBasicToolInfoList()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetBasicFuncInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	data, err := basic_tool_service.GetToolInfo(name)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func PostBasicFuncInfo(c *gin.Context) {
	type InputParam struct {
		Name string `json:"name" form:"name"`
		InputValue string `json:"inputValue" form:"inputValue"`
	}
	appG := app.Gin{C: c}
	ip := &InputParam{}
	if err := c.ShouldBind(ip); err !=nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	basicToolInfo, err := basic_tool_service.GetToolInfo(ip.Name)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	data, err := basic_tool_service.PostChatGPT(basicToolInfo.Prompt, ip.InputValue)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
