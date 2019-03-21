package core

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	// change the User-Agent here if you want.  Maybe add-in later to change from command line or text file.
	useragent = "User-Agent': 'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.82 Safari/537.36"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func BaseRequest(urlAddress string, proxy string, requestType string) *http.Response {
	// if the proxy string is set try to pipe through the proxy with the requests.  otherwise
	// no need to delay on the base request.
	if proxy != "" {
		fmt.Println("using Proxy: ", proxy)
		proxyUrl, err := url.Parse(proxy)
		// check for error?
		if err == nil {
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
			return resp
		}
	} else if proxy == "" {
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
		return resp
	}
	// shouldn't ever get here...
	return nil
}

func MakeRequest(urlAddress string, proxy string, cookie *http.Cookie, requestType string) *http.Response {
	// if the proxy string is set try to pipe through the proxy with the requests.  otherwise
	if proxy != "" {
		fmt.Println("Cookie Value: ", cookie.Value)
		proxyUrl, err := url.Parse(proxy)
		// check for error?
		if err == nil {
			client := &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}

			req, err := http.NewRequest(requestType, urlAddress, nil)
			if err != nil {
				log.Fatalln(err)
			}
			// set the useragent.
			req.Header.Set("User-Agent", useragent)
			req.AddCookie(cookie)
			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}
			defer resp.Body.Close()
			fmt.Println("Request Length: ", resp.ContentLength)
			fmt.Println("Status: ", resp.Status)
			return resp
		}
	} else if proxy == "" {
		fmt.Println("Cookie Value: ", cookie.Value)
		client := &http.Client{}

		req, err := http.NewRequest(requestType, urlAddress, nil)
		if err != nil {
			log.Fatalln(err)
		}
		// set the useragent.
		req.Header.Set("User-Agent", useragent)
		req.AddCookie(cookie)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		fmt.Println("Request Length: ", resp.ContentLength)
		fmt.Println("Status: ", resp.Status)
		return resp
	}
	// shouldn't ever get here...
	return nil
}

func FuzzyWuzzy(urlAddress string, proxy string, cookie *http.Cookie, delay time.Duration, requestType string) {
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

	fmt.Println("Fuzzing Cookie: ", cookie.Value)
	// For every line in fuzz file run three attacks, fuzz before, after and replace.
	for scanner.Scan() {
		// sleep for <x> delay after each request.
		// after cookie attack.
		cookie.Value = cookie.Value + scanner.Text()
		MakeRequest(urlAddress, proxy, cookie, requestType)
		cookie.Value = tempCookie
		time.Sleep(delay * time.Millisecond)

		// before cookie attack.
		cookie.Value = scanner.Text() + cookie.Value
		MakeRequest(urlAddress, proxy, cookie, requestType)
		cookie.Value = tempCookie
		time.Sleep(delay * time.Millisecond)

		// replace cookie attack.
		cookie.Value = scanner.Text()
		MakeRequest(urlAddress, proxy, cookie, requestType)
		cookie.Value = tempCookie
		time.Sleep(delay * time.Millisecond)

	}
	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}

func BaseLine(urlAddress string, proxy string, requestType string, delay time.Duration, fileLocation string) {
	// Get original request, collect all observed cookies and take a baseline screenshots

	// maybe open file here and pass the object needed for writing?
	fmt.Println("Base Request Time...")
	resp := BaseRequest(urlAddress, proxy, requestType)

	// get number of cookies - you can get this from the Original Make Request above.
	fmt.Println("Baseline Request Length: ", resp.ContentLength)
	fmt.Println("Baseline Status: ", resp.Status)

	//// for number of cookies Fuzz dem with a nested for loop of Fuzz params.
	if len(resp.Cookies()) < 1 {
		fmt.Println("No Cookies Found!. Try again with more cookies. ")
		os.Exit(0)
	} else {
		// for loop through the individual cookies
		for _, cookie := range resp.Cookies() {
			fmt.Println("Request Number: ", cookie)
			//MakeRequest(urlAddress,proxy,cookie)
			FuzzyWuzzy(urlAddress, proxy, cookie, delay, requestType)
		}
	}
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
