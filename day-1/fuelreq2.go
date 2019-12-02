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
	xes := strings.Split(string(dat), "\n")
	cum := int64(0)
	for i := 0; i < len(xes)-1; i++ {
		x, _ := strconv.ParseInt(xes[i], 10, 64)
		o := x
		for f := fuelreq(x); f > 0; f = fuelreq(f) {
			x = x + f
		}
		cum = cum + x - o
	}
	fmt.Println(cum)
	return
}
