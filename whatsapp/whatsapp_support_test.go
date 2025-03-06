package whatsapp

// MD to Whatsapp supports the following conversions
// - Bold: *text*
// - Italic: _text_
// - Strikethrough: ~text~
// - Code (inline): `text`
// - Monospace: ```text```
// - Quote : > text
// - Bulleted List: - text or * text
// - Numbered List: 1. text

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/verloop/markdown"
	"github.com/verloop/markdown/parser"
)

func TestAllCases(t *testing.T) {

	type TestCase struct {
		Input  []byte
		Output []byte
	}

	tests := []TestCase{
		TestCase{
			Input:  []byte(`testcase`),
			Output: []byte(`testcase`),
		},
		TestCase{
			Input:  []byte(`*test*`),
			Output: []byte(`_test_`),
		},
		TestCase{
			Input:  []byte(`_test_`),
			Output: []byte(`_test_`),
		},
		TestCase{
			Input:  []byte(`**test**`),
			Output: []byte(`*test*`),
		},
		TestCase{
			Input:  []byte(`__test__`),
			Output: []byte(`*test*`),
		},
		TestCase{
			Input: []byte(`*_test*
			testing`),
			Output: []byte(`__test_
			testing`),
		},
		TestCase{
			Input:  []byte(`_*test*_`),
			Output: []byte(`__test__`),
		},
		TestCase{
			Input: []byte(`_test_
			> asd
			` + "```" + `
			asd
			` + "```"),
			Output: []byte(`_test_
			> asd
			` + "```" + `
			asd
			` + "```"),
		},
		TestCase{
			Input: []byte(`This text is ___really important___`),
			Output: []byte(`This text is *_really important_*`),
		},
		TestCase{
			Input: []byte(`This text is **_really important_**`),
			Output: []byte(`This text is *_really important_*`),
		},
		TestCase{
			Input: []byte(`This is really***very***important text.`),
			Output: []byte(`This is really*_very_*important text.`),
		},
		TestCase{
			Input: []byte(`1. First item
2. Second item
3. Third item
4. Fourth item`),
			Output: []byte(`1. First item
2. Second item
3. Third item
4. Fourth item
`),
		},
		TestCase{
			Input: []byte(`~~This is really***very***important text.~~`),
			Output: []byte(`~This is really*_very_*important text.~`),
		},
		TestCase{
			Input: []byte(`- First item
- Second item
- Third item
- Fourth item`),
			Output: []byte(`- First item
- Second item
- Third item
- Fourth item
`),
		},
	}



	t.Log("tst")
	for _, test := range tests {
		testParser := parser.New()
		parsed := testParser.Parse(test.Input)
		opts := RendererOptions{}
		ren := NewRenderer(opts)
		res := markdown.Render(parsed, ren)
		t.Log(string(test.Input), "--", string(test.Output), "---->", string(res))
		assert.Equal(t, test.Output, res)
		fmt.Println(string(res))
	}

}
