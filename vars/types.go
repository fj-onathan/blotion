package vars

// + Record values and params to interact to api/record
type RecordValues struct {
	Role  string `json:"role"`
	Value struct {
		ID           string   `json:"id"`
		Version      int      `json:"version"`
		Type         string   `json:"type"`
		ViewIds      []string `json:"view_ids"`
		CollectionID string   `json:"collection_id"`
		Content      []string `json:"content"`
		Permissions  []struct {
			Role   string `json:"role"`
			Type   string `json:"type"`
			UserID string `json:"user_id,omitempty"`
		} `json:"permissions"`
		CreatedTime       int64  `json:"created_time"`
		LastEditedTime    int64  `json:"last_edited_time"`
		ParentID          string `json:"parent_id"`
		ParentTable       string `json:"parent_table"`
		Alive             bool   `json:"alive"`
		CreatedByTable    string `json:"created_by_table"`
		CreatedByID       string `json:"created_by_id"`
		LastEditedByTable string `json:"last_edited_by_table"`
		LastEditedByID    string `json:"last_edited_by_id"`
	} `json:"value"`
}

type Page struct {
	ID    string `json:"id"`
	Table string `json:"table"`
}

type RecordRequest struct {
	Requests []Page `json:"requests"`
}

// + Query Selection and params to interact to api/record
type QueryValues struct {
	CollectionID     string `json:"collectionId"`
	CollectionViewID string `json:"collectionViewId"`
	Loader           Loader `json:"loader"`
}

type QuerySelection struct {
	Type               string                `json:"type"`
	BlockIDS           []string              `json:"blockIds"`
	AggregationResults []string              `json:"aggregationResults"`
	Total              int                   `json:"total"`
	Collection         map[string]Collection `json:"collection"`
	Block              map[string]Block      `json:"block"`
}

type Loader struct {
	Type string `json:"type"`
}

type Collection struct {
	Role   string `json:"role"`
	Values struct {
		Schema map[string]Schema `json:"schema"`
	} `json:"value"`
}

type Schema struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Block struct {
	Role   string `json:"role"`
	Values struct {
		ID         string                     `json:"id"`
		Version    int                        `json:"version"`
		Type       string                     `json:"type"`
		Properties map[string][][]interface{} `json:"properties"`
		Format     map[string]interface{}     `json:"format"`
		Files      []string                   `json:"file_ids"`
	} `json:"value"`
}

// + Page Chunk on and params to interact to api/record
type ChunkRequest struct {
	PageID          string `json:"pageId"`
	Limit           int    `json:"limit"`
	ChunkNumber     int    `json:"chunkNumber"`
	VerticalColumns bool   `json:"verticalColumns"`
}

type PageChunk struct {
	Type  string           `json:"type"`
	Block map[string]Block `json:"block"`
}

// + Block Structures
type BlockSubLevel struct {
	T []string
}
type BlockSubLevel2 []string
