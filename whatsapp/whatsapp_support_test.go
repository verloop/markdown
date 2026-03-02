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
	"bytes"
	"fmt"
	"testing"

	"github.com/verloop/markdown"
	"github.com/verloop/markdown/parser"
)

func TestAllCases(t *testing.T) {

	type TestCase struct {
		Input  []byte
		Output []byte
	}

	tests := []TestCase{
		{
			Input:  []byte(`normal text`),
			Output: []byte("normal text\n"),
		},
		{
			Input:  []byte(`**bold**`),
			Output: []byte("*bold*\n"),
		},
		{
			Input:  []byte(`*italic*`),
			Output: []byte("_italic_\n"),
		},
		{
			Input:  []byte(`~~strike~~`),
			Output: []byte("~strike~\n"),
		},
		{
			Input:  []byte("> quote"),
			Output: []byte("> quote\n"),
		},
		{
			Input:  []byte("```\ncode block\n```"),
			Output: []byte("code block\n"),
		},
		{
			Input:  []byte("`code`"),
			Output: []byte("`code`\n"),
		},
		{
			Input:  []byte(`**bold** and *italic* and ~~strike~~`),
			Output: []byte("*bold* and _italic_ and ~strike~\n"),
		},
		{
			Input:  []byte("- First item\n- Second item\n- Third item"),
			Output: []byte("- First item\n- Second item\n- Third item\n"),
		},
		{
			Input:  []byte("1. First item\n2. Second item\n3. Third item"),
			Output: []byte("1. First item\n2. Second item\n3. Third item\n"),
		},
		{
			Input:  []byte("* First item\n* Second item\n* Third item"),
			Output: []byte("- First item\n- Second item\n- Third item\n"),
		},
		{
			Input:  []byte("```go\nfmt.Println(\"hello\")\n```"),
			Output: []byte("fmt.Println(\"hello\")\n"),
		},
		{
			Input:  []byte(`\"Cards\"`),
			Output: []byte("\"Cards\"\n"),
		},
		{
			Input:  []byte(`\'Option A\'`),
			Output: []byte("'Option A'\n"),
		},
		{
			Input:  []byte(`[Link Text](https://verloop.io)`),
			Output: []byte("https://verloop.io\n"),
		},
		{
			Input:  []byte(`[](https://verloop.io)`),
			Output: []byte("https://verloop.io\n"),
		},
		{
			Input:  []byte(`check this out [Link Text](https://verloop.io) and more text`),
			Output: []byte("check this out https://verloop.io and more text\n"),
		},
		{
			Input:  []byte(`https://verloop.io`),
			Output: []byte("https://verloop.io\n"),
		},
		{
			Input:  []byte(`# Heading level 1`),
			Output: []byte("Heading level 1\n"),
		},
		{
			Input:  []byte(`## Heading level 2`),
			Output: []byte("Heading level 2\n"),
		},
		{
			Input:  []byte(`### Heading level 3`),
			Output: []byte("Heading level 3\n"),
		},
		{
			Input:  []byte(`###### Heading level 6`),
			Output: []byte("Heading level 6\n"),
		},
		{
			Input:  []byte("some text\n# Heading"),
			Output: []byte("some text\n\nHeading\n"),
		},
		{
			Input:  []byte("# Heading 1\n## Heading 2\n### Heading 3"),
			Output: []byte("Heading 1\n\n\nHeading 2\n\n\nHeading 3\n"),
		},
		{
			Input:  []byte("# Heading\nsome text below"),
			Output: []byte("Heading\n\nsome text below\n"),
		},
	}

	t.Log("tst")
	for _, test := range tests {
		t.Run(string(test.Input), func(t *testing.T) {
			testParser := parser.New()
			parsed := testParser.Parse(test.Input)
			opts := RendererOptions{}
			ren := NewRenderer(opts)
			res := markdown.Render(parsed, ren)
			t.Log(string(test.Input), "--", string(test.Output), "---->", string(res))
			if !bytes.Equal(test.Output, res) {
				t.Errorf("input: %q\nexpected: %q\nactual:   %q", string(test.Input), string(test.Output), string(res))
			}
			fmt.Println(string(res))
		})
	}

}
