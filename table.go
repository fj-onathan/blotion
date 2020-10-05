package src

import (
	"encoding/json"
	"github.com/fj-onathan/blotion/config"
	"github.com/fj-onathan/blotion/vars"
	"log"
	"net/url"
	"strings"
)

// ! + Export struct to JSON
type Table struct {
	Fields map[string]interface{} `json:"fields"`
}

func getRecords(PageID string) vars.RecordValues {
	reqRecord := vars.RecordRequest{
		Requests: []vars.Page{
			{
				ID:    config.ToDashID(PageID),
				Table: "block",
			},
		},
	}
	bytesReq, err := json.Marshal(reqRecord)
	if err != nil {
		log.Fatalln(err)
	}
	record := config.GetRecordValues(bytesReq)
	return record["results"][0]
}

func getCollection(CollectionID, CollectionViewID string) map[string]vars.QuerySelection {
	reqCollection := vars.QueryValues{
		CollectionID:     CollectionID,
		CollectionViewID: CollectionViewID,
		Loader: vars.Loader{
			Type: "table",
		},
	}
	bytesReq, err := json.Marshal(reqCollection)
	if err != nil {
		log.Fatalln(err)
	}
	ConvertData := config.GetQueryCollection(bytesReq)

	return ConvertData
}

func ExportTable(PageID string) []Table {
	// Request to Record Values
	Records := getRecords(PageID)
	CollectionID := Records.Value.CollectionID
	CollectionViewID := Records.Value.ViewIds[0]

	// Request to Query Collection
	CollectionData := getCollection(CollectionID, CollectionViewID)

	var Schema = CollectionData["recordMap"].Collection[CollectionID].Values.Schema
	TableExport := make([]Table, len(CollectionData["result"].BlockIDS))

	for tb, block := range CollectionData["result"].BlockIDS {
		Page := CollectionData["recordMap"].Block[block]
		Fields := make(map[string]interface{})

		for id, _ := range Schema {
			SchemaDefenition := Schema[id]
			SHtype := SchemaDefenition.Type
			SHvalue := Page.Values.Properties

			// ! add to fields
			if len(SHvalue[id]) > 0 {
				var value interface{}
				value = SHvalue[id][0][0].(string)

				if SHtype == "multi_select" {
					c := strings.Split(value.(string), ",")
					value = c
				} else if SHtype == "file" {
					fileID := Page.Values.Files[len(Page.Values.Files)-1]
					fileURL := url.QueryEscape("https://s3-us-west-2.amazonaws.com/secure.notion-static.com/"+fileID+"/"+value.(string)) + "?table=block&id=" + Page.Values.ID
					value = "https://www.notion.so/image/" + fileURL
				}

				Fields[SchemaDefenition.Name] = value
			}
		}

		TableExport[tb].Fields = Fields
	}

	return TableExport
}
