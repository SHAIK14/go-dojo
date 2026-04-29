package main

import "fmt"

func main() {

	score = addscore(score)
	fmt.Println(score)
	addscoreP(&score)
	fmt.Println(score)

	fmt.Println(UserDetails{"asif", 26, 5.11})
	user := UserDetails{"asif", 26, 5.11}
	p := &user
	p.Name = "jhon"
	fmt.Println(user)

}
