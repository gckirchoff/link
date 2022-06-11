package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a parsed a tag link from an HTML page
type Link struct {
	Href string
	Text string
}

var Reader io.Reader

// Parse will take in an HTML document and return a slice of Links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var builtLink Link
	for _, attribute := range n.Attr {
		if attribute.Key == "href" {
			builtLink.Href = attribute.Val
			break
		}
	}
	builtLink.Text = text(n)
	return builtLink
}

func text(n *html.Node) string {
	if (n.Type == html.TextNode) {
		return n.Data
	}
	var builtText string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		builtText += text(c) + " "
	}
	return strings.Join(strings.Fields(builtText), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var returnNodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		returnNodes = append(returnNodes, linkNodes(c)...)
	}
	return returnNodes
}

// func dfs (node *html.Node, padding string) {
// 	var content string = node.Data
// 	if node.Type == html.ElementNode &&  node.Data == "a" {
// 		content = "<" + content + ">"
// 	}
// 	fmt.Println(padding, content)
// 	for c := node.FirstChild; c != nil; c = c.NextSibling {
// 		dfs(c, padding + "  ")
// 	}
// }