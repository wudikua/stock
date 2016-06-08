package trade

import (
	// "log"
	// "model"
	"stock"
	"testing"
)

func Test_NewHqHelper(t *testing.T) {
	all := stock.NewAllStock()
	all.LoadFromStorage()
	hq := NewHqHelper(all)
	hq.Update()
}
