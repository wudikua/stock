package constant

const (
	// 所有股票代码
	ALL_STOCK_API = "http://www.shdjt.com/js/lib/astock.js"
	// 日K线
	DAY_TRADE_API = "http://quotes.money.163.com/service/chddata.html?code=%s&start=%s&end=%s"
	// 分时图 1分钟级别，多只股票逗号分隔
	HQ_TRADE_API      = "https://analyse.leverfun.com/analyse/eom_prices?type=1&stockCode=%s"
	HQ_TRADE_SINA_API = "http://hq.sinajs.cn/?rn=1465362253123&list=%s"
	HQ_TRADE_126_API  = "http://api.money.126.net/data/feed/%s"
	// L2数据
	HQ_TIMLY_API = "https://app.leverfun.com/timelyInfo/timelyOrderForm?stockCode=%S"
	DB_USER      = "root"
	DB_PASSWORD  = "root"
	DB_NAME      = "stock"
)
