package input

type Aunts []Aunt

type Orders []Order
func (a Orders) Len() int {
	return len(a)
}

func (a Orders) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Orders) Less(i, j int) bool {
	return a[i].ServiceBeginTime < a[j].ServiceBeginTime
}
