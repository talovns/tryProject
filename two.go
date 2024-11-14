package main

import (
	"math"
)

var st string = " (3 + 2) * (4 - 2) / 2"

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
	st = "0123456789"
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

func ReturnZnacheniaVirajenia(lst []string) []string {
	if len(lst) == 1 {
		rez := lst
		return rez
	} else {
		for i := 0; i < len(lst)-1; i++ {
			if lst[i] == "*" {
				//  сделано: надо перевести строки в числа перемножить и обратно в строку чтобы заретернить в список с новыми значениями
				umnoj := PerevodStrVvInt(lst[i-1]) * PerevodStrVvInt(lst[i+1])
				lst[i] = string(umnoj)
				// понять как собрать список из срезлов +- так два раза сделать надо:
				// users = append(users[:n], users[n+1:]...)
				return lst[:i-1]
			}
		}
	}
	return []string{}
}

func main() {
	for _, value := range st {
		if (value) != ' ' {
			lst = append(lst, string(value))
		}
	}
	for i := 0; i < len(lst); i++ {
		promrez := ""
		if lst[i] == ")" {
			for j := i - 1; lst[j] != "("; j-- {
				promrez += lst[j]
			}
			promrez = ""
		}
		// заменять скобки в списке на значение из функции которую нужно написать для рассчитывания значений в скобках

	}
	// fmt.Print((PerevodStrVvInt("3")))
}
