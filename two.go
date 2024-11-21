package main

import (
	"fmt"
	"math"
	"strconv"
)

var st string = " (3 + 2 / 8) * ( (4 - 2) / 2)"

var pr string = "0123456789"

var SpisokIzSpiskov = []string{}

var lst = []string{}

func revVirajenia(x string) []string {
	i := 0
	sp := []string{}
	for i < len(x) {
		sp = append(sp, string(x[len(x)-1-i]))
		i++
	}
	return sp
}

func strokaVSpisok(x string) []string {
	rez := []string{}
	for i := 0; i < len(x); i++ {
		rez = append(rez, string(x[i]))
	}
	return rez
}

func PerevodStrVvInt(x string) (rez int) {
	st := "0123456789"
	for i := len(x) - 1; i >= 0; i-- {
		for j := 0; j < len(st); j++ {
			if st[j] == x[i] {
				rez += j * int(math.Pow(10, float64(len(x)-i-1)))
				break
			}
		}
	}
	return
}

func virajenia(x float64, y float64, s string) (rez string) {
	switch s {
	case "+":
		rez = fmt.Sprint(x + y)
	case "-":
		rez = fmt.Sprint(x - y)
	case "/":
		rez = fmt.Sprint(x / y)
	case "*":
		rez = fmt.Sprint(x * y)
	}
	return
}

func VoBrtnPolski(strx string) string {
	flag := false
	for _, val := range strx {
		if val == '*' || val == '/' {
			flag = true
		}
	}
	var st = []string{}
	var pluss = ""
	for p := 0; p < len(strx); p++ {
		if strx[p] != '*' && strx[p] != '-' && strx[p] != '/' && strx[p] != '+' {
			pluss += string(strx[p])
			if p == len(strx)-1 {
				st = append(st, pluss)
				break
			}
		} else {
			st = append(st, pluss)
			st = append(st, string(strx[p]))
			pluss = ""
		}
	}
	var rez string = ""
	if flag {
		for i := 0; i < len(st); i++ {
			if len(rez) == 0 {
				if st[i] == "*" {
					rez += string(st[i-1]) + " " + string(st[i+1]) + " *"
				}
				if st[i] == "/" {
					rez += string(st[i-1]) + " " + string(st[i+1]) + " /"
				}
			} else {
				if st[i] == "*" {
					rez += " " + string(st[i+1]) + " *"
				}
				if st[i] == "/" {
					rez += " " + string(st[i+1]) + " /"
				}
			}
		}
	}
	for i := 0; i < len(st); i++ {
		if len(rez) == 0 {
			if st[i] == "+" {
				rez += string(st[i-1]) + " " + string(st[i+1]) + " +"
			}
			if st[i] == "-" {
				rez += string(st[i-1]) + " " + string(st[i+1]) + " -"
			}
		} else {
			if st[i] == "+" {
				rez += " " + string(st[i+1]) + " +"
			}
			if st[i] == "-" {
				rez += " " + string(st[i+1]) + " -"
			}
		}
	}
	return rez
}

func polski(st string) string {
	var rezk = ""
	var lst1 = []string{}
	var lst2 = []string{}
	var prom = ""
	for i := 0; i < len(st); i++ {
		if st[i] != ' ' && st[i] != '*' && st[i] != '/' && st[i] != '+' && st[i] != '-' {
			prom += string(st[i])
		} else if st[i] != ' ' && (st[i] == '*' || st[i] == '/' || st[i] == '+' || st[i] == '-') {
			lst2 = append(lst2, string(st[i]))
		} else {
			lst1 = append(lst1, string(prom))
			prom = ""
		}
	}
	for len(lst1) > 1 {
		cnt := 0
		for i := 0; i < len(lst1); i++ {
			if lst1[i] == "" {
				cnt++
			}
		}

		if len(lst1)-cnt > 2 {
			if lst1[1] == "" {
				y, _ := strconv.ParseFloat(lst1[2], 64)
				x, _ := strconv.ParseFloat(lst1[0], 64)
				lst1[2] = virajenia(x, y, lst2[0])
				lst1 = lst1[2:]
				lst2 = lst2[1:]
			} else {
				y, _ := strconv.ParseFloat(lst1[1], 64)
				x, _ := strconv.ParseFloat(lst1[0], 64)
				lst1[1] = virajenia(x, y, lst2[0])
				lst1 = lst1[1:]
				lst2 = lst2[1:]
			}
		} else {
			y, _ := strconv.ParseFloat(lst1[len(lst1)-1], 64)
			x, _ := strconv.ParseFloat(lst1[0], 64)
			rezk = (virajenia(x, y, lst2[0]))
			break
		}

	}
	return rezk
}

func main() {
	// var pocti = ""
	for _, value := range st {
		if (value) != ' ' {
			lst = append(lst, string(value))
		}
	}
	// fmt.Print(lst)
	var h = 0
	for i := 0; i < len(lst); i++ {
		promrez := ""
		if lst[i] == ")" {
			for j := i - 1; lst[j] != "("; j-- {
				promrez += lst[j]
				h = j
			}
			// fmt.Print(revVirajenia(promrez))
			st := ""
			s := revVirajenia(promrez)
			for k := 0; k < len(s); k++ {
				st += string(s[k])
			}
			lst[h-1] = (polski(VoBrtnPolski(st)))
			lst = append(lst[:h], lst[i+1:]...)
			// fmt.Print(lst, "\n")
			i = 0
			promrez = ""
		}
		// заменять скобки в списке на значение из функции которую нужно написать для ->
		// ->рассчитывания значений в скобках
	}

	var itog = ""
	for g := 0; g < len(lst); g++ {
		itog += lst[g]
	}

	fmt.Print(polski(VoBrtnPolski(itog)))

}
