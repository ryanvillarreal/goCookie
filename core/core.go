package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func MakeRequest(urlAddress string, proxy string){
	// if the proxy string is set try to pipe through the proxy with the requests.  otherwise
	if proxy != ""{
		fmt.Println("using Proxy...", proxy)
		proxyUrl, err := url.Parse(proxy)
		// check for error?
		if err == nil{
			myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
			resp, err := myClient.Get(urlAddress)
			if err == nil {
				fmt.Println(resp)
			}
		}

	} else{
		fmt.Println("Not using proxy")
		resp, err := http.Get(urlAddress)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(body))
	}
}

func ProxySettings(){
	// add support for Proxy Settings here.
}

func OpenFuzzList(){
	// read in the file here.  Eventually need to change this over to reading in the argument
	dat, err := ioutil.ReadFile("./fuzz.txt")
	check(err)
	fmt.Println(string(dat))
}

func BaseLine(urlAddress string, proxy string){
	// Get original request, collect all observed cookies and take a baseline screenshots
	MakeRequest(urlAddress, proxy)
	// get number of cookies

	// for number of cookies Fuzz dem with a nested for loop of Fuzz params.
}
