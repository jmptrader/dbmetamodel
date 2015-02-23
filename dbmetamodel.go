package dbmetamodel

type SchemaProcessor interface {
	RetrieveTableMetaData(databaseName string, schemaName string) ([]Table, error)
}
