/* 01

package main

import "fmt"

func init () {
	// main関数より先に呼ばれる
	fmt.Println("init")
}

func bazz () {
	//定義しただけでは実行されない
	fmt.Println("Bazz")
}

func main () {
	//実行が必要
	bazz()
	//大文字から始まる関数はPublic
	fmt.Println("Hello, World!", "Hello, Go!")
	//コンマで繋げられる
}

*/

