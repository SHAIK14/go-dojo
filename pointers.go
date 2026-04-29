package main

var score int = 10

func addscore(x int) int {
	x += 5
	return x

}

func addscoreP(p *int) {
	*p += 5

}

// func main() {

// 	score = addscore(score)
// 	fmt.Println(score)
// 	addscoreP(&score)
// 	fmt.Println(score)

// }
