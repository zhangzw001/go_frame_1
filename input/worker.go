package input

import (
	"math"
)

func init() {
	for i := 0 ; i < len(AuntLists);i++ {
		FirstAunt[AuntLists[i].Id] =false
	}
}
func Worker(orders Orders, aunts Aunts) (int,int ){
	// 需要第一时间处理相同时间的订单
	firstOrderBeginTime := orders[0].ServiceBeginTime
	i := 1
	for ; i<len(orders);i++ {
		if orders[i].ServiceBeginTime != firstOrderBeginTime {
			break
		}
	}
	// 这里存的都是相同开始时间的订单
	firstOrders := orders[:i]
	// 从这些订单中找到最优的阿姨
	maxScore := math.MaxFloat64 * -1
	var maxOrder Order
	var maxAunt Aunt
	// 遍历所有相同开始时间的订单
	for k:=0;k<len(firstOrders);k++ {
		order := firstOrders[k]
		// 需要对每个阿姨都计算score
		for j:=0;j<len(aunts);j++ {
			aunt := aunts[j]
			// 分别计算A,B,C
			a := aunt.ServiceScore
			b := GetEuclideanDistance(aunt.X,aunt.Y,order.X,order.Y)
			c := GetC(aunt.Id,b)
			// 如果间隔太远就放弃分配
			if c > 0.5 {
				continue
			}
			// 如果阿姨还没有空闲,就是上次的任务时间+赶路的时间 大于等于 当前订单的开始时间
			if float64(aunt.ServiceEndTime) + c*3600 >= float64(order.ServiceBeginTime) {
				continue
			}
			score := GetScore(a,b,c)
			// 记录最高分
			if maxScore < score {
				maxScore = score
				maxAunt,maxOrder = aunt,order
			}
		}
	}
	// 假定我们一定会找到一个阿姨,而不会出现所有的阿姨都赶不到的情况
	// 更新阿姨的坐标,是否是第一单,订单结束时间
	UpdateAunt(maxAunt,maxOrder)
	// 更新这个全局订单表
	OrderLists = UpdateOrders(maxOrder.Id)
	return maxAunt.Id,maxOrder.Id

}

func UpdateAunt(a Aunt,o Order) {
	// id号刚好比这个数字的索引号大1
	AuntLists[a.Id-1].X,AuntLists[a.Id-1].Y = o.X,o.Y
	FirstAunt[a.Id] = true
	// 需要更新阿姨的订单结束时间
	AuntLists[a.Id-1].ServiceEndTime = o.ServiceBeginTime+o.ServiceUnitTime*60
	//fmt.Println(">>更新阿姨的位置")
	//fmt.Println(a,o)
	//fmt.Println(AuntLists[a.Id-1])
}
func UpdateOrders(id int) Orders{
	for i:=0;i<len(OrderLists);i++ {
		if id == OrderLists[i].Id {
			return append(OrderLists[:i],OrderLists[i+1:]...)
		}
	}
	return OrderLists
}

// GetC 阿姨行走B距离所需要的时间
func GetC(id int, b float64) float64 {
	var C float64
	if !FirstAunt[id] {
		C = 0.5
	}else {
		C = b/float64(15)
	}
	return C
}
// GetScore 计算分数
func GetScore(A,B,C float64) float64{
	return A-0.5*B/float64(15)-0.25*C
}

// GetEuclideanDistance 计算阿姨距离订单的坐标的欧式距离 单位是千米
func GetEuclideanDistance(x1,y1, x2,y2 int) float64 {
	return  math.Sqrt(math.Pow(float64(x1-x2),2)+math.Pow(float64(y1-y2),2))/float64(1000)
}

