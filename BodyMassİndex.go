package main

import "fmt"

func main() {

	var yourName string
	var yourWeight float32
	var yourLength float32

	fmt.Println("enter your your name")
	fmt.Scan(&yourName)

	fmt.Println("enter your weight")
	fmt.Scan(&yourWeight)

	fmt.Println("your lengthu write in meters")
	fmt.Scan(&yourLength)

	fmt.Println("your name  = ", yourName)
	fmt.Println("your weight = ", yourWeight)
	fmt.Println("your length = ", yourLength)

	bodyMassİndex := bodyMassİndexValue(yourName, yourWeight, yourLength)

	if bodyMassİndex >= 18 && bodyMassİndex <= 25 {

		fmt.Println("your mass index is normal ", bodyMassİndex)
	} else if bodyMassİndex < 18 {
		healthyWeight := 18 * (yourLength * yourLength)

		requiredWeight := healthyWeight - yourWeight
		fmt.Println("the weight you need to gain : +", requiredWeight)
		fmt.Println("body mass index", bodyMassİndex)
	} else if bodyMassİndex > 25 {
		healthyWeight := 25 * (yourLength * yourLength)

		weightYouNeedToLose := yourWeight - healthyWeight
		fmt.Println("weight you need to lose  -", weightYouNeedToLose)
		fmt.Println("body mass index", bodyMassİndex)
	}
	fmt.Println("healthy days..")

}

func bodyMassİndexValue(yourName string, yourWeight float32, yourLength float32) (bodyMassİndexValue float32) {
	bodyMassİndex := yourWeight / (yourLength * yourLength)
	return bodyMassİndex
}
