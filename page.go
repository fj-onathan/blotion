package src

import (
	"encoding/json"
	"fmt"
	"github.com/fj-onathan/blotion/config"
	"github.com/fj-onathan/blotion/vars"
	"log"
	"strings"
)

// ! + Export struct to JSON
type Page struct {
	Fields map[string]interface{} `json:"fields"`
}

func pageChunk(PageID string) map[string]vars.PageChunk {
	reqChunk := vars.ChunkRequest{
		PageID:          config.ToDashID(PageID),
		Limit:           9999,
		ChunkNumber:     0,
		VerticalColumns: false,
	}

	bytesReq, err := json.Marshal(reqChunk)
	if err != nil {
		log.Fatalln(err)
	}
	Chunk := config.LoadPageChunk(bytesReq)
	return Chunk
}

func blockText(block [][]interface{}) string {
	var Element string
	if len(block) == 1 {
		Element = "<p>" + block[0][0].(string) + "</p>"
	} else {
		Element = "<p>"
		for i := 0; i < len(block); i++ {
			if len(block[i]) > 1 {
				decorationProperty := fmt.Sprintf("%v", block[i][1])
				replacer := strings.NewReplacer("[[", "", "]]", "")
				property := replacer.Replace(decorationProperty)
				arguments := strings.Split(property, " ")

				if config.Contains(arguments[0], "b", "u", "i") {
					Element += "<" + arguments[0] + ">" + block[i][0].(string) + "</" + arguments[0] + ">"
				} else if config.Contains(arguments[0], "a") {
					Element += "<" + arguments[0] + " href='" + arguments[1] + "' target='_blank'>" + block[i][0].(string) + "</" + arguments[0] + ">"
				} else {
					Element += block[i][0].(string)
				}
			} else {
				Element += block[i][0].(string)
			}
		}
		Element += "</p>"
	}
	return Element
}

func ExportPage(PageID string) string {
	// Request to Record Values
	Records := getRecords(PageID)
	Chunk := pageChunk(PageID)
	var HTML string

	for _, block := range Records.Value.Content {
		Block := Chunk["recordMap"].Block[block].Values
		if config.Contains(Block.Type, "sub_header", "text") {
			if Block.Type == "text" {
				HTML += blockText(Block.Properties["title"])
			}
		} else if config.Contains(Block.Type, "image") {
			// TODO: Encode
			HTML += "<img src='" + Block.Format["display_source"].(string) + "'>"
		}

	}
	return HTML
}
