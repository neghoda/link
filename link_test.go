package link

import (
	"reflect"
	"strings"
	"testing"
)

func TestExtractLinks(t *testing.T) {
	testData := "<html><body><h1>Hello!</h1><a href=\"/other-page\">A link to another page</a></body></html>"
	expected := []Link{
		Link{
			Href: "/other-page",
			Text: "A link to another page",
		},
	}
	if result, _ := ExtractLinks(strings.NewReader(testData)); !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v to equal %v", expected, result)
	}

	testData = "<html><body><h1>Hello!</h1></body></html>"
	expected = []Link{}
	if result, _ := ExtractLinks(strings.NewReader(testData)); len(result) != 0 {
		t.Errorf("Expected %v to equal %v", expected, result)
	}

	testData = `<div>
    				<a href="https://www.twitter.com/joncalhoun">
      					Check me out on twitter
      					<i class="fa fa-twitter" aria-hidden="true"></i>
    				</a>
    				<a href="https://github.com/gophercises">
      					Gophercises is on <strong>Github</strong>!
    				</a>
  				</div>`
	expected = []Link{
		Link{
			Href: "https://www.twitter.com/joncalhoun",
			Text: "Check me out on twitter",
		},
		Link{
			Href: "https://github.com/gophercises",
			Text: "Gophercises is on Github!",
		},
	}
	if result, _ := ExtractLinks(strings.NewReader(testData)); !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v to equal %v", expected, result)
	}
	testData = "<a href=\"/dog-cat\">dog cat <!-- commented text SHOULD NOT be included! --></a>"
	expected = []Link{
		Link{
			Href: "/dog-cat",
			Text: "dog cat",
		},
	}
	if result, _ := ExtractLinks(strings.NewReader(testData)); !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v to equal %v", expected, result)
	}
}
