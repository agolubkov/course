package main

import (
	"fmt"
	"time"

	"modulestest/internal/packages/listeq"
	"modulestest/leap"

	"github.com/agolubkov/course"
	_ "github.com/agolubkov/course/week_1/leap"
	"github.com/agolubkov/course/week_1/sublist"
	td "github.com/agolubkov/course/week_1/twelve-days"
)

func main() {
	fmt.Println(td.Song())
	course.PrintSong()
	fmt.Println(leap.IsLeapYear(time.Now().Year()))
	fmt.Println(listeq.Sublist([]int{1}, []int{2}))
	fmt.Println(sublist.Sublist([]int{1}, []int{1}) == sublist.RelationEqual)
	fmt.Println(listeq.IsEqual([]int{1}, []int{2}))
}
