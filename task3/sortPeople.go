package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	birthDay  time.Time
}

type People []Person

func (obj People) Len() int { return len(obj) }

//There is a need to sort people in order of increasing age. Whenever two people have the same age - sort them by name.
func (obj People) Less(i, j int) bool {
	if obj[i].birthDay.Equal(obj[j].birthDay) {
		if obj[i].firstName < obj[j].firstName {
			return true
		}
	}
	return obj[i].birthDay.After(obj[j].birthDay)
}

func (obj People) Swap(i, j int) { obj[i], obj[j] = obj[j], obj[i] }

func sortPeople() {
	ivanIvanovDate, err := time.Parse("2006-Jan-02", "2005-Aug-10")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	konstantinKalininDate, err := time.Parse("2006-Jan-02", "2001-Oct-18")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	jackDanielDate, err := time.Parse("2006-Jan-02", "1850-Sep-05")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	georgeWashington, err := time.Parse("2006-Jan-02", "2001-Oct-18")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p := People{
		{"Jack", "Daniel", jackDanielDate},
		{"Ivan", "Ivanov", ivanIvanovDate},
		{"Konstantin", "Kalinin", konstantinKalininDate},
		{"George", "Washington", georgeWashington},
	}
	sort.Sort(p)
	//Unable to use sort.Sort(p) as value
	//fmt.Println(sort.Sort(p))
	fmt.Println("Sorted slice of people:")
	for i := 0; i < len(p); i++ {
		fmt.Println(p[i].firstName, p[i].lastName, p[i].birthDay)
	}
	fmt.Println()
}
