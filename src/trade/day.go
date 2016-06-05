package trade

import (
	"model"
	"net/http"
	"constant"
	"time"
	"storage"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
	"log"
	"io"
	"encoding/csv"
)

/**
 * 获取沪深A股，创业板个股的日线
 */

type DayTradeHelper struct {
}

func NewDayTradeHelper() *DayTradeHelper {
	return &DayTradeHelper{}
}

func (this *DayTradeHelper) getTrade(stock *model.Stock, begin string, end string) []*model.DayTrade {
	var url string
	if (stock.Type == model.HU_A) {
		url = fmt.Sprintf(constant.DAY_TRADE_API, "0"+stock.Code, begin, end)
	} else if (stock.Type == model.SHEN_A || stock.Type == model.CHUANGYE || stock.Type == model.ZHONG_XIAO) {
		url = fmt.Sprintf(constant.DAY_TRADE_API, "1"+stock.Code, begin, end)
	} else {
		return nil
	}
	
	log.Println("get trade", stock.Code, "from", begin, "to", end, url)
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	enc := mahonia.NewDecoder("gbk")
	_, utf8Body, _ := enc.Translate(body, true)

	days := make([]*model.DayTrade, 0)
	r := csv.NewReader(strings.NewReader(string(utf8Body)))
	// skip title
	r.Read()
	for {
		cols, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil
		}
		Close, _ := strconv.ParseFloat(cols[3], 32)
		Open, _ := strconv.ParseFloat(cols[6], 32)
		Low, _ := strconv.ParseFloat(cols[5], 32)
		High, _ := strconv.ParseFloat(cols[4], 32)
		Volume, _ := strconv.Atoi(cols[11])
		Money, _ := strconv.ParseFloat(cols[12], 32)
		if (Open == 0) {
			continue
		} 
		days = append(days, &model.DayTrade{
			Date: cols[0],
			Code: strings.TrimLeft(cols[1], "'"),
			Close: float32(Close),
			High: float32(High),
			Low: float32(Low),
			Open: float32(Open),
			Volume: Volume,
			Money: int(Money),
		})
	}
	return days
}

/**
 * 对某一只股票进行更新日交易历史
 * @param  {[type]} this *DayTradeHelper)    update(stock *model.Stock [description]
 * @return {[type]}      [description]
 */
func (this *DayTradeHelper) Update(stock *model.Stock) {
	day := storage.GetLatestDayStock(stock)
	now := time.Now().Format("20060102")
	var added []*model.DayTrade
	if day == nil || day.Date == "" {
		// init all
		before := time.Unix(time.Now().Unix() - 3600*24*365*5, 0).Format("20060102")
		added = this.getTrade(stock, before, now)
	} else {
		log.Println("has inserted ", stock.Code)
		return
		// init today
		added = this.getTrade(stock, now, now)
	}
	
	storage.InsertTradeHis(added)
	log.Println("add ", stock.Code, len(added), " days trade info")
}
