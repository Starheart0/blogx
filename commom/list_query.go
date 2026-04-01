package commom

import (
	"blogx_server/global"
	"fmt"

	"gorm.io/gorm"
)

type PageInfo struct {
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Order string `form:"order"`
}

func (p PageInfo) GetPage() int {
	if p.Page > 20 || p.Page <= 0 {
		return 1
	}
	return p.Page
}

func (p PageInfo) GetLimit() int {
	if p.Limit <= 0 || p.Limit > 100 {
		p.Limit = 10
	}
	return p.Limit
}

func (p PageInfo) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

type Option struct {
	PageInfo     PageInfo
	Likes        []string
	Preloads     []string
	Where        *gorm.DB
	Debug        bool
	Order        string
	DefaultOrder string
}

func ListQuery[T any](model T, option Option) (list []T, count int, err error) {
	query := global.DB.Where(model)

	if option.Debug {
		query = query.Debug()
	}
	for _, preload := range option.Preloads {
		query = query.Preload(preload)
	}

	if len(option.Likes) > 0 && option.PageInfo.Key != "" {
		likes := global.DB.Where("")
		for _, like := range option.Likes {
			likes.Or(fmt.Sprintf("%s like '%s'", like, fmt.Sprintf("%%%s%%", option.PageInfo.Key)))
		}
		query = query.Where(likes)
	}

	if option.Where != nil {
		query = query.Where(option.Where)
	}

	var _c int64
	global.DB.Model(model).Count(&_c)
	count = int(_c)

	if option.PageInfo.Order != "" {
		query = query.Order(option.PageInfo.Order)
	} else if option.DefaultOrder != "" {
		query = query.Order(option.DefaultOrder)
	}

	err = query.Model(&model).Offset(option.PageInfo.GetOffset()).Limit(option.PageInfo.GetLimit()).Find(&list).Error
	return
}
