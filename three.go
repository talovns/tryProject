// package main

// import (
// 	"fmt"
// )

// func TReugPascala(x int) {
// 	var lst1 = []int{}
// 	var lst2 = []int{}
// 	if x > 0 {
// 		fmt.Print(1, "\n")
// 		for i := 2; i <= x; i++ {
// 			if i == 2 {
// 				fmt.Print(1, 1, "\n")
// 			} else if i == 3 {
// 				fmt.Print(1, 2, 1, "\n")
// 			} else if i == 4 {
// 				lst1 = append(lst1, 3)
// 				lst1 = append(lst1, 3)
// 				fmt.Print(1, lst1[0], lst1[1], 1, "\n")
// 			} else {
// 				fmt.Print(1, " ")
// 				for i := 0; i < len(lst1)-1; i++ {
// 					if i == 0 {
// 						lst2 = append(lst2, 1+lst1[0])
// 						lst2 = append(lst2, lst1[1]+lst1[0])
// 					} else {
// 						lst2 = append(lst2, lst1[i]+lst1[i+1])
// 					}
// 					if i == len(lst1)-2 {
// 						lst2 = append(lst2, 1+lst1[0])
// 					}
// 				}
// 				for r := 0; r < len(lst2); r++ {
// 					fmt.Print(lst2[r], " ")
// 				}
// 				lst1 = lst2
// 				lst2 = []int{}
// 				fmt.Print(1, "\n")
// 			}
// 		}
// 	} else {
// 		panic("введите число больше 0")
// 	}

// }

// var a int

// func main() {
// 	fmt.Scan(&a)
// 	TReugPascala(a)
// }
