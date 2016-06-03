package model

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
