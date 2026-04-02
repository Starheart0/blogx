package image_api

import (
	"blogx_server/commom/res"
	"blogx_server/global"
	"blogx_server/models"
	"blogx_server/utils"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (ImageApi) ImageUploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	// size limit
	if fileHeader.Size > global.Config.Upload.Size*1024*1024 {
		res.FailWithMsg(fmt.Sprintf("file size greater than %dMB", global.Config.Upload.Size), c)
		return
	}

	// suffix limit
	filename := fileHeader.Filename
	suffix, err := imageSuffixJudge(filename)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	// hash limit
	file, err := fileHeader.Open()
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	byteData, _ := io.ReadAll(file)
	hash := utils.Md5(byteData)
	var model models.ImageModel
	err = global.DB.Take(&model, "hash = ?", hash).Error
	if err == nil {
		logrus.Infof("repeat upload same image %s <==> %s %s", filename, model.Filename, hash)
		res.Ok(model.WebPath(), "repeat upload same image", c)
		return
	}

	//upload
	filePath := fmt.Sprintf("%s/%s.%s", global.Config.Upload.UploadDir, hash, suffix)
	model = models.ImageModel{
		Filename: filename,
		Path:     filePath,
		Size:     fileHeader.Size,
		Hash:     hash,
	}
	err = global.DB.Create(&model).Error
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	err = c.SaveUploadedFile(fileHeader, filePath)
	if err != nil {
		logrus.Fatalf("file upload error %s", err)
		res.FailWithError(err, c)
		return
	}
	res.Ok(model.WebPath(), "image upload successfully", c)
}

func imageSuffixJudge(filename string) (suffix string, err error) {
	_list := strings.Split(filename, ".")
	if len(_list) == 1 {
		return "", errors.New("wrong filename")
	}
	suffix = _list[len(_list)-1]
	if !utils.InList(suffix, global.Config.Upload.WhiteList) {
		return "", errors.New("wrong filename")
	}
	return
}
