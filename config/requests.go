package config

import (
	"encoding/json"
	"github.com/fj-onathan/blotion/vars"
)

// + Request to Initial Values
// ? /api/v3/getRecordValues
func GetRecordValues(params []byte) map[string][]vars.RecordValues {
	path := "/getRecordValues"
	data := notionAPI(params, path)

	dataJSON := make(map[string][]vars.RecordValues)
	err := json.Unmarshal(data, &dataJSON)
	if err != nil {
		panic(err)
	}

	return dataJSON
}

// + Request to Query Collection
// ? /api/v3/queryCollection
func GetQueryCollection(params []byte) map[string]vars.QuerySelection {
	path := "/queryCollection"
	data := notionAPI(params, path)

	ConvertData := make(map[string]vars.QuerySelection)
	err := json.Unmarshal(data, &ConvertData)
	if err != nil {
		panic(err)
	}

	return ConvertData
}

// + Request to Page Chunk
// ? /api/v3/loadPageChunk
func LoadPageChunk(params []byte) map[string]vars.PageChunk {
	path := "/loadPageChunk"
	data := notionAPI(params, path)

	ConvertData := make(map[string]vars.PageChunk)
	err := json.Unmarshal(data, &ConvertData)
	if err != nil {
		panic(err)
	}

	return ConvertData
}
