package main
import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func fuelreq(mass int64) int {
	return (int)(math.Floor(float64(mass)/3)-2)
}

func main() {
	x, _ := strconv.ParseInt(os.Args[1], 10, 64)
	fmt.Println(fuelreq(x))
	return
}
