package enum

type ArticleStatus int8

const (
	ArticleStatusDraft     ArticleStatus = 1
	ArticleStatusExamine   ArticleStatus = 2
	ArticleStatusPublished ArticleStatus = 3
	ArticleStatusFail      ArticleStatus = 4
)
