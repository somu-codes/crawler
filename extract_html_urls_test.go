package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
    <html>
        <body>
            <a href="/path/one">Boot.dev</a>
            <a href="https://other.com/path/one">Other</a>
        </body>
    </html>
`,
			expected: []string{
				"https://blog.boot.dev/path/one",
				"https://other.com/path/one",
			},
		},
		{
			name:     "only relative URLs",
			inputURL: "https://example.com",
			inputBody: `
    <html>
        <body>
            <a href="/foo">Foo</a>
            <a href="/bar">Bar</a>
        </body>
    </html>
`,
			expected: []string{
				"https://example.com/foo",
				"https://example.com/bar",
			},
		},
		{
			name:     "only absolute URLs",
			inputURL: "https://test.com",
			inputBody: `
    <html>
        <body>
            <a href="https://a.com/x">A</a>
            <a href="https://b.com/y">B</a>
        </body>
    </html>
`,
			expected: []string{
				"https://a.com/x",
				"https://b.com/y",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getURLsFromHTML(tt.inputBody, tt.inputURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v, expected %v", got, tt.expected)
			}
		})
	}
}
