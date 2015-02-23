package dbmetamodel

type LimitQueryPart struct {
	LimitTo int
}

type SkipQueryPart struct {
	SkipOver int
}

type OrderQueryPart struct {
	FieldName string
	Direction string
}

type WhereQueryPart struct {
	FieldName string
	Operator  string
	Value     string
}

type FindByIdQuery struct {
	ModelId string
}

type FindQuery struct {
	Fields []string

	Where   []*WhereQueryPart
	OrderBy []*OrderQueryPart
	Limit   *LimitQueryPart
	Skip    *SkipQueryPart
}

type CountQuery struct {
	Where []*WhereQueryPart
}

type InsertQuery struct {
	Data map[string]string
}

type UpdateByIdQuery struct {
	ModelId string
	Data    map[string]string
}

type UpdateQuery struct {
	Data  map[string]string
	Where []*WhereQueryPart
}

type ExistsQuery struct {
	ModelId string
}

type DeleteByIdQuery struct {
	ModelId string
}

type QueryBuilder interface {
	BuildFindByIdQuery(query FindByIdQuery, tableMetaData *Table) (string, error)

	BuildFindQuery(query FindQuery, tableMetaData *Table) (string, error)

	BuildFindOneQuery(query FindQuery, tableMetaData *Table) (string, error)

	BuildCountQuery(query CountQuery, tableMetaData *Table) (string, error)

	BuildExistsQuery(query ExistsQuery, tableMetaData *Table) (string, error)

	BuildInsertQuery(query InsertQuery, tableMetaData *Table) (string, error)

	BuildDeleteByIdQuery(query DeleteByIdQuery, tableMetaData *Table) (string, error)

	BuildUpdateByIdQuery(query UpdateByIdQuery, tableMetaData *Table) (string, error)

	BuildUpdateQuery(query UpdateQuery, tableMetaData *Table) (string, error)
}

type UrlParser struct {
}

type BodyParser struct {
}
