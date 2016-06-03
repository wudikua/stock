package stock

import (
	"fmt"
	"testing"
)

func Test_GetAllStock(t *testing.T) {
	all := NewAllStock()
	all.updateFromApi()
	fmt.Println(all)
}
