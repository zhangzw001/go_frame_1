package input

import (
	"fmt"
	"testing"
)

func TestReadAunt(t *testing.T) {
	list := ReadAunt("./aunt.csv")
	fmt.Println(list[:3])
}


func TestReadOrder(t *testing.T) {
	list := ReadOrder("./order.csv")
	fmt.Println(list[:3])
}
