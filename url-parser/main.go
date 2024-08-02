package main

import (
    "flag"
    "fmt"
    "golang.org/x/net/html"
    "os"
    "strings"
)

var FlagExercise = flag.String("e", "", "Exercise to process")

type Link struct {
    Href string
    Text string
}

func NewLink(n *html.Node) Link {
    href := ""

    return Link{href, parseNodeText(n)}

}

func walkNodeChilds(n *html.Node) []*html.Node {
    siblings := make([]*html.Node, 0, 4)
    child := n.FirstChild

    for child != nil {
        siblings = append(siblings, child)
        child = child.NextSibling
    }

    return siblings
}

func parseNodeText(n *html.Node) (text string) {
    childs := walkNodeChilds(n)

    for _, child := range childs {
        if child.Type == html.TextNode {
            text += child.Data
            continue
        }

        if child.Type == html.ElementNode && child.Data == "a" {
            continue
        }

        text += parseNodeText(child)
    }

    text = strings.TrimSpace(text)
    return
}

func getLinks(f *os.File) []Link {
    node, err := html.Parse(f)
    if err != nil {
        panic(err)
    }

    walkStack := make([]*html.Node, 1, 10)
    links := make([]Link, 0, 10)
    walkStack[0] = node

    for len(walkStack) > 0 {
        node := walkStack[0]

        if node.Type == html.ElementNode && node.Data == "a" {
            walkStack = walkStack[1:]
            href := ""
            for _, attr := range node.Attr {
                if attr.Key == "href" {
                    href = attr.Val
                    break
                }
            }

            links = append(links, Link{href, parseNodeText(node)})
            continue
        }

        siblings := walkNodeChilds(node)
        walkStack = append(siblings, walkStack[1:]...)
    }

    return links
}

func main() {
    flag.Parse()

    exerciseFile, err := os.Open(*FlagExercise)
    if err != nil {
        panic(fmt.Errorf("Problem with exercise file: %s", err.Error()))
    }

    links := getLinks(exerciseFile)

    fmt.Println(links)
}
