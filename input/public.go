package input

var (
	AuntPath = "./input/aunt.csv"
	OrderPath = "./input/order.csv"
	AuntLists  = ReadAunt(AuntPath)
	OrderLists = ReadOrder(OrderPath)
	FirstAunt  = make(map[int]bool,len(AuntLists))
)
