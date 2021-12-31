package input

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Aunt struct {
	Id int
	ServiceScore float64
	X int
	Y int
	ServiceEndTime int
}
// ReadAunt 从路径读取阿姨的信息
func ReadAunt(s string) Aunts {
	var auntLists Aunts
	auntF , err := os.Open(s)
	if err != nil {
		log.Fatalln(err)
	}
	defer auntF.Close()

	buf := bufio.NewScanner(auntF)
	for buf.Scan() {
		s1 := buf.Text()
		auntOne := strings.Split(s1,",")
		if len(auntOne) != 4 {
			continue
		}
		id,_ := strconv.Atoi(auntOne[0])
		serviceScore,_ := strconv.ParseFloat(auntOne[1],6)
		x,_ := strconv.Atoi(auntOne[2])
		y,_ := strconv.Atoi(auntOne[3])
		aunt := Aunt{id,serviceScore,x,y,0}
		auntLists = append(auntLists, aunt)
	}
	return auntLists
}



type Order struct {
	Id int
	ServiceBeginTime int
	ServiceUnitTime int
	X int
	Y int
}

// ReadOrder 从路径读取订单的信息
func ReadOrder(s string) Orders {
	var orderLists Orders
	auntF , err := os.Open(s)
	if err != nil {
		log.Fatalln(err)
	}
	defer auntF.Close()

	buf := bufio.NewScanner(auntF)
	// 一行行读取
	for buf.Scan() {
		s1 := buf.Text()
		auntOne := strings.Split(s1,",")
		if len(auntOne) != 5 {
			continue
		}
		// 这里直接强制转换了
		id,_ := strconv.Atoi(auntOne[0])
		serviceBeginTime,_ := strconv.Atoi(auntOne[1])
		serviceUnitTime,_ := strconv.Atoi(auntOne[2])
		x,_ := strconv.Atoi(auntOne[3])
		y,_ := strconv.Atoi(auntOne[4])
		order := Order{id,serviceBeginTime,serviceUnitTime,x,y}
		orderLists = append(orderLists, order)
	}
	// 排序
	sort.Sort(orderLists)
	return orderLists
}

