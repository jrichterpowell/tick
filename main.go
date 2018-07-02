package main

import (
	"io/ioutil"
	"github.com/gizak/termui"
)

func main () {

	//Read the API key
	b_key, err := ioutil.ReadFile("key")
	check(err)
	API_KEY := string(b_key)

	//Initialize the UI
	initTermui()
	defer termui.Close()

	tickerHandler(API_KEY)

}
