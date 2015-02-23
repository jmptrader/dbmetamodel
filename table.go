package dbmetamodel

type Table struct {
	DatabaseName string
	SchemaName   string
	TableName    string
	Columns      []*Column
	ForeignKeys  []*ForeignKey

	InsertColumns     []*Column
	UpdateColumns     []*Column
	PrimaryKeyColumns []*Column
	HasPrimaryKeys    bool

	SerialPrimaryKey          string
	SerialPrimaryKeyFieldName string
	SerialPrimaryKeyFieldType string

	ModelName                string
	ModelNameLowerCase       string
	ModelNameLowerCasePlural string
}
