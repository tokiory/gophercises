package main

import (
    "os"
    "testing"
)

func TestOneLink(t *testing.T) {
    exerciseFile, err := os.Open("data/ex1.html")
    if err != nil {
        t.Errorf("Problem with exercise file: %s", err.Error())
    }

    expected := []Link{
        {
            Href: "/other-page",
            Text: "A link to another page",
        },
    }

    actual := getLinks(exerciseFile)

    if len(actual) != len(expected) {
        t.Errorf("Expected %d links, got %d", len(expected), len(actual))
    }
}

func TestTwoLinks(t *testing.T) {
    exerciseFile, err := os.Open("data/ex2.html")
    if err != nil {
        t.Errorf("Problem with exercise file: %s", err.Error())
    }

    expected := []Link{
        {
            Href: "https://www.twitter.com/joncalhoun",
            Text: "Check me out on twitter",
        },
        {
            Href: "https://github.com/gophercises",
            Text: "Gophercises is on Github!",
        },
    }

    actual := getLinks(exerciseFile)

    if len(actual) != len(expected) {
        t.Errorf("Expected %d links, got %d", len(expected), len(actual))
    }
}
