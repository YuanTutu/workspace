package main
//一个结构体（struct）就是一组字段（field）。

import "fmt"

type Vertex struct {
	X int
	Y int
}
func main() {
	fmt.Println(Vertex{1,2})
}
