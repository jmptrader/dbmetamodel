package dbmetamodel

type Column struct {
	ColumnName            string
	IsPrimaryKeyColumn    bool
	IsAutoIncrementColumn bool
	AutoIncrementDataType string
	IsNullable            bool
	HasDefault            bool
	IsUnique              bool
	DataType              string
	DataTypeSize          int64

	FieldName          string
	FieldNameLowerCase string
	FieldType          string
}
