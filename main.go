package main

import "fmt"

// func PerevodStrVvInt(x string) float64 {
// 	st := "0123456789"
// 	rez := 0
// 	for i := len(x) - 1; i >= 0; i-- {
// 		for j := 0; j < len(st); j++ {
// 			if st[j] == x[i] {
// 				rez += j * int(math.Pow(10, float64(len(x)-i-1)))
// 				break
// 			}
// 		}
// 	}
// 	pr := float64(rez)
// 	return pr
// }

// func ReturnZnacheniaVirajenia(lst string) (rez string) {
// 	if strings.Count(lst, "*") == 0 && strings.Count(lst, "/") == 0 {
// 		rez = lst
// 	} else {
// 		for i := 0; i < len(lst)-1; i++ {
// 			if string(lst[i]) == "/" {
// 				umnoj := PerevodStrVvInt(string(lst[i-1])) / PerevodStrVvInt(string(lst[i+1]))
// 				pr := strconv.Itoa(umnoj)
// 				rez = lst[:i-1] + pr + lst[i+2:]
// 			}
// 			if string(lst[i]) == "/" {
// 				umnoj := PerevodStrVvInt(string(lst[i-1])) / PerevodStrVvInt(string(lst[i+1]))
// 				pr := strconv.Itoa(umnoj)
// 				rez = lst[:i-1] + pr + lst[i+2:]
// 			}
// 		}
// 	}
// 	return ReturnZnacheniaVirajenia(rez)
// }

// // func ReturnZnacheniaVirajenia(lst []string) []string {
// // 	for len(lst) != 1 {

// // 		for i := 0; i < len(lst)-1; i++ {
// // 			if string(lst[i]) == "*" {
// // 				umnoj := PerevodStrVvInt(string(lst[i-1])) * PerevodStrVvInt(string(lst[i+1]))
// // 				pr := strconv.FormatFloat(umnoj, 'f', -1, 64)
// // 				rez = lst[:i-1] + pr + lst[i+2:]
// // 			}
// // 			if string(lst[i]) == "/" {
// // 				umnoj := (PerevodStrVvInt(string(lst[i-1]))) / (PerevodStrVvInt(string(lst[i+1])))
// // 				pr := strconv.FormatFloat(umnoj, 'f', -1, 64)
// // 				rez = lst[:i-1] + pr + lst[i+2:]
// // 			}
// // 		}
// // 		lst = rez
// // 	}
// // 	return lst
// // }

// // func esheReturn(lst string) (rez string) {
// // 	for (strings.Count(lst, "-") > 0) && (strings.Count(lst, "+") > 0) {

// // 		for i := 0; i < len(lst)-1; i++ {
// // 			if string(lst[i]) == "+" {
// // 				umnoj := PerevodStrVvInt(string(lst[i-1])) + PerevodStrVvInt(string(lst[i+1]))
// // 				pr := strconv.FormatFloat(umnoj, 'f', -1, 64)
// // 				rez = lst[:i-1] + pr + lst[i+2:]
// // 			}
// // 			if string(lst[i]) == "-" {
// // 				umnoj := (PerevodStrVvInt(string(lst[i-1]))) - (PerevodStrVvInt(string(lst[i+1])))
// // 				pr := strconv.FormatFloat(umnoj, 'f', -1, 64)
// // 				rez = lst[:i-1] + pr + lst[i+2:]
// // 			}
// // 		}
// // 		lst = rez
// // 	}
// // 	return
// // }

// // переделать все обратно в список для деления, так как будет флоат а там сложно будет ориентироваться по элементам
// func main() {
// 	// lst := []string{"1", "+", "3"}
// 	// fmt.Print(esheReturn(ReturnZnacheniaVirajenia(lst)))
// 	fmt.Print(ReturnZnacheniaVirajenia("1 + 3"))
// }

// занового ретерн:

func VoBrtnPolski(st string) {
	flag := false
	for _, val := range st {
		if val == '*' || val == '/' {
			flag = true
		}
	}
	var rez string = ""
	if flag {
		for i := 0; i < len(st); i++ {
			if len(rez) == 0 {
				if st[i] == '*' {
					rez += string(st[i-1]) + " " + string(st[i+1]) + " *"
				}
				if st[i] == '/' {
					rez += string(st[i-1]) + " " + string(st[i+1]) + " /"
				}
			} else {
				if st[i] == '*' {
					rez += " " + string(st[i+1]) + " *"
				}
				if st[i] == '/' {
					rez += " " + string(st[i+1]) + " /"
				}
			}
		}
	}
	for i := 0; i < len(st); i++ {
		if len(rez) == 0 {
			if st[i] == '+' {
				rez += string(st[i-1]) + " " + string(st[i+1]) + " +"
			}
			if st[i] == '-' {
				rez += string(st[i-1]) + " " + string(st[i+1]) + " -"
			}
		} else {
			if st[i] == '+' {
				rez += " " + string(st[i+1]) + " +"
			}
			if st[i] == '-' {
				rez += " " + string(st[i+1]) + " -"
			}
		}
	}
	fmt.Print(rez)
}

// func retoorn([]string){

// }

// var lst = []string{}

func main() {
	VoBrtnPolski("3+2*5/7-1")
	// retoorn(lst)
}
