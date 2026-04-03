// package main

// import (
// 	"first-program/mathlab"
// 	"first-program/testin"
// 	"fmt"
// )

// func add(a int, b int) int {
// 	res := a + b
// 	return res
// }
// func check(a int, b int) bool {
// 	if a == b {
// 		return true
// 	} else {
// 		return false
// 	}
// }
// func init() {
// 	fmt.Println("I'm Init function")
// }

// func main() {
// 	var name string
// 	var age int

// 	var a int
// 	var b int
// 	var res int
// 	var equality bool

// 	fmt.Scan(&a)
// 	fmt.Scan(&b)

// 	res = add(a, b)

// 	fmt.Scan(&name)
// 	fmt.Scan(&age)

// 	equality = check(a, b)
// 	test := mathlab.Add(a, b)
// 	sub := testin.Sub(a, b)

// 	fmt.Println("name is: ", name)
// 	fmt.Println("age is: ", age)
// 	fmt.Println("res is: ", res)
// 	fmt.Println("test is: ", test)
// 	fmt.Println("sub is: ", sub)

// 	if equality {
// 		fmt.Println("True")
// 	} else {
// 		fmt.Println("False")
// 	}

// 	// IIFE
// 	func(x, y int) {
// 		fmt.Println(x + y)
// 	}(5, 6)

// 	// Assign function in variable
// 	add := func(x, y int) {
// 		fmt.Println(x + y)
// 	}

// 	add(4, 6)
// }

// ***************************************** Basics **************************************************************
/*
func greeting(name string) {
	fmt.Println("You are Welcome here Mr. ", name)
}
func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func mul(a int, b int) int {
	return a * b
}
func div(a int, b int) int {
	return a / b
}

func main() {
	var name string
	fmt.Scanln(&name)
	greeting(name)

	fmt.Print("Please take 2 number as input: ")

	var a, b int
	fmt.Scanln(&a, &b)

	fmt.Println("Which operation you want to do?\n+\n-\n*\n/")
	var ch string
	fmt.Scanln(&ch)

	switch ch {

	case "+":
		fmt.Println("sum is : ", add(a, b))
		break
	case "-":
		fmt.Println("sub is : ", sub(a, b))
		break
	case "*":
		fmt.Println("multi is : ", mul(a, b))
		break
	case "/":
		if b != 0 {
			fmt.Println("div is : ", div(a, b))
		} else {
			fmt.Println("Divide by Zero")
		}
		break
	default:
		fmt.Println("Don't Joke with me! I don't have any time for joking.")
	}
}

*/

package main

import (
	"first-program/cmd"
)

func main() {

	// fmt.Println(cnf.Version)
	// fmt.Println(cnf.ServiceName)
	// fmt.Println(cnf.HttpPort)
	cmd.Serve()

	//********** Based 64 *************

	// s := "a"

	// byteArr := []byte(s)
	// fmt.Println(s)
	// fmt.Println(byteArr)

	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)
	// b64str := enc.EncodeToString(byteArr)

	// fmt.Println(b64str) // check in encode to0 base64 format site

	// decodedstr, err := enc.DecodeString(b64str)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(decodedstr)

	//************** SHA - Secure Hash Algorithm **************

	// data := []byte("Hello World")
	// hash := sha256.Sum256(data)
	// fmt.Println("Hash after SHA-256: ", hash)

	//**************HMAC - Hash-based Message Authentication Code **********

	// secret := []byte("my-secret")
	// message := []byte("Hello World")

	// h := hmac.New(sha256.New, secret)
	// h.Write(message)

	// text := h.Sum(nil)

	// fmt.Println(text)

	// jwt, err := util.CreateJWT("my-secret", util.Payload{
	// 	Sub:         40,
	// 	FirstName:   "Habibur",
	// 	LastName:    "Rahman",
	// 	Email:       "habibur@gmail.com",
	// 	IsShopOwner: true,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(jwt) // check from jwt.io site
}
