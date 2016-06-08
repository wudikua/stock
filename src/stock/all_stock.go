package stock

import (
	. "constant"
	"fmt"
	"io/ioutil"
	"log"
	"model"
	"net/http"
	"regexp"
	"storage"
	"time"
)

type AllStock struct {
	Stocks     map[string]*model.Stock
	UpdateTime int64
}

var reg = regexp.MustCompile("~(?P<code>\\d+)`(?P<name>.*?)`(?P<py>\\w+)")

func GetAllStock() (map[string]*model.Stock, error) {
	log.Println("get all stock")
	resp, err := http.Get(ALL_STOCK_API)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	matches := reg.FindAllStringSubmatch(string(body), -1)
	stocks := make(map[string]*model.Stock)
	for _, m := range matches {
		s := &model.Stock{
			Code:   m[1],
			CnName: m[2],
			PyName: m[3],
			Type:   model.Code2Type(m[1]),
		}
		stocks[s.Code] = s
	}
	return stocks, nil
}

func NewAllStock() *AllStock {
	return &AllStock{}
}

func (this *AllStock) LoadFromStorage() {
	all := storage.GetAllStocks()
	this.Stocks = make(map[string]*model.Stock)
	for _, s := range all {
		this.Stocks[s.Code] = s
	}
}

func (this *AllStock) UpdateFromApi() {
	stocks, err := GetAllStock()
	if err != nil {
		fmt.Errorf("UpdateFromApi", err)
	}
	this.Stocks = stocks
	this.UpdateTime = time.Now().Unix()
}

func (this *AllStock) UpdateStorage() {
	s := make([]*model.Stock, len(this.Stocks))
	i := 0
	for _, v := range this.Stocks {
		s[i] = v
		i += 1
	}
	storage.UpSertStockInfo(s)
}

func (this *AllStock) String() string {
	stocks := this.Stocks
	str := ""
	if stocks == nil || len(stocks) == 0 {
		return str
	}
	for _, v := range stocks {
		str += fmt.Sprintf("%v\n", v)
	}
	tm := time.Unix(this.UpdateTime, 0)
	str += fmt.Sprintf("总共%d只股票\n", len(stocks))
	str += fmt.Sprintf("更新时间%s\n", tm.Format("2006-01-02 15:04:05"))
	return str
}
