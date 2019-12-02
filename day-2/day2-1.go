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
	dat, _ := ioutil.ReadFile("input1")

	xes := strings.Split(string(dat[:len(dat)-1]), ",")
	fmt.Println(xes)
	l := make([]int, len(xes))
	for i := 0; i < len(xes); i++ {
		x, _ := strconv.ParseInt(xes[i], 10, 64)
		l[i] = int(x)
	}

	cur := 0
	for cur < len(l) && l[cur] != 99 {
		opcode := l[cur]
		if opcode == 1 {
			//fmt.Printf("%d: adding %d(#%d) and %d(#%d) to %d and storing to #%d\n", cur, l[l[cur+1]], l[cur+1], l[l[cur+2]], l[cur+2], l[l[cur+1]]+l[l[cur+2]], l[l[cur+3]])
			l[l[cur+3]] = l[l[cur+1]] + l[l[cur+2]]
		}
		if opcode == 2 {
			l[l[cur+3]] = l[l[cur+1]] * l[l[cur+2]]
		}
		cur = cur + 4
	}
	fmt.Println(l)
	return
}
