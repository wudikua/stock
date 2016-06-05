package web

import (
    "net/http"
    "html/template"
    "storage"
    "model"
    "log"
)

func Start() {
	http.Handle("/", http.FileServer(http.Dir("C:/Users/wudikua/IdeaProjects/stock/static")))
	http.HandleFunc("/query", query)
	http.ListenAndServe(":8080", nil)
}

type TradeQueryResult struct {
	Days []*model.DayTrade
	Code string
	Name string
}

func query(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	stock := &model.Stock {
		Code: r.Form["code"][0],
	}
	log.Println(stock, r.Form["begin"], r.Form["end"])
	days := storage.GetTradeHis(stock, r.Form["begin"][0], r.Form["end"][0])
	storage.GetStockInfo(stock)
	t := template.New("trade")
	t, _ = template.ParseFiles("C:/Users/wudikua/IdeaProjects/stock/static/trade.html")
    t.Execute(w, &TradeQueryResult{
    	Days: days,
    	Code: stock.Code,
    	Name: stock.CnName,
    })
}
