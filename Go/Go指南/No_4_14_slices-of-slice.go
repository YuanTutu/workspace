package main
//切片可包含任何类型，甚至包括其它的切片。

import (
	"fmt"
	"strings"
)
func main() {
	//创建一个井字板
	board:=[][]string{
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}
	//两个玩家轮流打上X和O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0;i < len(board);i++{
		fmt.Printf("%s\n",strings.Join(board[i]," "))
	}
}
