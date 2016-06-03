package storage

import (
	"constant"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"model"
)

var db = NewMysql()

func NewMysql() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", constant.DB_USER, constant.DB_PASSWORD, constant.DB_NAME))
	if err != nil {
		panic(err.Error())
	}
	return db
}

/**
 * 插入某个股票的一天交易数据
 */
func InsertTradeHis(day *model.DayTrade) error {
	sql := `
	INSERT INTO trade_his
	(
		code, 
		date,
		open,
		close,
		high,
		low,
		volume,
		money
	) VALUES (
		?, 
		?,
		?,
		?,
		?,
		?,
		?,
		?
	)`
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(day.Code, day.Date,
		day.Open, day.Close, day.High, day.Low,
		day.Volume, day.Money)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

/**
 * 某只股票最新一天的交易数据
 */
func GetLatestDayStock(stock *model.Stock) *model.DayTrade {
	sql := `
	SELECT 
		code, 
		date,
		open,
		close,
		high,
		low,
		volume,
		money
	FROM trade_his
	WHERE code = ?
	ORDER BY date DESC
	LIMIT 1
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	day := &model.DayTrade{}
	row := stmt.QueryRow(stock.Code)
	if row == nil {
		return day
	}
	err = row.Scan(
		&day.Code,
		&day.Date,
		&day.Open,
		&day.Close,
		&day.High,
		&day.Low,
		&day.Volume,
		&day.Money,
	)
	if err != nil {
		return day
	}
	return day
}
