package dbmetamodel

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

type LimitQueryPart struct {
	LimitTo uint64
}

type SkipQueryPart struct {
	SkipOver uint64
}

type OrderQueryPartItem struct {
	FieldName string
	Direction string
}

type OrderQueryPart struct {
	OrderBy []*OrderQueryPartItem
}

type WhereQueryPart struct {
	FieldName string
	Operator  string
	Value     string
}

type FieldsQueryPart struct {
	Included []string
	Excluded []string
}

type IncludeQueryPart struct {
}

type FindByIdQuery struct {
	ModelId string

	Fields *FieldsQueryPart
}

type FindOneQuery struct {
	Fields  *FieldsQueryPart
	Include []*IncludeQueryPart
	Where   []*WhereQueryPart
	OrderBy *OrderQueryPart
	Skip    *SkipQueryPart
}

type FindQuery struct {
	Fields  *FieldsQueryPart
	Include []*IncludeQueryPart
	Where   []*WhereQueryPart
	OrderBy *OrderQueryPart
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

type DestroyByIdQuery struct {
	ModelId string
}

type DestroyQuery struct {
	Where []*WhereQueryPart
}

type SqlQueryBuilder interface {
	BuildFindByIdQuery(query FindByIdQuery, tableMetaData *Table) (string, error)

	BuildFindQuery(query FindQuery, tableMetaData *Table) (string, error)

	BuildFindOneQuery(query FindQuery, tableMetaData *Table) (string, error)

	BuildCountQuery(query CountQuery, tableMetaData *Table) (string, error)

	BuildExistsQuery(query ExistsQuery, tableMetaData *Table) (string, error)

	BuildInsertQuery(query InsertQuery, tableMetaData *Table) (string, error)

	BuildDestroyByIdQuery(query DestroyByIdQuery, tableMetaData *Table) (string, error)

	BuildDestroyQuery(query DestroyQuery, tableMetaData *Table) (string, error)

	BuildUpdateByIdQuery(query UpdateByIdQuery, tableMetaData *Table) (string, error)

	BuildUpdateQuery(query UpdateQuery, tableMetaData *Table) (string, error)
}

type QueryUrlParser interface {
	ParseFindByIdQueryString(findByIdQuery *FindByIdQuery, queryString url.Values) error

	ParseFindQueryString(findQuery *FindQuery, queryString url.Values) error

	ParseFindOneQueryString(findOneQuery *FindOneQuery, queryString url.Values) error
}

type FilterUrlParser interface {
	ParseUrlFieldsFilterPart(filterString string, filterValueString string) (fields *FieldsQueryPart, err error)

	ParseUrlLimitFilterPart(filterString string, filterValueString string) (limit *LimitQueryPart, err error)

	ParseUrlSkipFilterPart(filterString string, filterValueString string) (skip *SkipQueryPart, err error)

	ParseUrlOrderFilterPart(filterString string, filterValueString string) (order *OrderQueryPart, err error)

	ParseUrlWhereFilterPart(filterString string, filterValueString string) (where []*WhereQueryPart, err error)
}

type QueryStringUrlParser struct {
	filterTypeRegex    *regexp.Regexp
	limitFilterRegex   *regexp.Regexp
	skipFilterRegex    *regexp.Regexp
	orderByFilterRegex *regexp.Regexp
	orderByValueRegex  *regexp.Regexp
	whereFilterRegex   *regexp.Regexp
}

func NewQueryStringUrlParser() *QueryStringUrlParser {
	parser := QueryStringUrlParser{
		filterTypeRegex:    regexp.MustCompile("^(?i)\\s*filter\\s*\\[\\s*(?P<filterType>\\w*)\\s*\\]\\s*$"),
		limitFilterRegex:   regexp.MustCompile("^(?i)\\s*filter\\s*\\[\\s*limit\\s*\\]\\s*$"),
		skipFilterRegex:    regexp.MustCompile("^(?i)\\s*filter\\s*\\[\\s*skip\\s*\\]\\s*$"),
		orderByFilterRegex: regexp.MustCompile("^(?i)\\s*filter\\s*\\[\\s*order\\s*\\]\\s*$"),
		orderByValueRegex:  regexp.MustCompile("(?i)(?P<fieldName>\\w*)\\s*(?P<direction>ASC|DESC)"),
		whereFilterRegex:   regexp.MustCompile("^(?i)\\s*filter\\s*\\[\\s*where\\s*\\]\\s*"),
		//^(?i)\s*filter\s*\[\s*where\s*\]\s*\[\s*(?<fieldName>\w*)\s*\]\s*\[\s*(?<operation>\w*)\s*\]\s*$
	}

	return &parser
}

func getFilterType(regex *regexp.Regexp, queryStringPart string) (isFilter bool, filterType string) {
	// check if the queryStringPart is a filter
	if regex.MatchString(queryStringPart) {
		n1 := regex.SubexpNames()
		r2 := regex.FindAllStringSubmatch(queryStringPart, -1)[0]

		md := map[string]string{}
		for i, n := range r2 {
			md[n1[i]] = n
		}

		// return the filter type found
		filterType := md["filterType"]

		return true, strings.ToLower(filterType)
	}
	return false, ""
}

func (parser *QueryStringUrlParser) ParseFindByIdQueryString(findByIdQuery *FindByIdQuery, queryString url.Values) (err error) {
	for queryStringPart, values := range queryString {
		// get the filter type
		isFilter, filterType := getFilterType(parser.filterTypeRegex, queryStringPart)

		if isFilter {
			for _, value := range values {
				switch filterType {
				case "fields":
					findByIdQuery.Fields, err = parser.ParseUrlFieldsFilterPart(queryStringPart, value)
					break
				default:
					return errors.New(filterType + " is not a valid filter for a findByIdQuery")
				}
			}
		}
	}

	return nil
}

func (parser *QueryStringUrlParser) ParseFindQueryString(findQuery *FindQuery, queryString url.Values) (err error) {
	for queryStringPart, values := range queryString {
		// get the filter type
		isFilter, filterType := getFilterType(parser.filterTypeRegex, queryStringPart)

		if isFilter {
			fmt.Println(filterType)
			for _, value := range values {
				switch filterType {
				case "fields":
					findQuery.Fields, err = parser.ParseUrlFieldsFilterPart(queryStringPart, value)
					break
				case "where":
					findQuery.Where, err = parser.ParseUrlWhereFilterPart(queryStringPart, value)
					break
				case "order":
					findQuery.OrderBy, err = parser.ParseUrlOrderFilterPart(queryStringPart, value)
					break
				case "limit":
					findQuery.Limit, err = parser.ParseUrlLimitFilterPart(queryStringPart, value)
					break
				case "skip":
					findQuery.Skip, err = parser.ParseUrlSkipFilterPart(queryStringPart, value)
					break
				default:
					return errors.New(filterType + " is not a valid filter for a findQuery")
				}
			}
		}
	}

	return nil
}

func (parser *QueryStringUrlParser) ParseFindOneQueryString(findOneQuery *FindOneQuery, queryString url.Values) error {
	return nil
}

func (parser *QueryStringUrlParser) ParseUrlFieldsFilterPart(filterString string, filterValueString string) (fields *FieldsQueryPart, err error) {
	return nil, nil
}

func (parser *QueryStringUrlParser) ParseUrlLimitFilterPart(filterString string, filterValueString string) (limit *LimitQueryPart, err error) {
	if parser.limitFilterRegex.MatchString(filterString) {

		// we have a limit query part
		limitTo, err := strconv.ParseUint(filterValueString, 10, 64)

		if err != nil {
			return nil, errors.New("Error parsing the value for the limit filter, it is not a valid integer.\n" + err.Error())
		}
		limit = &LimitQueryPart{
			LimitTo: limitTo,
		}

		return limit, nil
	}

	return nil, nil
}

func (parser *QueryStringUrlParser) ParseUrlSkipFilterPart(filterString string, filterValueString string) (skip *SkipQueryPart, err error) {

	if parser.skipFilterRegex.MatchString(filterString) {

		// we have a skip query part
		skipBy, err := strconv.ParseUint(filterValueString, 10, 64)
		if err != nil {
			return nil, err
		}
		skip = &SkipQueryPart{
			SkipOver: skipBy,
		}

		return skip, nil
	}

	return nil, nil
}

func (parser *QueryStringUrlParser) ParseUrlOrderFilterPart(filterString string, filterValueString string) (order *OrderQueryPart, err error) {
	if parser.orderByFilterRegex.MatchString(filterString) {

		// we have an order query part
		if parser.orderByValueRegex.MatchString(filterValueString) {
			n1 := parser.orderByValueRegex.SubexpNames()
			r2 := parser.orderByValueRegex.FindAllStringSubmatch(filterValueString, -1)

			orderItems := []*OrderQueryPartItem{}
			for _, v := range r2 {

				md := map[string]string{}
				for i, n := range v {
					md[n1[i]] = n
				}

				fieldName := md["fieldName"]
				direction := md["direction"]

				fmt.Println(md)

				orderItems = append(orderItems, &OrderQueryPartItem{
					FieldName: fieldName,
					Direction: strings.ToLower(direction),
				})
			}

			if len(orderItems) > 0 {
				order = &OrderQueryPart{
					OrderBy: orderItems,
				}
			}
		}

		return order, nil
	}

	return nil, nil
}

func (parser *QueryStringUrlParser) ParseUrlWhereFilterPart(filterString string, filterValueString string) (where []*WhereQueryPart, err error) {
	return nil, nil
}
