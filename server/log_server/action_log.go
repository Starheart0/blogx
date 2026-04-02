package log_server

import (
	"blogx_server/core"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/models/enum"
	"blogx_server/utils/jwts"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	e "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type ActionLog struct {
	c                  *gin.Context
	level              enum.LogLevelType
	title              string
	requestBody        []byte
	responseBody       []byte
	log                *models.LogModel
	showRequestHeader  bool
	showRequest        bool
	showResponseHeader bool
	showResponse       bool
	itemList           []string
	responseHeader     http.Header
	isMiddleware       bool
}

func (ac *ActionLog) ShowRequest() {
	ac.showRequest = true
}

func (ac *ActionLog) ShowResponse() {
	ac.showResponse = true
}

func (ac *ActionLog) SetTitle(title string) {
	ac.title = title
}

func (ac *ActionLog) SetLevel(level enum.LogLevelType) {
	ac.level = level
}

func (ac *ActionLog) SetLink(label string, href string) {
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_item link\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\"><a href=\"%s\" target=\"_blank\">%s</a></div></div>\n",
		label,
		href, href))
}

func (ac *ActionLog) SetImage(src string) {
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_image\"><img src=\"%s\" alt=\"\"></div>\n", src))
}

func (ac *ActionLog) ShowRequestHeader() {
	ac.showRequestHeader = true
}

func (ac *ActionLog) ShowResponseHeader() {
	ac.showResponseHeader = true
}

func (ac *ActionLog) setItem(label string, value any, logLevelType enum.LogLevelType) {
	var v string
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice:
		byteData, _ := json.Marshal(value)
		v = string(byteData)
	default:
		v = fmt.Sprintf("%v", value)
	}
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_item %s\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\">%s</div></div>\n",
		logLevelType,
		label, v))
}

func (ac *ActionLog) SetItem(label string, value any) {
	ac.setItem(label, value, enum.LogInfoLevel)
}

func (ac *ActionLog) SetItemInfo(label string, value any) {
	ac.setItem(label, value, enum.LogInfoLevel)
}
func (ac *ActionLog) SetItemWarn(label string, value any) {
	ac.setItem(label, value, enum.LogWarnLevel)
}
func (ac *ActionLog) SetItemError(label string, value any) {
	ac.setItem(label, value, enum.LogErrLevel)
}

func (ac *ActionLog) SetError(label string, err error) {
	msg := e.WithStack(err)
	logrus.Errorf("%s %s\n", label, err.Error())
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_error\"><div class=\"line\"><div class=\"label\">%s</div><div class=\"value\">%s</div><div class=\"type\">%s</div></div><div class=\"stack\">%s</div></div>\n",
		label, err, err, msg))
}

func (ac *ActionLog) SetRequest(c *gin.Context) {
	byteData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Errorf(err.Error())
	}
	c.Request.Body = io.NopCloser(bytes.NewReader(byteData))
	ac.requestBody = byteData
}

func (ac *ActionLog) SetResponse(data []byte) {
	ac.responseBody = data
}

func (ac *ActionLog) SetResponseHeader(header http.Header) {
	ac.responseHeader = header
}

func (ac *ActionLog) MiddlewareSave() {
	_saveLog, _ := ac.c.Get("saveLog")
	saveLog, _ := _saveLog.(bool)
	if !saveLog {
		return
	}
	if ac.log == nil {
		ac.isMiddleware = true
		ac.Save()
		return
	}
	// responseHeader
	if ac.showResponseHeader {
		byteData, _ := json.Marshal(ac.responseHeader)
		ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_response_header\"><pre class=\"log_json_body\">%s</pre></div>\n", string(byteData)))
	}
	// set response
	if ac.showResponse {
		ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_response\"><pre class=\"log_json_body\">%s</pre></div>", string(ac.responseBody)))
	}
	ac.Save()
}
func (ac *ActionLog) Save() (id uint) {
	if ac.log != nil {
		newContent := strings.Join(ac.itemList, "\n")
		content := ac.log.Content + "\n" + newContent
		global.DB.Model(ac.log).Updates(map[string]any{
			"content": content,
		})
		return ac.log.ID
	}
	var newItemList []string
	// requestHeader
	if ac.showRequestHeader {
		byteData, _ := json.Marshal(ac.c.Request.Header)
		newItemList = append(newItemList, fmt.Sprintf("<div class=\"log_request_header\"><pre class=\"log_json_body\">%s</pre></div>\n", string(byteData)))
	}
	// set request
	if ac.showRequest {
		newItemList = append(newItemList, fmt.Sprintf("<div class=\"log_request\"><div class=\"log_request_head\"><span class=\"log_request_method %s\">%s</span><span class=\"log_request_path\">%s</span></div><div class=\"log_request_body\"><pre class=\"log_json_body\">%s</pre></div></div>\n",
			strings.ToLower(ac.c.Request.Method),
			ac.c.Request.Method,
			ac.c.Request.URL.String(),
			string(ac.requestBody),
		))
	}
	// middle content
	newItemList = append(newItemList, ac.itemList...)

	if ac.isMiddleware {
		// responseHeader
		if ac.showResponseHeader {
			byteData, _ := json.Marshal(ac.responseHeader)
			newItemList = append(newItemList, fmt.Sprintf("<div class=\"log_response_header\"><pre class=\"log_json_body\">%s</pre></div>\n", string(byteData)))
		}
		// set response
		if ac.showResponse {
			newItemList = append(newItemList, fmt.Sprintf("<div class=\"log_response\"><pre class=\"log_json_body\">%s</pre></div>", string(ac.responseBody)))
		}
	}
	ip := ac.c.ClientIP()
	addr := core.GetIpAddr(ip)
	claims, err := jwts.ParseTokenByGin(ac.c)
	userID := uint(0)
	if err == nil && claims != nil {
		userID = claims.UserID
	}

	log := models.LogModel{
		LogType: enum.ActionLogType,
		Title:   ac.title,
		Content: strings.Join(newItemList, "\n"),
		Level:   ac.level,
		UserID:  userID,
		IP:      ip,
		Addr:    addr,
	}
	err = global.DB.Create(&log).Error
	if err != nil {
		logrus.Errorf("log create err : %s", err)
		return
	}
	ac.log = &log
	ac.itemList = []string{}
	return id
}

func NewActionLogByGin(c *gin.Context) *ActionLog {
	return &ActionLog{
		c: c,
	}
}

func GetLog(c *gin.Context) *ActionLog {
	_log, ok := c.Get("log")
	if !ok {
		return NewActionLogByGin(c)
	}
	log, ok := _log.(*ActionLog)
	if !ok {
		return NewActionLogByGin(c)
	}
	c.Set("saveLog", true)
	return log
}
