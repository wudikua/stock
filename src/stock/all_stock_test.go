package stock

import (
	"fmt"
	"testing"
)

func Test_GetAllStock(t *testing.T) {
	all := NewAllStock()
	all.UpdateFromApi()
	fmt.Println(all)
}
