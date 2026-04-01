package log_api

import (
	"blogx_server/commom"
	"blogx_server/commom/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/server/log_server"
	"fmt"

	"github.com/gin-gonic/gin"
)

type LogApi struct {
}

type LogListRequest struct {
	commom.PageInfo
	LogType     enum.LogType      `from:"logType"`
	Level       enum.LogLevelType `from:"level"`
	UserID      uint              `from:"userID"`
	IP          string            `from:"IP"`
	LoginStatus bool              `from:"loginStatus"`
	ServiceName string            `from:"serviceName"`
}

type LogListResponse struct {
	models.LogModel
	UserNickname string `json:"userNickname"`
}

func (LogApi) LogListView(c *gin.Context) {
	var cr LogListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	list, count, err := commom.ListQuery(models.LogModel{
		LogType:     cr.LogType,
		Level:       cr.Level,
		UserID:      cr.UserID,
		IP:          cr.IP,
		LoginStatus: cr.LoginStatus,
		ServiceName: cr.ServiceName,
	}, commom.Option{
		PageInfo:     cr.PageInfo,
		Likes:        []string{"Title"},
		Preloads:     []string{"UserModel"},
		DefaultOrder: "created_at desc",
	})

	var _list = make([]LogListResponse, 0)
	for _, logModel := range list {
		_list = append(_list, LogListResponse{
			LogModel:     logModel,
			UserNickname: logModel.UserModel.Nickname,
		})
	}
	res.OkWithList(_list, int(count), c)
}

func (LogApi) LogReadView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var log models.LogModel
	err = global.DB.Take(&log, cr.ID).Error
	if err != nil {
		res.FailWithMsg("not exist", c)
		return
	}
	if !log.IsRead {
		global.DB.Model(&log).Update("is_Read", true)
		res.OkWithMsg("log read successfully", c)
	} else {
		res.FailWithMsg("already read before", c)
	}
}

func (LogApi) LogRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	log := log_server.GetLog(c)
	log.ShowRequest()
	log.ShowResponse()

	var logList []models.LogM odel
	global.DB.Find(&logList, "id in ?", cr.IDlist)

	if len(logList) > 0 {
		global.DB.Delete(&logList)
	}
	msg := fmt.Sprintf("log delete successful, total %d logs", len(logList))
	res.OkWithMsg(msg, c)
}
