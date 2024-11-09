package main

import "fmt"

func OBShPrEF(lst []string) {
	var pr = []string{}
	var rez = []string{}
	for _, value := range lst {
		for i := 1; i < len(value); i++ {
			pr = append(pr, value[:i])
		}
	}
	for r := 0; r < len(pr); r++ {
		zn := pr[r]
		cnt := 0
		for h := 0; h < len(pr); h++ {
			if zn == pr[h] {
				cnt++
			}
		}
		if cnt == len(lst) {
			rez = append(rez, zn)
		}
	}
	fmt.Print(rez[len(rez)-1])

}

func main() {
	var lst = []string{"flow", "flower", "flight"}
	OBShPrEF(lst)
}
