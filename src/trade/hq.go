package trade

import (
	"constant"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"model"
	"net"
	"net/http"
	"stock"
	"strings"
	"time"
)

type HqHelper struct {
	client http.Client
	stocks *stock.AllStock
}

func NewHqHelper(stocks *stock.AllStock) *HqHelper {
	client := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*3)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(time.Now().Add(3 * time.Second))
				return c, nil
			},
		},
	}
	return &HqHelper{
		client: client,
		stocks: stocks,
	}
}

func (this *HqHelper) request(query string, c chan int) {
	url := fmt.Sprintf(constant.HQ_TRADE_126_API, query)
	// log.Println("request hq ", url)
	resp, err := this.client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	s := string(body)
	s = strings.TrimLeft(s, "_ntes_quote_callback(")
	s = strings.TrimRight(s, ");")
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		return
	}
	if data != nil {
		for code, o := range data {
			fmt.Println(o.(map[string]interface{})["update"], code,
				o.(map[string]interface{})["name"],
				o.(map[string]interface{})["price"])
		}
	}
	<-c
}

func (this *HqHelper) getFromUrl() *model.Hq {
	hq := &model.Hq{}
	query := ""
	n := 0
	c := make(chan int, 100)
	for _, stock := range this.stocks.Stocks {
		if stock.Type == model.HU_A {
			query = query + "0" + stock.Code + ","
			n += 1
		} else if stock.Type == model.SHEN_A || stock.Type == model.CHUANGYE || stock.Type == model.ZHONG_XIAO {
			query = query + "1" + stock.Code + ","
			n += 1
		} else {
			continue
		}
		if n == 100 {
			// 一批请求
			go this.request(query, c)
			c <- 1
			log.Println("fire ", len(c))
			query = ""
			n = 0
		}
	}

	return hq
}

/**
 * 对某一只股票持续获取分时图
 * @param  {[type]} this *HqHelper)    Update(stock *model.Stock [description]
 * @return {[type]}      [description]
 */
func (this *HqHelper) Update() {
	for true {

		this.getFromUrl()
		time.Sleep(time.Second * 5)
	}
}
