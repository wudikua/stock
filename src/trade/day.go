package trade

import (
	"model"
	"net/http"
	"storage"
)

/**
 * 获取沪深A股，创业板个股的日线
 */

type DayTrade struct {
}

func NewDayTrade() *DayTrade {
	return &DayTrade{}
}

/**
 * 对某一只股票进行更新日交易历史
 * @param  {[type]} this *DayTrade)    update(stock *model.Stock [description]
 * @return {[type]}      [description]
 */
func (this *DayTrade) update(stock *model.Stock) {
	day := storage.GetLatestDayStock(stock)
	if day == nil || day.Date == nil {

	}
}
