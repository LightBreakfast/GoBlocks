package blocks

import (
	"encoding/json"
	"fmt"
)

//Root - Starting Block
type Root struct {
	Blocks []Blocks `json:"blocks"`
}

//Text - (Without Inline Type Definitions)
type Text struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

//Accessory - (Without Inline Type Definitions)
type Accessory struct {
	Type        string       `json:"type,omitempty"`
	ImageURL    string       `json:"image_url,omitempty"`
	AltText     string       `json:"alt_text,omitempty"`
	Options     *[]Options   `json:"options,omitempty"`
	PlaceHolder *PlaceHolder `json:"placeholder,omitempty"`
}

//OptionsText - (Without Inline Type Definitions)
type OptionsText struct {
	Type  string `json:"type,omitempty"`
	Text  string `json:"text,omitempty"`
	Emoji bool   `json:"emoji,omitempty"`
}

//Options - (Without Inline Type Definitions)
type Options struct {
	Value       string      `json:"value,omitempty"`
	OptionsText OptionsText `json:"text,omitempty"`
}

//Elements - (Without Inline Type)
type Elements struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

//PlaceHolder - (Without Inline Type)
type PlaceHolder struct {
	Type  string `json:"type,omitempty"`
	Text  string `json:"text,omitempty"`
	Emoji bool   `json:"emoji,omitempty"`
}

//Blocks - (Without Inline Type Definitions)
type Blocks struct {
	Type      string      `json:"type"`
	Text      *Text       `json:"text,omitempty"`
	BlockID   string      `json:"block_id,omitempty"`
	Accessory *Accessory  `json:"accessory,omitempty"`
	Elements  *[]Elements `json:"elements,omitempty"`
}

//GenerateSection - Generates A Standard Text Section
func GenerateSection(text string) (section Blocks) {
	Section := Blocks{Type: "section",
		Text: &Text{
			Text: text,
			Type: "mrkdwn",
		},
	}

	return Section
}

//GenerateSectionWithImage - Generates A Blocks Section with Inline Image
func GenerateSectionWithImage(text string, imageurl string, alttext string) (section Blocks) {

	Section := Blocks{Type: "section",
		Text: &Text{
			Text: text,
			Type: "mrkdwn",
		},
		Accessory: &Accessory{
			Type:     "image",
			ImageURL: imageurl,
			AltText:  alttext,
		},
	}

	return Section

}

//GenerateDivder - Generates a divider block
func GenerateDivder() (section Blocks) {
	Divider := Blocks{Type: "divider"}

	return Divider
}

//GenerateContext - Generates a Context Block
func GenerateContext(text string) (cntext Blocks) {

	Cntext := Blocks{Type: "context",
		Elements: &[]Elements{
			Elements{
				Text: text,
				Type: "mrkdwn",
			},
		},
	}

	return Cntext

}

//GenerateBlocksString - Generates string representation of a slice of blocks
func GenerateBlocksString(blocks []Blocks) (blockstr string) {

	root := Root{blocks}

	prettyJSON, err := json.MarshalIndent(root, "", "    ")
	if err != nil {
		fmt.Println("Failed to generate json", err)
	}

	return string(prettyJSON)

}

//GenerateOverFlow - This Generates an Overflow block from a Slice Of Strings - (Emoji is always on)
func GenerateOverFlow(items []string, title string) (options Blocks) {

	var o []Options

	for i, item := range items {
		o = append(o, Options{
			Value: fmt.Sprintf("value-%v", i),
			OptionsText: OptionsText{
				Type:  "plain_text",
				Text:  item,
				Emoji: true,
			},
		},
		)
	}

	Section := Blocks{Type: "section",
		Text: &Text{
			Text: title,
			Type: "mrkdwn",
		},
		Accessory: &Accessory{
			Type:    "overflow",
			Options: &o,
		},
	}

	return Section

}

//GenerateMultiSelect - This Generates an MultiSelect block from a Slice Of Strings - (Emoji is always on)
func GenerateMultiSelect(items []string, text string, phtext string) (options Blocks) {

	var o []Options

	for i, item := range items {
		o = append(o, Options{
			Value: fmt.Sprintf("value-%v", i),
			OptionsText: OptionsText{
				Type:  "plain_text",
				Text:  item,
				Emoji: true,
			},
		},
		)
	}

	Section := Blocks{Type: "section",
		Text: &Text{
			Text: text,
			Type: "mrkdwn",
		},
		Accessory: &Accessory{
			Type: "multi_static_select",
			PlaceHolder: &PlaceHolder{
				Type:  "plain_text",
				Text:  phtext,
				Emoji: true,
			},
			Options: &o,
		},
	}

	return Section

}
