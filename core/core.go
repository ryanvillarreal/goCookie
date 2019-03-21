package core

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	useragent = "User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.82 Safari/537.36"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func BaseRequest(urlAddress string, proxy string, requestType string) []*http.Cookie {
	// if the proxy string is set try to pipe through the proxy with the requests.  otherwise
	if proxy != "" {
		fmt.Println("using Proxy...", proxy)
		proxyUrl, err := url.Parse(proxy)
		// check for error?
		if err == nil {
			fmt.Println("Not using proxy")
			client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

			req, err := http.NewRequest(requestType, urlAddress, nil)
			if err != nil {
				log.Fatalln(err)
			}
			// set the useragent.
			req.Header.Set("User-Agent", useragent)

			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()
			return resp.Cookies()
		}
	} else {
		fmt.Println("Not using proxy")
		client := &http.Client{}

		req, err := http.NewRequest(requestType, urlAddress, nil)
		if err != nil {
			log.Fatalln(err)
		}
		// set the useragent.
		req.Header.Set("User-Agent", useragent)

		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		return resp.Cookies()
	}
	// shouldn't ever get here...
	return nil
}

func MakeRequest(urlAddress string, proxy string, cookie *http.Cookie) []*http.Cookie {
	// if the proxy string is set try to pipe through the proxy with the requests.  otherwise
	if proxy != "" {
		proxyUrl, err := url.Parse(proxy)
		// check for error?
		if err == nil {
			myClient := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
			resp, err := myClient.Get(urlAddress)
			if err == nil {
				return resp.Cookies()
			}
		}
	} else {
		// make the call without proxy
		resp, err := http.Get(urlAddress)
		if err == nil {
			return resp.Cookies()
		}
	}
	// shouldn't ever get here?
	return nil
}

func ProxySettings() {
	// add support for Proxy Settings here.
}

func FuzzyWuzzy(urlAddress string, proxy string, cookie *http.Cookie) {
	// need to figure out how to put the data before, overwrite, and after the original cookie.
	file, err := os.Open("./fuzz.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// make a copy of the cookie.Value for resetting during the for loop.
	var tempCookie string
	tempCookie = cookie.Value
	for scanner.Scan() {
		// after cookie attack.
		cookie.Value = cookie.Value + scanner.Text()
		MakeRequest(urlAddress, proxy, cookie)
		fmt.Println("Cookie Before Request: ", cookie)
		cookie.Value = tempCookie

		// before cookie attack.
		cookie.Value = scanner.Text() + cookie.Value
		MakeRequest(urlAddress, proxy, cookie)
		fmt.Println("Cookie After Request: ", cookie)
		cookie.Value = tempCookie

		// replace cookie attack.
		cookie.Value = scanner.Text()
		MakeRequest(urlAddress, proxy, cookie)
		fmt.Println("Replace Cookie: ", cookie)
		cookie.Value = tempCookie
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func BaseLine(urlAddress string, proxy string, requestType string) {
	// Get original request, collect all observed cookies and take a baseline screenshots
	fmt.Println("Base Request Time...")
	cookies := BaseRequest(urlAddress, proxy, requestType)
	// get number of cookies - you can get this from the Original Make Request above.
	fmt.Println(cookies)

	//// for number of cookies Fuzz dem with a nested for loop of Fuzz params.
	//if len(cookies) < 1 {
	//	fmt.Println("No Cookies Found!. Try again with more cookies. ")
	//	os.Exit(0)
	//} else {
	//	// for loop through the individual cookies
	//	for _, cookie := range cookies {
	//		fmt.Println("Request Number: ", cookie)
	//		//MakeRequest(urlAddress,proxy,cookie)
	//		FuzzyWuzzy(urlAddress, proxy, cookie)
	//	}
	//}
}
