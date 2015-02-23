package dbmetamodel

type ForeignKey struct {
	DatabaseName string
	SchemaName   string
	TableName    string
	KeyName      string
	FkColumns    []string
	UkColumns    []string
	UpdateRule   string
	DeleteRule   string
	MatchOption  string
}
