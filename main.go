package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var selectedFunction int
	fmt.Print("Select a function (1 for number generator or 2 for hot and cold game or 3 to exit): ")
	fmt.Scan(&selectedFunction)

	if selectedFunction > 0 {
		switch {
		case selectedFunction == 1:
			{
				var limit int
				fmt.Print("Welcome to the random number generator!", ", Please select your maximum number : ")
				fmt.Scan(&limit)
				fmt.Println("Your number is", rand.Intn(limit))

			}
		case selectedFunction == 2:
			{
				var maxNum int
				fmt.Print("Welcome to the Hot and cold game!", ", Please select your maximum number: ")
				fmt.Scan(&maxNum)

				actualNumber := rand.Intn(maxNum)
				var guessedNumber int
				for {
					fmt.Print("Guess the number: ")
					fmt.Scan(&guessedNumber)

					if guessedNumber == actualNumber {
						fmt.Println("You got it correctly")
						return
					} else if guessedNumber < actualNumber {
						fmt.Println("Value too low! Please try again")
					} else {
						fmt.Println("Value too high! Please try again")
					}
				}
			}
<<<<<<< HEAD
			case selectedFunction == 3: {
				fmt.Println("Exiting...")
				return
			}
			default: {
				fmt.Println ("Err: Invalid value")
				fmt.Println ("Exiting...")
				return


=======
		case selectedFunction == 3:
			{
				fmt.Println("Exiting...")
				return
			}
		default:
			{
				fmt.Println("Err: Invalid value")
				fmt.Println("Exiting...")
				return

>>>>>>> 3666f9b (Initial commit)
			}
		}
	}
}
