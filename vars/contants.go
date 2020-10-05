package vars

// HTTP
const (
	NotionHost    = "https://www.notion.so/api/"
	NotionVersion = "v3"
	NotionHTTP    = NotionHost + NotionVersion
	// modern Chrome
	UserAgent   = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3483.0 Safari/537.36"
	AcceptLang  = "en-US,en;q=0.9"
	ContentType = "application/json"
)

// + for User.Permissions
const (
	// PermissionTypeUser describes permissions for a user
	PermissionTypeUser = "user_permission"
	// PermissionTypePublic describes permissions for public
	PermissionTypePublic = "public_permission"
)

// + for Schema.Type
const (
	ColumnTypeCheckbox       = "checkbox"
	ColumnTypeCreatedBy      = "created_by"
	ColumnTypeCreatedTime    = "created_time"
	ColumnTypeDate           = "date"
	ColumnTypeEmail          = "email"
	ColumnTypeFile           = "file"
	ColumnTypeForumula       = "formula"
	ColumnTypeLastEditedBy   = "last_edited_by"
	ColumnTypeLastEditedTime = "last_edited_time"
	ColumnTypeMultiSelect    = "multi_select"
	ColumnTypeNumber         = "number"
	ColumnTypePerson         = "person"
	ColumnTypePhoneNumber    = "phone_number"
	ColumnTypeRelation       = "relation"
	ColumnTypeRollup         = "rollup"
	ColumnTypeSelect         = "select"
	ColumnTypeText           = "text"
	ColumnTypeTitle          = "title"
	ColumnTypeURL            = "url"
)
