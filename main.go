package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Banana", "Watermelon", "Grapes"}
	fmt.Printf("type of fruitlist is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Apple")

	fmt.Println(fruitList)

	fruitList = append(fruitList[2:5])
	fmt.Println(fruitList)

	highScores := make([]int, 5)

	highScores[0] = 155
	highScores[1] = 169
	highScores[2] = 325
	highScores[3] = 203
	highScores[4] = 185

	highScores = append(highScores, 501, 422, 369)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

}
