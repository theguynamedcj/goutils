package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"

	//	"log"
	"math/big"
	mathrand "math/rand" // use mathrand for an alias for math/rand to avoid importing error
	"net/http"

	"github.com/atotto/clipboard"
	"github.com/showwin/speedtest-go/speedtest"
)

type Quote struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers   = "0123456789"
	special   = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func generatePassword(length int, includeUpper, includeLower, includeNumbers, includeSpecial, copyToClipboard bool) (string, error) {
	var chars string

	if includeUpper {
		chars += uppercase
	}

	if includeLower {
		chars += lowercase
	}

	if includeNumbers {
		chars += numbers
	}

	if includeSpecial {
		chars += special
	}

	if chars == "" {
		return "", fmt.Errorf("Must select at least one character type")
	}

	password := make([]byte, length)
	charsLen := big.NewInt(int64(len(chars)))

	for i := range length {
		num, err := rand.Int(rand.Reader, charsLen)
		if err != nil {
			return "", err
		}
		password[i] = chars[num.Int64()]
	}

	return string(password), nil
}

func main() {
	var selectedFunction int
	fmt.Print("Select a function (1 for number generator, 2 for hot and cold game, 3 for password generator, 4 for random quote, 5 for speed test or 0 to exit): ")
	_, err := fmt.Scan(&selectedFunction)
	if err != nil {
		fmt.Println("Err: Invalid selection, Please select 1, 2, 3, 4, 5 or 0")
		return
	}

	if selectedFunction > 0 {
		switch selectedFunction {
		case 1:
			{
				var limit int
				fmt.Print("Welcome to the Random Number generator! Please select your maximum number: ")
				_, err := fmt.Scan(&limit)
				if err != nil {
					fmt.Println("Error: Invalid input")
					return
				}
			}
		case 2:
			{
				var maxNum int
				fmt.Print("Welcome to the Hot and cold game! Please select your maximum number: ")
				_, err := fmt.Scan(&maxNum)
				if err != nil {
					fmt.Println("Error: Invalid input")
					return
				}
				actualNumber := mathrand.Intn(maxNum) // Use mathrand.Intn
				var guessedNumber int
				for {
					fmt.Print("Guess the number: ")
					_, err := fmt.Scan(&guessedNumber)
					if err != nil {
						fmt.Println("Error: Invalid input")
						continue

					}
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

		case 3:
			{
				var length int
				var includeUpper, includeLower, includeNumbers, includeSpecial, copyToClipboard bool
				var input string

				fmt.Print("Welcome to the password generator! Please, enter password length: ")
				_, err := fmt.Scan(&length)
				if err != nil || length <= 0 {
					fmt.Println("Error: Invalid Input,")
					return
				}
				if length > 10000000 {
					fmt.Print("WARNING: VALUES THIS LARGE CAN CAUSE YOUR MEMORY TO FILL UP AND COULD INITAITE A SIGKILL, WHICH COULD PREVENT THE PASSWORD FROM GENERATING!!!! WOULD YOU STILL LIKE TO PROCEED(NOT RECOMMENDED)? (y/n): ")
					fmt.Scan(&input)
					if input != "y" && input != "Y" {
						return
					}
				}
				fmt.Print("Include lowercase? (y/n): ")
				fmt.Scan(&input)
				includeLower = input == "y" || input == "Y"

				fmt.Print("Include upper? (y/n): ")
				fmt.Scan(&input)
				includeUpper = input == "y" || input == "Y"

				fmt.Print("Include numbers? (y/n): ")
				fmt.Scan(&input)
				includeNumbers = input == "y" || input == "Y"

				fmt.Print("Include special characters? (y/n): ")
				fmt.Scan(&input)
				includeSpecial = input == "y" || input == "Y"

				fmt.Print("Copy to Clipboard? (y/n): ")
				fmt.Scan(&input)
				copyToClipboard = input == "y" || input == "Y"

				password, err := generatePassword(length, includeUpper, includeLower, includeNumbers, includeSpecial, copyToClipboard)
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				fmt.Println("Password generated successfully!")
				fmt.Println("Your password is", password)
				if copyToClipboard {
					err := clipboard.WriteAll(password)
					if err != nil {
						fmt.Print(err)
					} else {
						fmt.Println("Password copied to the clipboard")
					}
				}
			}

		case 4:
			{
				resp, err := http.Get("https://zenquotes.io/api/random")
				if err != nil {
					fmt.Println("Failed to fetch code: ", err)
					return
				}
				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Failed to read response body:", err)
					return
				}

				var quotes []Quote
				err = json.Unmarshal(body, &quotes)
				if err != nil {
					fmt.Println("Failed to parse response body", err)
					return

				}
				if len(quotes) > 0 {
					fmt.Println("\n", quotes[0].Quote)
					fmt.Println("-", quotes[0].Author)
				}
			}
		case 5:
			fmt.Println("\nStarting speed test...")

			var speedtestClient = speedtest.New()
			serverList, err := speedtestClient.FetchServers()
			if err != nil {
				fmt.Println("Error: Failed to connect to server: ", err)
			}

			user, _ := speedtestClient.FetchUserInfo()
			fmt.Printf("Testing from %s (%s)...", user.Isp, user.IP)
			targets, _ := serverList.FindServer([]int{})

			for _, s := range targets {
				fmt.Println("\nTesting Upload speed...")
				if err := s.UploadTest(); err != nil {
					fmt.Println("Error: Failed to connect to server")
				}
				fmt.Printf("Upload Speed: %s", s.ULSpeed)

				fmt.Println("\nTesting Download speed...")
				if err := s.DownloadTest(); err != nil {
					fmt.Println("Error: Failed to connect to server")
				}
				fmt.Printf("Download Speed: %s", s.DLSpeed)

				fmt.Println("\nPinging Servers...")
				if err := s.PingTest(nil); err != nil {
					fmt.Println("Error: Failed to connect to server")
				}
				fmt.Printf("Ping: %s\n", s.Latency)
				s.Context.Reset()
			}

		default:
			{
				fmt.Println("Err: Invalid selection, Please select 1, 2, 3, 4, 5 or 0")
			}
		}
	}
}
