package jsonextractor

import (
	"fmt"
	"testing"
)

func TestFixJson(t *testing.T) {
	raw, ok := FixJson([]byte(`{"badi": "\r\x11"}`))
	if string(raw) != `{"badi": "\r\u0011"}` || !ok {
		panic("Fix json error: fix unicode error")
	}

	raw, ok = FixJson([]byte(`{"abc": 123,}`))
	if string(raw) != `{"abc": 123}` || !ok {
		panic("fix invalid json error")
	}

	raw, ok = FixJson([]byte(`{"message": "This is not a valid JSON string \x0A"}`))
	if string(raw) != `{"message": "This is not a valid JSON string \u000a"}` || !ok {
		panic("fix invalid json error")
	}
}

func TestExtractJSONWithRaw(t *testing.T) {
	raw := `<html>

aasdfasd
df
{
  "code" : "0",
  "message" : "success",
  "responseTime" : 2,
  "traceId" : "a469b12c7d7aaca5",
  "returnCode" : null,
  "result" : {
    "total" : 0,
    "navigatePages" : 8,
    "navigatepageNums" : [ ],
    "navigateFirstPage" : 0,
    "navigateLastPage" : 0
  }
}

</html>
{{
{"abc": 123}

{{{{{{   }} {"test":                     123}}
`
	results, rawStr, err := ExtractJSONWithRaw(raw)

	if err != nil {
		t.Fatal(err)
	}

	if results[1] != `{"abc": 123}` {
		panic(`Extract json error: no.2 json != {"abc": 123}`)
	}
	if results[2] != `{   }` {
		panic(`Extract json error: no.3 json != {   }`)
	}
	if results[3] != `{"test":                     123}` {
		panic(`Extract json error: no.4 json != {"test":                     123}`)
	}
	if rawStr[0] != `{{{   }} {"test":                     123}}` {
		panic(`Extract json error: no.1 raw json != {{{   }} {"test":                     123}}`)
	}
	if rawStr[1] != `{{   }}` {
		panic(`Extract json error: no.2 raw json != {{   }}`)
	}

}

func TestExtractStandardJSON(t *testing.T) {
	str := `// something trush...
    <html> <!-- something html-->
    {"name": "John", "age": 25}{"name": "Jim", "age": 30}</html>`
	results, _, _ := ExtractJSONWithRaw(str)
	for _, r := range results {
		fmt.Println(r)
	}
	// or use ExtractStandardJSON
	results2 := ExtractStandardJSON(str)
	for _, r := range results2 {
		fmt.Println(r)
	}
}
