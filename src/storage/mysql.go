package storage

import (
	"constant"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"model"
)

/**
 	------------------------------------------
	|	table:                               |
	|		trade_his   日K线信息表          |
	|	cols:                                |
	| 		code        |char     |代码      |
	|		date        |date     |日期      |
	|		open        |float    |开盘价    |
	|		close       |float    |收盘价    |
	|		high        |float    |最高价    |
	|		low         |float    |最低价    |
	|		volume      |int      |成交量    |
	|		money       |int      |成交金额  |
	------------------------------------------

	------------------------------------------
	|	table:                               |
	|		stock       个股信息             |
	|	cols:                                |
	| 		code        |char     |代码      |
	|		cn_name     |char     |中文名    |
	------------------------------------------

	------------------------------------------
	|	table:                               |
	|		hq          5/10档行情信息       |
	|	cols:                                |
	| 		code        |char     |代码      |
	------------------------------------------
*/

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
func InsertTradeHis(days []*model.DayTrade) error {
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
	for _, day := range days {
		if day == nil {
			continue
		}
		_, err = stmt.Exec(day.Code, day.Date,
			day.Open, day.Close, day.High, day.Low,
			day.Volume, day.Money)
		if err != nil {
			panic(err.Error())
		}
	}
	return nil
}

func GetTradeHis(stock *model.Stock, begin string, end string) []*model.DayTrade {
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
	AND date >= ?
	AND date <= ?
	ORDER BY DATE ASC
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	days := make([]*model.DayTrade, 0)
	rows, err := stmt.Query(stock.Code, begin, end)
	defer rows.Close()
	if rows == nil || err != nil {
		return days
	}
	for rows.Next() {
		day := &model.DayTrade{}
		err = rows.Scan(
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
			continue
		}
		days = append(days, day)
	}
	return days
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

func GetAllStocks() []*model.Stock {
	sql := `
		SELECT code, cn_name FROM stock
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	rows, _ := stmt.Query()
	if rows == nil {
		return nil
	}
	stocks := make([]*model.Stock, 0)
	for rows.Next() {
		stock := &model.Stock{}
		err = rows.Scan(&stock.Code, &stock.CnName)
		if err != nil {
			return nil
		}
		stock.Type = model.Code2Type(stock.Code)
		stocks = append(stocks, stock)
	}
	return stocks
}

func GetStockInfo(stock *model.Stock) bool {
	sql := `
		SELECT cn_name FROM stock WHERE code = ?
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	row := stmt.QueryRow(stock.Code)
	if row == nil {
		return false
	}
	err = row.Scan(&stock.CnName)
	if err != nil {
		return false
	}
	stock.Type = model.Code2Type(stock.Code)
	return true

}

// 更新股票实体信息
func UpSertStockInfo(stocks []*model.Stock) error {
	sql := `
	INSERT INTO stock
	(
		code,
		cn_name
	)
	VALUES
	(
		?,
		?
	)
	ON DUPLICATE KEY UPDATE
	cn_name = ?
	`
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	for _, stock := range stocks {
		if stock == nil {
			continue
		}
		_, err = stmt.Exec(stock.Code, stock.CnName, stock.CnName)
		if err != nil {
			panic(err.Error())
		}
	}
	return nil
}
