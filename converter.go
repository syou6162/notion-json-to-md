package main

type NotionResponse struct {
	Object  string  `json:"object"`
	Results []Block `json:"results"`
}

type Block struct {
	Type             string     `json:"type"`
	Heading1         *Heading   `json:"heading_1,omitempty"`
	Heading2         *Heading   `json:"heading_2,omitempty"`
	Heading3         *Heading   `json:"heading_3,omitempty"`
	Paragraph        *Paragraph `json:"paragraph,omitempty"`
	BulletedListItem *ListItem  `json:"bulleted_list_item,omitempty"`
	NumberedListItem *ListItem  `json:"numbered_list_item,omitempty"`
	Code             *CodeBlock `json:"code,omitempty"`
}

type Heading struct {
	RichText []RichText `json:"rich_text"`
}

type Paragraph struct {
	RichText []RichText `json:"rich_text"`
}

type ListItem struct {
	RichText []RichText `json:"rich_text"`
}

type CodeBlock struct {
	Language string     `json:"language"`
	RichText []RichText `json:"rich_text"`
}

type RichText struct {
	PlainText   string      `json:"plain_text"`
	Annotations Annotations `json:"annotations"`
	Href        *string     `json:"href"`
}

type Annotations struct {
	Bold          bool `json:"bold"`
	Italic        bool `json:"italic"`
	Strikethrough bool `json:"strikethrough"`
	Code          bool `json:"code"`
}

func formatRichText(richTexts []RichText) string {
	var result string
	for _, rt := range richTexts {
		text := rt.PlainText

		// Apply annotations
		if rt.Annotations.Code {
			text = "`" + text + "`"
		}
		if rt.Annotations.Bold {
			text = "**" + text + "**"
		}
		if rt.Annotations.Italic {
			text = "*" + text + "*"
		}
		if rt.Annotations.Strikethrough {
			text = "~~" + text + "~~"
		}

		// Apply link
		if rt.Href != nil && *rt.Href != "" {
			text = "[" + text + "](" + *rt.Href + ")"
		}

		result += text
	}
	return result
}

func Convert(response NotionResponse) string {
	if len(response.Results) == 0 {
		return ""
	}

	var result string
	for _, block := range response.Results {
		switch block.Type {
		case "heading_1":
			if block.Heading1 != nil && len(block.Heading1.RichText) > 0 {
				result += "# " + formatRichText(block.Heading1.RichText) + "\n\n"
			}
		case "heading_2":
			if block.Heading2 != nil && len(block.Heading2.RichText) > 0 {
				result += "## " + formatRichText(block.Heading2.RichText) + "\n\n"
			}
		case "heading_3":
			if block.Heading3 != nil && len(block.Heading3.RichText) > 0 {
				result += "### " + formatRichText(block.Heading3.RichText) + "\n\n"
			}
		case "paragraph":
			if block.Paragraph != nil && len(block.Paragraph.RichText) > 0 {
				result += formatRichText(block.Paragraph.RichText) + "\n\n"
			}
		case "bulleted_list_item":
			if block.BulletedListItem != nil && len(block.BulletedListItem.RichText) > 0 {
				result += "- " + formatRichText(block.BulletedListItem.RichText) + "\n"
			}
		case "numbered_list_item":
			if block.NumberedListItem != nil && len(block.NumberedListItem.RichText) > 0 {
				result += "1. " + formatRichText(block.NumberedListItem.RichText) + "\n"
			}
		case "code":
			if block.Code != nil && len(block.Code.RichText) > 0 {
				result += "```" + block.Code.Language + "\n"
				result += formatRichText(block.Code.RichText) + "\n"
				result += "```\n\n"
			}
		}
	}
	return result
}
