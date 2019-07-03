package book

import (
	"fmt"
	"encoding/json"
	"strings"
)
type Book struct {
	Title string `json:title`
	Author string `json:author`
	Pages int `json:pages`
}

func (book Book) CategoryByLength() string {
	if book.Pages > 300 {
		return "NOVEL"
	} else {
		return "SHORT STORY"
	}
}

func NewBookFromJSON(jsonStr string) (Book, error) {
	var book Book
	err := json.Unmarshal([]byte(jsonStr), &book)
	if err != nil {
		fmt.Println(err)
	}
	return book, err
}

func (book *Book) AuthorLastName() string {
	slist := strings.Split(book.Author, " ")
	slen := len(slist)
	return slist[slen-1]
}