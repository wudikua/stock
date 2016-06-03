package storage

import (
	"fmt"
	"model"
	"testing"
)

func Test_InsertDay(t *testing.T) {
	day := &model.DayTrade{
		Code:   "123",
		Date:   "2016-06-02",
		Open:   1.0,
		Close:  1.0,
		Volume: 1,
		Money:  2,
		High:   1.0,
		Low:    1.0,
	}
	InsertTradeHis(day)
	t.Log("insert success")
}

func Test_GetDay(t *testing.T) {
	stock := &model.Stock{
		Code: "33",
	}
	day := GetLatestDayStock(stock)
	fmt.Println(day)
}
