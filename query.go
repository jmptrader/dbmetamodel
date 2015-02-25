package dbmetamodel

import (
	"net/url"
	"regexp"
	"strconv"
)

type LimitQueryPart struct {
	LimitTo uint64
}

type SkipQueryPart struct {
	SkipOver uint64
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

type FieldsQueryPart struct {
	Included []string
	Excluded []string
}
type FindQuery struct {
	Fields *FieldsQueryPart

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

type UrlParser interface {
	ParseUrlFieldsFilterPart(fields *FieldsQueryPart, queryString string) error

	ParseUrlLimitFilterPart(limit *LimitQueryPart, queryString string) error

	ParseUrlSkipFilterPart(skip *SkipQueryPart, queryString string) error

	ParseUrlOrderFilterPart(order *OrderQueryPart, queryString string) error

	ParseWhereFilterPart(where []*WhereQueryPart, queryString string) error

	//ParseIncludeFilterPart(queryString string)
}

type QueryStringUrlParser struct {
	limitRegex *regexp.Regexp
	skipRegex  *regexp.Regexp
}

func (parser *QueryStringUrlParser) ParseUrlFieldsFilterPart(fields *FieldsQueryPart, queryString url.Values) error {
	return nil
}

func (parser *QueryStringUrlParser) ParseUrlLimitFilterPart(limit *LimitQueryPart, queryString url.Values) error {
	for k, _ := range queryString {
		if parser.limitRegex.MatchString(k) {

			// we have a limit query part
			limitTo, err := strconv.ParseUint(queryString.Get(k), 10, 64)
			if err != nil {
				return err
			}
			limit = &LimitQueryPart{
				LimitTo: limitTo,
			}
			return nil
		}
	}

	return nil
}

func (parser *QueryStringUrlParser) ParseUrlSkipFilterPart(skip *SkipQueryPart, queryString url.Values) error {
	for k, _ := range queryString {
		if parser.skipRegex.MatchString(k) {

			// we have a skip query part
			skipBy, err := strconv.ParseUint(queryString.Get(k), 10, 64)
			if err != nil {
				return err
			}
			skip = &SkipQueryPart{
				SkipOver: skipBy,
			}
			return nil
		}
	}

	return nil
}

func (parser *QueryStringUrlParser) ParseUrlOrderFilterPart(order *OrderQueryPart, queryString url.Values) error {
	return nil
}

func (parser *QueryStringUrlParser) ParseWhereFilterPart(where []*WhereQueryPart, queryString url.Values) error {
	return nil
}
