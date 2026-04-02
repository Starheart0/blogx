package site

type SiteInfo struct {
	Title string `json:"title" yaml:"title"`
	Logo  string `json:"logo" yaml:"logo"`
	Beian string `json:"beian" yaml:"beian"`
	Mode  int8   `json:"mode" yaml:"mode" binding:"oneof=1 2"`
}

type Project struct {
	Title   string `json:"title" yaml:"title"`
	Icon    string `json:"icon" yaml:"icon"`
	WebPath string `json:"webPath" yaml:"webPath"`
}

type Seo struct {
	Keywords    string `json:"keywords" yaml:"keywords"`
	Description string `json:"description" yaml:"description"`
}

type About struct {
	SiteDate string `json:"siteDate" yaml:"siteDate"`
	QQ       string `json:"qq" yaml:"QQ"`
	Version  string `yaml:"-" json:"version"`
	Wechat   string `json:"wechat" yaml:"wechat"`
	Gitee    string `json:"gitee" yaml:"gitee"`
	Bilibili string `json:"bilibili" yaml:"bilibili"`
	Github   string `json:"github" yaml:"github"`
}

type Login struct {
	QQLogin          bool `json:"qqLogin" yaml:"QQLogin"`
	UsernamePwdLogin bool `json:"usernamePwdLogin" yaml:"UsernamePwdLogin"`
	EmailLogin       bool `json:"emailLogin" yaml:"EmailLogin"`
	Captcha          bool `json:"captcha" yaml:"Captcha"`
}

type ComponentInfo struct {
	Title  string `json:"title" yaml:"title"`
	Enable bool   `json:"enable" yaml:"enable"`
}

type IndexRight struct {
	List []ComponentInfo `json:"list" yaml:"list"`
}

type Article struct {
	NoExamine bool `json:"noExamine" yaml:"noExamine"`
}
