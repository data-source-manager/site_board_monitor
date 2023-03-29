package initservice

import (
	"errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"site_board_monitor/service/mysqlservice"
)

var (
	NoExist     = errors.New("不存在该数据")
	ServerError = errors.New("网路错误")
)

type (
	SiteBoard interface {
		// QuerySiteByUUID 查询站点信息
		QuerySiteByUUID(siteUUID string) (mysqlservice.Site, error)
		// QueryBoardByUUID 查询板块信息
		QueryBoardByUUID(boardUUID string) (mysqlservice.SiteBoard, error)
		// QueryBoardBySiteUUID 根据站点id获取所有的板块信息
		QueryBoardBySiteUUID(siteUUID string) ([]mysqlservice.SiteBoard, error)
		QueryAllSite() ([]mysqlservice.Site, error)
		QueryAlLBoard() ([]mysqlservice.SiteBoard, error)
		QueryPushBoard() []mysqlservice.PushSiteBoard
		UpdateBoardByUUID(boardUUID, error string) error
	}

	Board struct {
		db  *gorm.DB
		rdb *redis.Client
	}
)

func NewBoard() SiteBoard {
	return &Board{db: SqlService, rdb: Rdb}
}

func (b *Board) QuerySiteByUUID(siteUUID string) (mysqlservice.Site, error) {
	var site mysqlservice.Site
	qRes := b.db.Where("site_uuid=?", siteUUID).Find(&site)
	if qRes.Error != nil {
		return mysqlservice.Site{}, ServerError
	}
	if qRes.RowsAffected != 1 {
		return mysqlservice.Site{}, NoExist
	}

	return site, nil
}

func (b *Board) QueryBoardByUUID(boardUUID string) (mysqlservice.SiteBoard, error) {
	var siteBoard mysqlservice.SiteBoard
	qRes := b.db.Where("board_uuid=?", boardUUID).Find(&siteBoard)
	if qRes.Error != nil {
		return mysqlservice.SiteBoard{}, ServerError
	}
	if qRes.RowsAffected != 1 {
		return mysqlservice.SiteBoard{}, NoExist
	}

	return siteBoard, nil

}

func (b *Board) QueryBoardBySiteUUID(siteUUID string) ([]mysqlservice.SiteBoard, error) {
	var boards []mysqlservice.SiteBoard
	qRes := b.db.Where("site_uuid=?", siteUUID).Find(&boards)
	if qRes.Error != nil {
		return nil, ServerError
	}
	if qRes.RowsAffected != 1 {
		return nil, NoExist
	}

	return boards, nil
}

func (b *Board) QueryAllSite() ([]mysqlservice.Site, error) {
	var sites []mysqlservice.Site
	qRes := b.db.Find(&sites)
	if qRes.Error != nil {
		return nil, ServerError
	}
	if qRes.RowsAffected != 1 {
		return nil, NoExist
	}

	return sites, nil
}

func (b *Board) QueryAlLBoard() ([]mysqlservice.SiteBoard, error) {
	var boards []mysqlservice.SiteBoard
	qRes := b.db.Find(&boards)
	if qRes.Error != nil {
		return nil, ServerError
	}
	if qRes.RowsAffected != 1 {
		return nil, NoExist
	}

	return boards, nil
}

func (b *Board) UpdateBoardByUUID(boardUUID, errmsg string) error {
	var board mysqlservice.SiteBoard
	board.BoardUuid = boardUUID
	b.db.Model(&board).Where("board_uuid=?", boardUUID).Updates(map[string]interface{}{"error_msg": errmsg, "board_status": 0})
	return nil
}

func (b *Board) QueryPushBoard() []mysqlservice.PushSiteBoard {
	var allPushBoardInfo []mysqlservice.PushSiteBoard

	b.db.Raw("select site.app_id,site.country,site.site_name,site_board.board_uuid,site_board.board_name,site_board.board_theme,site_board.if_font_position,site_board.post_data,site_board.post_url,site_board.header from site_board ,site where site_board.site_uuid=site.site_uuid").Scan(&allPushBoardInfo)

	return allPushBoardInfo
}

// 6251 8008 8196 5103
