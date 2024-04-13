package main

// roguh outline:
// 	stream in html page from http into parser
// 	find all hrefs and recursively stream them into parser to find more hrefs

// func reqs:
//	accepts a subdomain			/
// 	traverses it in its entirely		..
// 	ignores external domains
// 	handles recursive linking

// issues:
//
//	handling recursivity
//	directory structure will not suffice- no hierarchy on the page

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func showErr(e error) {
	log.Fatalln(e)
}

func printArr(input []string) {
	fmt.Printf("%+q", input)
}

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	rootSubdomain := "https://test-website.sppcontests.org/"

	fmt.Println("traverser started!")

	//init site root structure
	treeRoot := Sitemap{}
	treeRoot.append(rootSubdomain, "root")

	//traverse given root, assign references
	// treeRoot.browse()

	// for _, v := range rootLinks {}

	// traverse(rootLinks)

}

func traverse(rootHrefs []string) {
	printArr(rootHrefs)

}

// browse page for anchors, retroactively adds references
func (targetPage *Page) browse() {
	response, e := http.Get(targetPage.DocumentLocation)
	var references []string

	if e != nil {
		showErr(e)
	} else {
		// reading html body on given page root
		html, _ := io.ReadAll(response.Body)
		response.Body.Close()
		references = parseInstance(bytes.NewReader(html))
		for _, v := range references {
			nP := newPage(v, v)
			targetPage.addRef(nP)
		}
	}
}

// browse page for anchors, retyurns string array
func visit(link string) (hrefs []string) {
	response, e := http.Get(link)
	var references []string

	if e != nil {
		showErr(e)
	} else {
		// reading html body on given page root
		html, _ := io.ReadAll(response.Body)
		response.Body.Close()
		references = parseInstance(bytes.NewReader(html))
	}
	return references
}

func parseInstance(in io.Reader) (hrefs []string) {
	tokenizerInstance := html.NewTokenizer(in)

	for {
		iter := tokenizerInstance.Next()

		if iter == html.ErrorToken {
			if tokenizerInstance.Err() == io.EOF {
				break
			}
			showErr(tokenizerInstance.Err())
		} else if iter == html.StartTagToken {
			tagString, _ := tokenizerInstance.TagName()
			if len(tagString) == 1 && tagString[0] == 'a' {
				rAttr, rVal, _ := tokenizerInstance.TagAttr()
				if len(rAttr) != 0 || len(rVal) != 0 {
					if string(rAttr) == "href" {
						if string(rVal) != "#" {
							// prunedUrl := strings.Split(string(rVal), "/")[2]
							hrefs = append(hrefs, string(rVal))
						}
					}
				}
			}
		}

	}

	// fmt.Println(hrefs)
	return hrefs
}
