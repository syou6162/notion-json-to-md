package main

import (
	"testing"
)

func TestConvertEmptyResults(t *testing.T) {
	input := NotionResponse{
		Object:  "list",
		Results: []Block{},
	}

	got := Convert(input)
	want := ""

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertHeading1(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "heading_1",
				Heading1: &Heading{
					RichText: []RichText{
						{
							PlainText: "Test Heading 1",
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "# Test Heading 1\n\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertHeading2(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "heading_2",
				Heading2: &Heading{
					RichText: []RichText{
						{
							PlainText: "Test Heading 2",
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "## Test Heading 2\n\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertHeading3(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "heading_3",
				Heading3: &Heading{
					RichText: []RichText{
						{
							PlainText: "Test Heading 3",
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "### Test Heading 3\n\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertParagraph(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "paragraph",
				Paragraph: &Paragraph{
					RichText: []RichText{
						{
							PlainText: "Test paragraph text",
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "Test paragraph text\n\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertBulletedListItem(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "bulleted_list_item",
				BulletedListItem: &ListItem{
					RichText: []RichText{
						{
							PlainText: "Test bullet item",
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "- Test bullet item\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertNumberedListItem(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "numbered_list_item",
				NumberedListItem: &ListItem{
					RichText: []RichText{
						{
							PlainText: "Test numbered item",
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "1. Test numbered item\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertCodeBlock(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "code",
				Code: &CodeBlock{
					Language: "go",
					RichText: []RichText{
						{
							PlainText: "func main() {}",
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "```go\nfunc main() {}\n```\n\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertBoldAnnotation(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "paragraph",
				Paragraph: &Paragraph{
					RichText: []RichText{
						{
							PlainText: "bold text",
							Annotations: Annotations{
								Bold: true,
							},
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "**bold text**\n\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertItalicAnnotation(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "paragraph",
				Paragraph: &Paragraph{
					RichText: []RichText{
						{
							PlainText: "italic text",
							Annotations: Annotations{
								Italic: true,
							},
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "*italic text*\n\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertCodeAnnotation(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "paragraph",
				Paragraph: &Paragraph{
					RichText: []RichText{
						{
							PlainText: "code text",
							Annotations: Annotations{
								Code: true,
							},
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "`code text`\n\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertStrikethroughAnnotation(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "paragraph",
				Paragraph: &Paragraph{
					RichText: []RichText{
						{
							PlainText: "strike text",
							Annotations: Annotations{
								Strikethrough: true,
							},
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "~~strike text~~\n\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertMultipleAnnotations(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "paragraph",
				Paragraph: &Paragraph{
					RichText: []RichText{
						{
							PlainText: "bold italic text",
							Annotations: Annotations{
								Bold:   true,
								Italic: true,
							},
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "***bold italic text***\n\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertLink(t *testing.T) {
	href := "https://example.com"
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "paragraph",
				Paragraph: &Paragraph{
					RichText: []RichText{
						{
							PlainText: "link text",
							Href:      &href,
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "[link text](https://example.com)\n\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertMultipleRichTexts(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "paragraph",
				Paragraph: &Paragraph{
					RichText: []RichText{
						{
							PlainText: "Normal ",
						},
						{
							PlainText: "bold",
							Annotations: Annotations{
								Bold: true,
							},
						},
						{
							PlainText: " text",
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "Normal **bold** text\n\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestConvertMultipleBlocks(t *testing.T) {
	input := NotionResponse{
		Object: "list",
		Results: []Block{
			{
				Type: "heading_1",
				Heading1: &Heading{
					RichText: []RichText{
						{
							PlainText: "Title",
						},
					},
				},
			},
			{
				Type: "paragraph",
				Paragraph: &Paragraph{
					RichText: []RichText{
						{
							PlainText: "First paragraph",
						},
					},
				},
			},
			{
				Type: "bulleted_list_item",
				BulletedListItem: &ListItem{
					RichText: []RichText{
						{
							PlainText: "Item 1",
						},
					},
				},
			},
			{
				Type: "bulleted_list_item",
				BulletedListItem: &ListItem{
					RichText: []RichText{
						{
							PlainText: "Item 2",
						},
					},
				},
			},
		},
	}

	got := Convert(input)
	want := "# Title\n\nFirst paragraph\n\n- Item 1\n- Item 2\n"

	if got != want {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}
