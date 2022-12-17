package main

import (
	"fmt"
	"github.com/JimouChen/ds-set/easyset"
)

func main() {
	s := easyset.Set{}
	s.InitSet()
	s.Add("2")
	s.Add(1)
	s.Add(456)
	s.Show()
	_ = s.Delete(1)
	s.Show()
	fmt.Println(s.Length())
	fmt.Println(s.Contain("22"))
	s.Clear()
	fmt.Println(s.Length())
	s.Show()
}
