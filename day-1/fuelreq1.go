package main
import (
	"fmt"
	"math"
	"strings"
	"strconv"
	"io/ioutil"
)

func fuelreq(mass int64) int {
	return (int)(math.Floor(float64(mass)/3)-2)
}

func main() {
	dat, _ := ioutil.ReadFile("input")
	xes := strings.Split(string(dat), "\n")
	cum := 0
	for i := 0; i < len(xes)-1; i++ {
		x, _ := strconv.ParseInt(xes[i], 10, 64)
		cum = cum + fuelreq(x)
		fmt.Printf("%d - %d\t%d\n", x, fuelreq(x), cum)
	}
	fmt.Println(cum)
	return
}
