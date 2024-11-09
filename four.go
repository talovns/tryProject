package main

import "fmt"

func OBShPrEF(lst []string) {
	var pr = []string{}
	var rez = []string{}
	for _, value := range lst {
		for i := 0; i < len(value); i++ {
			for j := i; j < len(value); j++ {
				pr = append(pr, value[i:j])
			}
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
	var mx string
	for k := 1; k < len(rez); k++ {
		if len(rez[k-1]) >= len(rez[k]) {
			mx = rez[k-1]
		} else {
			mx = rez[k]
		}
	}
	fmt.Print(mx)

}

func main() {
	var lst = []string{"aflow", "bflower", "cflight"}
	OBShPrEF(lst)

}
