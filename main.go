package main

import (
	"fmt"

	s "github.com/samajain/goalgo/stack"
)

func main() {

	st := s.NewStack()
	st.Push("s1")
	st.Push("s2")
	st.Push("s3")
	st.Push("s4")
	a := st.Pop()
	fmt.Println(a)
	fmt.Println(st.Top())
	a = st.Pop()
	fmt.Println(a)
	fmt.Println(st.Top())
}
