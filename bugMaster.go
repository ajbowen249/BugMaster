package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ajbowen249/GoSandbox/console"
)

func main() {
	numCols := 80
	numRows := 25
	console.ClearScreen(numCols, numRows)
	console.SetTitle("Bug Master!")

	email := "alex.bowen@promessinc.com"
	baseURL := "https://promessdev.com/bugz/api"

	fmt.Print("password for ")
	fmt.Print(email)
	fmt.Print(": ")
	password := getPassword()
	fmt.Print("\n")

	res, err := http.Get(loginURL(baseURL, email, password))
	if err != nil {
		log.Fatal(err)
	}

	responseXML, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", responseXML)

	logR := new(webResponse)
	xml.Unmarshal([]byte(responseXML), logR)

	fmt.Println(logR.log.token)
}

func getPassword() string {
	password := ""

	for {
		isHit, char := console.GetKey()
		if isHit {
			if char == "\r" {
				return password
			}

			fmt.Print("*")
			password += char
		}
	}
}

func loginURL(baseURL string, email string, password string) string {
	return baseURL + ".asp?cmd=logon" + "&email=" + email + "&password=" + password
}

type webResponse struct {
	log loginResponse
}

type loginResponse struct {
	token string
}
