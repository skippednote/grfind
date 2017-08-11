Goodreads Find
---

> Find books on goodreads using title, isbn or author.

### Install
`go get github.com/skippednote/grfind`

### Usages
```go
package main

import (
	"encoding/json"	
	"log"
	"os"

	"github.com/skippednote/grfind"
)

func main() {
	gr := grfind.GRfind{
		Key:    "GoodReads Key",
		Secret: "GoodReads Secret",
	}
	books, err := gr.GetBooks("Elon Musk")
	if err != nil {
		log.Fatal(err)
	}

	// You can print the results
	for _, v := range books {
		log.Println(v.BestBook.Title)
	}

	// or return as JSON
	json.NewEncoder(os.Stdout).Encode(books)
}
```
