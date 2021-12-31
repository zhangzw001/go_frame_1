package input

import (
	"fmt"
	"testing"
)

func TestWorker(t *testing.T) {
	aunts := []Aunt{
		{1 ,0.9996, 8119, 6033,0} ,
		{2, 0.9991, -11796, 1050,0},
		{3, 0.9989, -720, -1818,0}}
	orders :=  []Order{
		{1518,1640736000,60,-1022,4758},
		{1355,1640736000,90,25229,2859},
		{1233,1640736000,120,-9082,1694}}
	fmt.Println(Worker(orders,aunts))
}
