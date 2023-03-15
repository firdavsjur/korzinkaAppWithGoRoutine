package main

import (
	"fmt"
	"time"
)

type product struct {
	productName  string
	productPrice int
}
type user struct {
	userName string
	products []product
}

func Cassir(users []user, from, to int) {
	for i := from; i < to; i++ {
		totalPrice := 0
		for _, v := range UsersArray[i].products {
			totalPrice += v.productPrice
		}
		fmt.Printf("User: %v,TotalPrice: %v\n", users[i].userName, totalPrice)
	}
}

func main() {
	start := time.Now()
	go Cassir(UsersArray, 0, len(UsersArray)/4)
	go Cassir(UsersArray, len(UsersArray)/4, len(UsersArray)/2)
	go Cassir(UsersArray, len(UsersArray)/2, len(UsersArray)/4*3)
	go Cassir(UsersArray, len(UsersArray)/4*3, len(UsersArray))
	elapsed := time.Since(start)
	time.Sleep(time.Millisecond * 1)
	fmt.Println("time: ", elapsed)

}
