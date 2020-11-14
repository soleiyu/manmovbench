// THERMAL Benchmark

package main

import (
	"fmt"
	"time"
	"strconv"
	"./pictFunc"
)

func main(){
	t0 := time.Now()

//	outFile := pictFunc.MkMandelMapQ_pr(720, 480, 4, 0.016, -0.667, 0.009, 0, 0, 0.36)
	
	for i := 0; i < 100; i ++ {
		outFile := pictFunc.MkMandelMapQ_pr(320, 240, 4, 4.0 / float64((i + 1) * (i + 1)), -0.667, 3.0 / float64((i + 1) *(i + 1)), 0, 0, 0.36)
		t1 := time.Now()
		outFile.Save("results/res" + strconv.Itoa(i) + ".png")
		fmt.Printf("%3d:%f\n", i, (t1.Sub(t0)).Seconds())
	}
//	fmt.Println("FINISH")
}

