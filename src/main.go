package main

import (
	"model"
	"stock"
	"trade"
	"log"
	"web"
	"flag"
	"fmt"
)

func update(helper *trade.DayTradeHelper, stock *model.Stock, sem chan int) {
	helper.Update(stock)
	<- sem
}

func dailyTradeUpdate() {
	all := stock.NewAllStock()
	all.UpdateFromApi()
	helper := trade.NewDayTradeHelper()
	log.Println("update")
	sem := make(chan int, 1)
	for _, stock := range all.Stocks {
		if stock.Type == model.HU_A || 
			stock.Type == model.SHEN_A ||
			stock.Type == model.CHUANGYE ||
			stock.Type == model.ZHONG_XIAO {
			sem <- 1
			go update(helper, stock, sem)
		}
	}
}

func dailyAllStockUpdate() {
	all := stock.NewAllStock()
	all.UpdateFromApi()
	all.UpdateStorage()
}
func httpServer() {
	web.Start()
}

func main() {
	boolServer := flag.Bool("web", false, "start up a webserver to query stocks info")
	boolTrade := flag.Bool("dailyTrade", false, "update day trade info")
	boolStock := flag.Bool("dailyStock", false, "update stock code->name info")
	flag.Parse()
	if (*boolServer) {
		httpServer()
	} else if (*boolTrade) {
		dailyTradeUpdate()
	} else if (*boolStock) {
		dailyAllStockUpdate()
	} else {
		fmt.Println("error input, use --help shows cmd")
	}
}