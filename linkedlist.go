package main

import (
	"fmt"
)

type Page struct {
	DocumentLocation string
	References       []*Page
	DocumentName     string
	Parent           *Page
}

type Sitemap struct {
	Head *Page
}

// create new page:
// expecting str documentLoc, documentName
func newPage(documentLocation string, documentName string) *Page {
	return &Page{
		DocumentLocation: documentLocation,
		References:       make([]*Page, 0),
		DocumentName:     documentName,
		Parent:           nil,
	}
}

func (ll *Sitemap) append(documentLocation string, documentName string) {
	newNode := newPage(documentLocation, documentName)
	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	current := ll.Head
	for current.Parent != nil {
		current = current.Parent
	}
	current.Parent = newNode
}

func refName(pageMemAddr []*Page) string {
	var buffer string
	for _, v := range pageMemAddr {
		buffer += /*  v + ": " + */ v.DocumentName + ", "
	}

	return buffer

}

// prints contents
func (ll *Sitemap) print() {
	current := ll.Head
	// for current != nil {
	// 	fmt.Printf("Document Location: %s, Document Name: %s\n", current.DocumentLocation, current.DocumentName)
	// 	fmt.Printf("References: %v\n", current.References)
	// 	current = current.Parent
	// }
	for current != nil {
		fmt.Printf("document name: %s contains references to %s\n", current.DocumentName, refName(current.References))
		current = current.Parent

	}
}

func (node *Page) addRef(refNode *Page) {
	node.References = append(node.References, refNode)
}

// func main() {

// 	sitemap := Sitemap{}
// 	new := newPage("facebook.com", "find your")
// 	sitemap.append("https://test-website.sppcontests.org/south-pacific-icpc/", "Page 1")
// 	sitemap.Head.addRef(new)
// 	sitemap.print()
// }
