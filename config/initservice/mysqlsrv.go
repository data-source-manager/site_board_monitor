package initservice

import (
	"errors"
	"gorm.io/gorm"
	"site_board_monitor/service/mysqlservice"
)

var (
	NoExist     = errors.New("不存在该数据")
	ServerError = errors.New("网路错误")
)

type (
	Board interface {
		// QuerySiteByUUID 查询站点信息
		QuerySiteByUUID(siteUUID string) (mysqlservice.Site, error)
		// QueryBoardByUUID 查询板块信息
		QueryBoardByUUID(boardUUID string) (mysqlservice.SiteBoard, error)
		// QueryBoardBySiteUUID 根据站点id获取所有的板块信息
		QueryBoardBySiteUUID(siteUUID string) ([]mysqlservice.SiteBoard, error)
		QueryAllSite() ([]mysqlservice.Site, error)
		QueryAlLBoard() ([]mysqlservice.SiteBoard, error)
		UpdateBoardByUUID(boardUUID, error string) error
	}

	board struct {
		db *gorm.DB
	}
)

func NewBoard() Board {
	return &board{db: SqlService}
}

func (b *board) QuerySiteByUUID(siteUUID string) (mysqlservice.Site, error) {
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

func (b *board) QueryBoardByUUID(boardUUID string) (mysqlservice.SiteBoard, error) {
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

func (b *board) QueryBoardBySiteUUID(siteUUID string) ([]mysqlservice.SiteBoard, error) {
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

func (b *board) QueryAllSite() ([]mysqlservice.Site, error) {
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

func (b *board) QueryAlLBoard() ([]mysqlservice.SiteBoard, error) {
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

func (b *board) UpdateBoardByUUID(boardUUID, errmsg string) error {
	var board mysqlservice.SiteBoard
	board.BoardUuid = boardUUID
	b.db.Model(&board).Where("board_uuid=?", boardUUID).Updates(map[string]interface{}{"error_msg": errmsg, "board_status": 0})
	return nil
}
