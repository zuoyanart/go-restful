package serializer

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"testing"
)

func TestSerializerFull(t *testing.T) {
	// create a new  serializer  which will return the markdown as []byte but of the HTML form
	myMarkdownSerializer := SerializeFunc(func(val interface{}, options ...map[string]interface{}) ([]byte, error) {
		var b []byte
		if s, isString := val.(string); isString {
			b = []byte(s)
		} else {
			b = val.([]byte)
		}
		buf := blackfriday.MarkdownCommon(b)
		if len(options) > 0 && len(options[0]) > 0 && options[0]["sanitize"] != nil && options[0]["sanitize"].(bool) == true {
			buf = bluemonday.UGCPolicy().SanitizeBytes(buf)
		}
		return buf, nil
	})

	RegisterDefaults(defaultSerializers)
	originalLen := Len()
	//mySerializers := Serializers{}
	// or just use the default one which will cover more code

	// will be putted
	For("markdown", myMarkdownSerializer)

	// will not be putted becuase of dot on the key which is not allowed.
	For("markdown2.", myMarkdownSerializer)
	if l := Len(); l != originalLen+1 {
		t.Fatalf("Expecting serializers length: %d but got %d", originalLen+1, l)
	}

	// lets test the serialize
	expectingResult := "<h2>Hello</h2>\n"
	result, err := Serialize("markdown", "## Hello", Options{"sanitize": true})

	if err != nil {
		t.Fatalf("Serialization failed with error: %s", err.Error())
	}

	val := string(result)
	if val != expectingResult {
		t.Fatalf("Serialization wrong result, expecting: '%s' but got: '%s'", expectingResult, val)
	}

	// lets add a second serialier with the same key
	// it should give us the final result of both of the serializers' results.
	For("markdown", myMarkdownSerializer) // we use the same serializer so it should be process the same input two times and give us two the <h1>Hello</h1>\n<h1>Hello</h1>\n

	if l := Len(); l != originalLen+1 {
		t.Fatalf("Expecting serializers length: %d but got %d. We still have %d because the len is by the key, which both times is 'markdown'", originalLen+1, originalLen+1, l)
	}

	expectingFinalResult := expectingResult + expectingResult
	fresult, err := Serialize("markdown", "## Hello")

	if err != nil {
		t.Fatalf("Serialization failed with error: %s", err.Error())
	}

	fval := string(fresult)
	if fval != expectingFinalResult {
		t.Fatalf("expectingFinalResult wrong result, expecting: '%s' but got: '%s'", expectingFinalResult, fval)
	}

}
