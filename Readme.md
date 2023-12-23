## Installation

```
go get github.com/Galangrs/helper
```

## Example Code

### Fetching

```
package main

import (
	"fmt"

	fetch "github.com/Galangrs/helper/fetch"
)

func main() {
	headers := fetch.Header{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	data := fetch.Data{
		"name": "John",
		"age":  30,
	}

	jsonBytes, err := fetch.SendRequest("POST", "https://example.com", data, headers)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonBytes))
}
```
