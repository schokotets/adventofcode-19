package main
import (
	"fmt"
	"math"
	"strings"
	"strconv"
	"io/ioutil"
)

func fuelreq(mass int64) int64 {
	return (int64)(math.Floor(float64(mass)/3)-2)
}

func main() {
	dat, _ := ioutil.ReadFile("input")

	xes := strings.Split(string(dat[:len(dat)-1]), ",")
	//fmt.Println(xes)
	L := make([]int, len(xes))
	for i := 0; i < len(xes); i++ {
		x, _ := strconv.ParseInt(xes[i], 10, 64)
		L[i] = int(x)
	}
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			l := make([]int, len(L))
			copy(l, L)
			l[1] = n
			l[2] = v
			cur := 0
			for cur < len(l) && l[cur] != 99 {
				opcode := l[cur]
				a := l[cur+1]
				b := l[cur+2]
				c := l[cur+3]

				if !(a < len(l) && b < len(l) && c < len(l)) {
					break
				}

				if opcode == 1 {
					l[c] = l[a] + l[b]
				}
				if opcode == 2 {
					l[c] = l[a] * l[b]
				}
				cur = cur + 4
			}
			fmt.Println(l)
			if l[0] == 19690720 {
				fmt.Printf("noun %v and value %v make %v", n, v, 100*n+v)
			}
		}
	}
	return
}
