package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"apple", "tomato", "peach"}

	fmt.Println("Type of FruitList is %T\n", fruitList)

	fruitList = append(fruitList, "mango", "Banana")

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 334
	highScore[2] = 434
	highScore[3] = 134

	highScore = append(highScore, 555, 667, 321)

	fmt.Println(highScore)
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

}
