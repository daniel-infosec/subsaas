package main

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"net"
	"strings"
	"regexp"
	"net/url"
	"flag"
	"bufio"
	"os"
)

func main() {
	var org = flag.String("org", "", "The organization to search for")
	var orgListInput = flag.String("orglist", "", "File path with list of organizations")
	flag.Parse()
	var orgList []string
	if *orgListInput != "" {
	    file, err := os.Open(*orgListInput)
	    if err != nil {
	        log.Fatal(err)
	    }
	    defer file.Close()

	    scanner := bufio.NewScanner(file)
	    for scanner.Scan() {
	        //fmt.Println(scanner.Text())
	        orgList = append(orgList, scanner.Text())
	    }

	    if err := scanner.Err(); err != nil {
	        log.Fatal(err)
	    }
	}
	if *org != "" {
		orgList = append(orgList, *org)
	}
	
	fmt.Println("Splunk")
	fmt.Println(splunk(orgList))
	fmt.Println("Slack")
	fmt.Println(slack(orgList))
	fmt.Println("Zoom")
	fmt.Println(zoom(orgList))
	fmt.Println("Atlassian")
	fmt.Println(atlassian(orgList))
	fmt.Println("Okta")
	fmt.Println(okta(orgList))
	fmt.Println("Box")
	fmt.Println(box(orgList))
	fmt.Println("Adobe Creative Cloud")
	fmt.Println(adobecreativecloud(orgList))
}

type SlackMatch struct {
	Name string
	Email string
}

func splunk(s []string) []string {
	var success []string
	for _, name := range s {
		if resolveMatch(name + ".splunkcloud.com") {
			success = append(success, name)
		}
	}
	return success
}

func slack(s []string) []SlackMatch {
	var matches []SlackMatch
	for _, name := range s {
		exists, res := bodyMatch("https://" + name + ".slack.com", "There's been a glitch")
		if !exists {
			r, _ := regexp.Compile("(?:data-team-email-domains-formatted=\")[^\"]*")
			t := r.FindString(res)
			if t != "" {
				email := strings.Split(r.FindString(res),"\"")[1]
				match := SlackMatch {
					Name: name,
					Email: email,
				}
				matches = append(matches, match)
			} else {
				match := SlackMatch {
					Name: name,
					Email: "nil",
				}
				matches = append(matches, match)
			}
		}
	}
	return matches
}

func zoom(s []string) []string {
	var matches []string
	for _, name := range s {
		if resolveMatch(name + ".zoom.us") {
			matches = append(matches, name)
		}
	}
	return matches
}

func atlassian(s []string) []string {
	var matches []string
	for _, name := range s {
		exists, _ := bodyMatch("https://" + name + ".atlassian.net", "Your Atlassian Cloud site is currently unavailable.")
		if !exists {
			matches = append(matches, name)
		}
	}
	return matches
}

func okta(s []string) []string {
	var matches []string
	for _, name := range s {
		exists, _ := bodyMatch("https://" + name + ".okta.com", "' logo',")
		if !exists {
			matches = append(matches, name)
		}
	}
	return matches
}

func box(s []string) []string {
	var matches []string
	for _, name := range s {
		resp, err := http.Get("https://" + name + ".box.com")
		if err != nil {
	    	log.Fatalf("http.Get => %v", err.Error())
		}
		finalURL := resp.Request.URL.String()
		if finalURL != "https://account.box.com/login" {
			matches = append(matches, name)
		}
	}
	return matches
}

func salesforce(s []string) []string {
	var matches []string
	for _, name := range s {
		if resolveMatch(name + ".my.salesforce.com") {
			matches = append(matches, name)
		}
	}
	return matches
}

func adobecreativecloud(s []string) []string {
	var matches []string
	for _, name := range s {
		form := url.Values{}
		form.Add("client_id", "adobedotcom2")
		form.Add("username", "qwertydeadbeef@"+name+".com")
		form.Add("reauthenticate", "false")
		form.Add("idp_flow_type", "login")
		req, err := http.NewRequest("POST", "https://adobeid-na1.services.adobe.com/renga-idprovider/pages/login_flow", strings.NewReader(form.Encode()))
		if err != nil {
			log.Fatalln(err)
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		hc := http.Client{}
		resp, err := hc.Do(req)
		if err != nil {
			log.Fatalln(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		if strings.Contains(string(body), "while(1);\"fed\"") {
			matches = append(matches, name)
		}
	}
	return matches
}

func getBody(url string) string {
	resp, err := http.Get(url)
	
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}

func resolveMatch(s string) bool {
	_, err := net.LookupIP(s)
	if err == nil {
		return true
	} else {
		return false
	}
}

func bodyMatch(url string, s string) (bool, string) {
	res := getBody(url)
	if strings.Contains(res, s) {
		return true, res
	} else {
		return false, res
	}
}