package main

import (
	"encoding/json"
	"flag"
	"fmt"       // 'flag' helps you parse command line arguments
	"io/ioutil" // we'll only use 'ioutil' to maek reading and printing
	"log"
	"net/http" //this is the 'http' package we'll be using to retrieve a page
	"os"       // 'os' gives you access to system calls
)

func main() {
	flag.Parse()        // This stuff with 'flag' converts the command line
	args := flag.Args() // arguments to a new variable named 'args'

	// if dates  aren't provided as an argument, show a message and exit.
	if len(args) < 2 {
		fmt.Println("Please specify a from and to date eg. 01-01-15 31-01-15. It will only return the last 31 days from today's date")
		os.Exit(1)
	}
	grvGetRaceMeetings(args[0], args[1])
	//grvGetRaceMeetings("25-09-2015", "24-09-2015")
	grvGetRaceMeetingsJSON(args[0], args[1])
}

func check(e error) { //Check for errors when writing to a file
	if e != nil {
		panic(e)
	}
	fmt.Println("Succesfully wrote output to file.")
}

func grvGetRaceMeetings(fromDate string, toDate string) {
	apiURL := "http://www.grv.org.au/api/api.php?action=getpastMeetings&from_date=" + fromDate + "&to_date=" + toDate
	res, err := http.Get(apiURL)
	fmt.Println("The http get return status code:", res.StatusCode) // returns the http statuscode.
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	grvRaceMeetingInfo := []byte(string(body))
	errfile := ioutil.WriteFile("grvRaceMeetings.txt", grvRaceMeetingInfo, 0644)
	check(errfile)
	text := string(body)
	fmt.Println(len(body), len(text))

	//use the ?race=<num> from this file to load all the races for that day
	//ie. http://www.grv.org.au/api/result.php?&race=1501150040
}

func grvGetRaceMeetingsJSON(fromDate string, toDate string) {
	apiURL := "http://www.grv.org.au/api/api.php?action=getpastMeetings&from_date=" + fromDate + "&to_date=" + toDate
	res, err := http.Get(apiURL)
	fmt.Println("The http get return status code:", res.StatusCode) // returns the http statuscode.
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	type grvRaceMeeting struct {
		Title string `json:"title,omitempty"`
		Start string `json:"start,omitempty"`
		URL   string `json:"url,omitempty"`
	}

	//var y map[string]interface{}
	var r grvRaceMeeting
	var parsed map[string]interface{}
	//str := json.Unmarshal([]byte(body), &parsed)
	str := json.Unmarshal([]byte(string(body)), &parsed)
	if str != nil {
		fmt.Println(str)
		return
	}
	fmt.Printf("%+v\n", r)
	//fmt.Printf("%s, %s, %s\n", RaceMeeting.Title, RaceMeeting.Start, RaceMeeting.URL)
}
