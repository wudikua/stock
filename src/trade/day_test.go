package trade

import (
	"testing"
	"model"
	"stock"
	"log"
)

func Test_NewDayTradeHelper(t *testing.T) {
	stock := &model.Stock{
		Code: "600343",
		Type: model.HU_A,
	}
	// stock2 := &model.Stock{
	// 	Code: "002563",
	// 	Type: model.SHEN_A,
	// }
	// stock3 := &model.Stock{
	// 	Code: "300075",
	// 	Type: model.CHUANGYE,
	// }
	helper := NewDayTradeHelper()
	helper.Update(stock)
	// helper.update(stock2)
	// helper.update(stock3)
}

func update(helper *DayTradeHelper, stock *model.Stock, sem chan int) {
	helper.Update(stock)
	<- sem
}

func Test_Update(t *testing.T) {
	all := stock.NewAllStock()
	all.UpdateFromApi()
	helper := NewDayTradeHelper()
	log.Println("update")
	sem := make(chan int, 20)
	for _, stock := range all.Stocks {
		if stock.Type == model.HU_A || 
			stock.Type == model.SHEN_A ||
			stock.Type == model.CHUANGYE {
			sem <- 1
			go update(helper, stock, sem)
		}
	}
}
