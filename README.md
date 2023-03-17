# JSON Extractor
JSON Extractor is a Go module for extracting valid JSON objects from anywhere, It is also used in the [yak](https://github.com/yaklang/yakit).

## Installation
```bash
go get -u github.com/WAY29/jsonextractor
```

## Usage
### Fix JSON
```go
package main

import (
	"fmt"

	"github.com/WAY29/jsonextractor"
)

func main() {
    // fix invalid json escape
	str := `{"message": "This is not a valid JSON string \x0A"}`
	fixed, ok := jsonextractor.FixJson([]byte(str))
	fmt.Println(string(fixed)) // Output: {"message": "This is not a valid JSON string \u000a"}
    fixed, ok = FixJson([]byte(`{"abc": 123,}`))
	fmt.Println(string(fixed)) // Output: {"abc": 123}
}
```

### Extract JSON Objects
```go
package main

import (
	"fmt"

	"github.com/WAY29/jsonextractor"
)

func main() {
	str := `// something trush...
    <html> <!-- something html-->
    {"name": "John", "age": 25}{"name": "Jim", "age": 30}</html>`
	results, _, err := jsonextractor.ExtractJSONWithRaw(str)
    if err != nil {
        panic(err)
    }
	for _, r := range results {
		fmt.Println(r)
	}
    // Output:
	// {"name": "John", "age": 25}
	// {"name": "Jim", "age": 30}

    // or use ExtractStandardJSON
    results2 := jsonextractor.ExtractStandardJSON(str)
	for _, r := range results2 {
		fmt.Println(r)
	}
    // Output:
	// {"name": "John", "age": 25}
	// {"name": "Jim", "age": 30}
}
```

## License
The module is licensed under the MIT license.


## Thanks
Original author: [VillanCh](https://github.com/VillanCh)


