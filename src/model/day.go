package model

import "fmt"

/**
 * 日线模型
 */

type DayTrade struct {
	Code string
	Date string
	// 开盘价 收盘价
	Open  float32
	Close float32
	// 成交量 成交金额
	Volume int
	Money  int
	// 最高价 最低价
	High float32
	Low  float32
}

func (this *DayTrade) String() string {
	return fmt.Sprintf("%s [%s] 开:%f 收:%f 高:%f 低:%f 量:%d 金:%d\n", this.Date, this.Code,
	this.Open, this.Close, this.High, this.Low, this.Volume, this.Money)
}
