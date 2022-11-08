package main

import (
	"fmt"
	"sort"
)

type MyStruct struct {
	Num  int
	Name string
}

type MyInt int

func (m *MyInt) showYourSelf() {
	fmt.Printf("%T %v\n", m, m)
}

func (m *MyInt) add(i MyInt) {
	*m = *m + MyInt(i)
}

func LogSomeData(in fmt.Stringer) {
	fmt.Printf("data: %s\n", in.String())
}

type kilometer int

func (k *kilometer) String() string {
	return fmt.Sprintf("%09d km", k)
}

func (m *MyStruct) String() string {
	return "Num: " + fmt.Sprintf("%d", m.Num) + " and Name: " + m.Name
}

type mySliceSorted interface {
	sorter() mySliceStruct
}
type mySliceStruct []MyStruct

func main() {
	i := MyInt(0)

	i.add(3)
	i.showYourSelf()
	toSort := mySliceStruct{{9, "nine"}, {3, "some"}, {1, "first"}, {2, "any"}}
	var sorter mySliceSorted = toSort
	sorter = sorter.sorter()
	fmt.Println(sorter)
	tenKm := kilometer(10)
	LogSomeData(&tenKm)
	LogSomeData(&toSort[0])
}

// sorter implements interface mySliceSorted
func (sl mySliceStruct) sorter() mySliceStruct {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i].Num < sl[j].Num
	})
	return sl
}
