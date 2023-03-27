package mysqlservice

import "time"

// SiteBoard  数据源板块
type SiteBoard struct {
	ID             int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	SiteUuid       string    `gorm:"column:site_uuid" db:"site_uuid" json:"site_uuid" form:"site_uuid"`                             //  站点id(属于哪个站点)
	BoardUuid      string    `gorm:"column:board_uuid" db:"board_uuid" json:"board_uuid" form:"board_uuid"`                         //  站点板块id
	BoardName      string    `gorm:"column:board_name" db:"board_name" json:"board_name" form:"board_name"`                         //  站点中文名
	BoardTheme     string    `gorm:"column:board_theme" db:"board_theme" json:"board_theme" form:"board_theme"`                     //  板块主题
	BoardUrl       string    `gorm:"column:board_url" db:"board_url" json:"board_url" form:"board_url"`                             //  板块url
	ReqMethod      string    `gorm:"column:req_method" db:"req_method" json:"req_method" form:"req_method"`                         //  请求方法
	PostUrl        string    `gorm:"column:post_url" db:"post_url" json:"post_url" form:"post_url"`                                 //  post请求的url
	PostData       string    `gorm:"column:post_data" db:"post_data" json:"post_data" form:"post_data"`                             //  post请求data
	Header         string    `gorm:"column:header" db:"header" json:"header" form:"header"`                                         //  请求头
	IfFontPosition int64     `gorm:"column:if_font_position" db:"if_font_position" json:"if_font_position" form:"if_font_position"` //  是否是头条新闻
	BoardStatus    int64     `gorm:"column:board_status" db:"board_status" json:"board_status" form:"board_status"`                 //  板块状态, 1:正常 0 异常
	ErrorMsg       string    `gorm:"column:error_msg" db:"error_msg" json:"error_msg" form:"error_msg"`                             //  异常信息
	InsertTime     time.Time `gorm:"column:insert_time" db:"insert_time" json:"insert_time" form:"insert_time"`                     //  插入时间
	UpdateTime     string    `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`
}

func (SiteBoard) TableName() string {
	return "site_board"
}

type Site struct {
	ID         int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	AppId      int64     `gorm:"column:app_id" db:"app_id" json:"app_id" form:"app_id"`
	SiteUuid   string    `gorm:"column:site_uuid" db:"site_uuid" json:"site_uuid" form:"site_uuid"`
	SiteName   string    `gorm:"column:site_name" db:"site_name" json:"site_name" form:"site_name"` //  站点名
	Domain     string    `gorm:"column:domain" db:"domain" json:"domain" form:"domain"`             //  域名
	Country    string    `gorm:"column:country" db:"country" json:"country" form:"country"`         //  国家
	Language   string    `gorm:"column:language" db:"language" json:"language" form:"language"`     //  语言
	SiteType   string    `gorm:"column:site_type" db:"site_type" json:"site_type" form:"site_type"` //  站点类型，ex:新闻、军事、社交
	InsertTime time.Time `gorm:"column:insert_time" db:"insert_time" json:"insert_time" form:"insert_time"`
	UpdateTime string    `gorm:"column:update_time" db:"update_time" json:"update_time" form:"update_time"`
}

func (Site) TableName() string {
	return "site"
}
