package dbmetamodel

import (
	"bitbucket.org/pkg/inflect"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

type Tables struct {
	TableCatalog string `db:"table_catalog"`
	TableSchema  string `db:"table_schema"`
	TableName    string `db:"table_name"`
	TableType    string `db:"table_type"`
}

type Columns struct {
	TableCatalog           string         `db:"table_catalog"`
	TableSchema            string         `db:"table_schema"`
	TableName              string         `db:"table_name"`
	ColumnName             string         `db:"column_name"`
	OrdinalPosition        sql.NullInt64  `db:"ordinal_position"`
	ColumnDefault          sql.NullString `db:"column_default"`
	IsNullable             sql.NullString `db:"is_nullable"`
	DataType               string         `db:"data_type"`
	CharacterMaximumLength sql.NullInt64  `db:"character_maximum_length"`
	CharacterOctetLength   sql.NullInt64  `db:"character_octet_length"`
	NumericPrecision       sql.NullInt64  `db:"numeric_precision"`
	NumericPrecisionRadix  sql.NullInt64  `db:"numeric_precision_radix"`
	NumericScale           sql.NullInt64  `db:"numeric_scale"`
	DatetimePrecision      sql.NullInt64  `db:"datetime_precision"`
	IntervalType           sql.NullString `db:"interval_type"`
	IntervalPrecision      sql.NullInt64  `db:"interval_precision"`
	CharacterSetCatalog    sql.NullString `db:"character_set_catalog"`
	CharacterSetSchema     sql.NullString `db:"character_set_schema"`
	CharacterSetName       sql.NullString `db:"character_set_name"`
	CollationCatalog       sql.NullString `db:"collation_catalog"`
	CollationSchema        sql.NullString `db:"collation_schema"`
	CollationName          sql.NullString `db:"collation_name"`
	DomainCatalog          sql.NullString `db:"domain_catalog"`
	DomainSchema           sql.NullString `db:"domain_schema"`
	DomainName             sql.NullString `db:"domain_name"`
	UdtCatalog             sql.NullString `db:"udt_catalog"`
	UdtSchema              sql.NullString `db:"udt_schema"`
	UdtName                sql.NullString `db:"udt_name"`
	ScopeCatalog           sql.NullString `db:"scope_catalog"`
	ScopeSchema            sql.NullString `db:"scope_schema"`
	ScopeName              sql.NullString `db:"scope_name"`
	MaximumCardinality     sql.NullInt64  `db:"maximum_cardinality"`
	DtdIdentifier          sql.NullString `db:"dtd_identifier"`
	IsSelfReferencing      sql.NullString `db:"is_self_referencing"`
	IsIdentity             sql.NullString `db:"is_identity"`
	IdentityGeneration     sql.NullString `db:"identity_generation"`
	IdentityStart          sql.NullString `db:"identity_start"`
	IdentityIncrement      sql.NullString `db:"identity_increment"`
	IdentityMaximum        sql.NullString `db:"identity_maximum"`
	IdentityMinimum        sql.NullString `db:"identity_minimum"`
	IdentityCycle          sql.NullString `db:"identity_cycle"`
	IsGenerated            sql.NullString `db:"is_generated"`
	GenerationExpression   sql.NullString `db:"generation_expression"`
	IsUpdatable            sql.NullString `db:"is_updatable"`
}

type KeyColumnUsage struct {
	ConstraintCatalog          string        `db:"constraint_catalog"`
	ConstraintSchema           string        `db:"constraint_schema"`
	ConstraintName             string        `db:"constraint_name"`
	TableCatalog               string        `db:"table_catalog"`
	TableSchema                string        `db:"table_schema"`
	TableName                  string        `db:"table_name"`
	ColumnName                 string        `db:"column_name"`
	ConstraintType             string        `db:"constraint_type"`
	OrdinalPosition            int           `db:"ordinal_position"`
	PositionInUniqueConstraint sql.NullInt64 `db:"position_in_unique_constraint"`
}

type TableConstraints struct {
	ConstraintCatalog string `db:"constraint_catalog"`
	ConstraintSchema  string `db:"constraint_schema"`
	ConstraintName    string `db:"constraint_name"`
	TableCatalog      string `db:"table_catalog"`
	TableSchema       string `db:"table_schema"`
	TableName         string `db:"table_name"`
	ConstraintType    string `db:"constraint_type"`
	IsDeferrable      string `db:"is_deferrable"`
	InitiallyDeferred string `db:"initially_deferred"`
}

type TableForeignKeys struct {
	FkTableCatalog    string `db:"fk_table_catalog"`
	FkTableSchema     string `db:"fk_table_schema_name"`
	FkTableName       string `db:"fk_table_name"`
	FkConstraintName  string `db:"fk_constraint_name"`
	FkColumnName      string `db:"fk_column_name"`
	FkOrdinalPosition int    `db:"fk_ordinal_position"`
	UqTableCatalog    string `db:"uq_table_catalog"`
	UqTableSchema     string `db:"uq_table_schema_name"`
	UqTableName       string `db:"uq_table_name"`
	UqConstraintName  string `db:"uq_constraint_name"`
	UqColumnName      string `db:"uq_column_name"`
	UqOrdinalPosition int    `db:"uq_ordinal_position"`
	UpdateRule        string `db:"update_rule"`
	DeleteRule        string `db:"delete_rule"`
	MatchOption       string `db:"match_option"`
}

type SchemaLoader interface {
	GetTables(databaseName string, schemaName string) ([]Tables, error)

	GetColumns(databaseName string, schemaName string, tableName string) ([]Columns, error)

	GetColumnKeyUsage(databaseName string, schemaName string, tableName string, columnName string) ([]KeyColumnUsage, error)

	GetForeignKeys(databaseName string, schemaName string, tableName string) ([]TableForeignKeys, error)
}

type schemaLoader struct {
	ConnectionString string
}

func (lodr *schemaLoader) GetTables(databaseName string, schemaName string) ([]Tables, error) {
	db, err := sqlx.Connect("postgres", lodr.ConnectionString)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer db.Close()

	tables := []Tables{}

	err = db.Select(&tables, "select table_catalog, table_schema, table_name, table_type from information_schema.tables where table_catalog = $1 and table_schema = $2 and table_type = 'BASE TABLE'", databaseName, schemaName)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return tables, err
}

func (lodr *schemaLoader) GetColumns(databaseName string, schemaName string, tableName string) ([]Columns, error) {

	db, err := sqlx.Connect("postgres", lodr.ConnectionString)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer db.Close()

	columns := []Columns{}
	err = db.Select(&columns, "select * from information_schema.columns where table_catalog = $1 and table_schema = $2 and table_name = $3 order by ordinal_position", databaseName, schemaName, tableName)

	return columns, err
}

func (lodr *schemaLoader) GetColumnKeyUsage(databaseName string, schemaName string, tableName string, columnName string) ([]KeyColumnUsage, error) {
	db, err := sqlx.Connect("postgres", lodr.ConnectionString)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer db.Close()

	keyUsage := []KeyColumnUsage{}
	err = db.Select(&keyUsage, `select kcu.constraint_catalog, kcu.constraint_schema, kcu.constraint_name, kcu.table_catalog, kcu.table_schema, kcu.table_name, kcu.column_name, tc.constraint_type, kcu.ordinal_position, kcu.position_in_unique_constraint from information_schema.key_column_usage kcu 
								join information_schema.table_constraints tc on tc.constraint_catalog = kcu.constraint_catalog and tc.constraint_schema = kcu.constraint_schema and tc.constraint_name = kcu.constraint_name
								where kcu.table_catalog = $1 and kcu.table_schema = $2 and kcu.table_name = $3 and kcu.column_name = $4`, databaseName, schemaName, tableName, columnName)

	return keyUsage, err
}

func (lodr *schemaLoader) GetForeignKeys(databaseName string, schemaName string, tableName string) ([]TableForeignKeys, error) {
	db, err := sqlx.Connect("postgres", lodr.ConnectionString)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	defer db.Close()

	foreignKeys := []TableForeignKeys{}
	err = db.Select(&foreignKeys, `select 
		     kcu1.table_catalog as fk_table_catalog
		   , kcu1.table_schema as fk_table_schema_name
		   , kcu1.table_name as fk_table_name
		   , kcu1.constraint_name as fk_constraint_name   
		   , kcu1.column_name as fk_column_name
		   , kcu1.ordinal_position as fk_ordinal_position
		   , kcu2.table_catalog as uq_table_catalog
		   , kcu2.table_schema as uq_table_schema_name		   
		   , kcu2.table_name as uq_table_name
		   , kcu2.constraint_name as uq_constraint_name		   
		   , kcu2.column_name as uq_column_name
		   , kcu2.ordinal_position as uq_ordinal_position
		   , rc.update_rule
		   , rc.delete_rule
		   , rc.match_option
		from information_schema.referential_constraints rc
		join information_schema.key_column_usage kcu1
		on kcu1.constraint_catalog = rc.constraint_catalog 
		   and kcu1.constraint_schema = rc.constraint_schema
		   and kcu1.constraint_name = rc.constraint_name
		join information_schema.key_column_usage kcu2
		on kcu2.constraint_catalog = 
		rc.unique_constraint_catalog 
		   and kcu2.constraint_schema = 
		rc.unique_constraint_schema
		   and kcu2.constraint_name = 
		rc.unique_constraint_name
		   and kcu2.ordinal_position = kcu1.ordinal_position

		   and kcu1.table_catalog = $1 and kcu1.table_schema = $2 and kcu1.table_name = $3
		   order by uq_table_catalog, uq_table_schema_name, uq_table_name, uq_constraint_name, uq_column_name, uq_ordinal_position`, databaseName, schemaName, tableName)

	return foreignKeys, err
}

type postgresqlSchemaProcessor struct {
	schemaLoader SchemaLoader
}

func NewPostgresqlSchemaProcessor(schemaLoader SchemaLoader) SchemaProcessor {
	schemaProcessor := &postgresqlSchemaProcessor{
		schemaLoader: schemaLoader,
	}

	return schemaProcessor
}

func (proc *postgresqlSchemaProcessor) populateFk(table *Table, fk TableForeignKeys) (currentFk *ForeignKey) {
	var matchOption string
	if fk.MatchOption == "NONE" {
		matchOption = "SIMPLE"
	} else {
		matchOption = fk.MatchOption
	}

	currentFk = &ForeignKey{
		DatabaseName: fk.UqTableCatalog,
		SchemaName:   fk.UqTableSchema,
		TableName:    fk.UqTableName,
		KeyName:      fk.FkConstraintName,
		FkColumns: []string{
			fk.FkColumnName,
		},
		UkColumns: []string{
			fk.UqColumnName,
		},
		UpdateRule:  fk.UpdateRule,
		DeleteRule:  fk.DeleteRule,
		MatchOption: matchOption,
	}

	return currentFk
}

func (proc *postgresqlSchemaProcessor) RetrieveTableMetaData(databaseName string, schemaName string) ([]Table, error) {

	// fetch the schema tables
	dbTables, err := proc.schemaLoader.GetTables(databaseName, schemaName)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	tables := []Table{}
	for _, dbTable := range dbTables {

		sanitisedModelName := toCapitalCase(dbTable.TableName)
		sanitisedModelNameLowerCase := strings.ToLower(sanitisedModelName)
		pluralisedModelName := inflect.Pluralize(sanitisedModelNameLowerCase)

		table := Table{
			DatabaseName:             dbTable.TableCatalog,
			SchemaName:               dbTable.TableSchema,
			TableName:                dbTable.TableName,
			ModelName:                sanitisedModelName,
			ModelNameLowerCase:       sanitisedModelNameLowerCase, // sanitised tableName
			ModelNameLowerCasePlural: pluralisedModelName,         // pluralised sanitised tableName
			Columns:                  []*Column{},
			ForeignKeys:              []*ForeignKey{},
			InsertColumns:            []*Column{},
			UpdateColumns:            []*Column{},
			PrimaryKeyColumns:        []*Column{},
		}

		// fetch the schema columns
		dbColumns, err := proc.schemaLoader.GetColumns(databaseName, schemaName, dbTable.TableName)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}

		// get the foreign key information
		var currentFk *ForeignKey
		dbForeignKeys, err := proc.schemaLoader.GetForeignKeys(databaseName, schemaName, dbTable.TableName)
		for _, fk := range dbForeignKeys {
			if currentFk == nil {
				currentFk = proc.populateFk(&table, fk)
				table.ForeignKeys = append(table.ForeignKeys, currentFk)
			} else {
				if currentFk.DatabaseName == fk.UqTableCatalog && currentFk.SchemaName == fk.UqTableSchema && currentFk.TableName == fk.UqTableName {
					currentFk.FkColumns = append(currentFk.FkColumns, fk.FkColumnName)
					currentFk.UkColumns = append(currentFk.UkColumns, fk.UqColumnName)
				} else {
					currentFk = proc.populateFk(&table, fk)
					table.ForeignKeys = append(table.ForeignKeys, currentFk)
				}
			}
		}

		for _, dbColumn := range dbColumns {

			// get the primay and unique key column usage
			keyColumns, err := proc.schemaLoader.GetColumnKeyUsage(databaseName, schemaName, dbTable.TableName, dbColumn.ColumnName)
			if err != nil {
				log.Fatalln(err)
				return nil, err
			}

			sanitisedColumnName := toCapitalCase(dbColumn.ColumnName)
			sanitisedColumnNameLowerCase := strings.ToLower(sanitisedColumnName)

			column := Column{

				ColumnName:         dbColumn.ColumnName,
				FieldName:          sanitisedColumnName,
				FieldNameLowerCase: sanitisedColumnNameLowerCase,
				FieldType:          dataType(dbColumn.DataType),
				DataType:           dbColumn.DataType,
				IsNullable:         false,
				IsPrimaryKeyColumn: false,
				HasDefault:         false,
			}

			if dbColumn.CharacterMaximumLength.Valid {
				column.DataTypeSize = dbColumn.CharacterMaximumLength.Int64
			}

			if dbColumn.IsNullable.Valid && dbColumn.IsNullable.String == "YES" {
				column.IsNullable = true
			}

			if dbColumn.ColumnDefault.Valid {
				column.HasDefault = true
			}

			for _, colKey := range keyColumns {
				if colKey.ConstraintType == "PRIMARY KEY" {
					column.IsPrimaryKeyColumn = true
					table.PrimaryKeyColumns = append(table.PrimaryKeyColumns, &column)
				}

				if colKey.ConstraintType == "UNIQUE" {
					column.IsUnique = true
				}
			}

			// assumption about sequence columns on postgres, if it has a default, is not nullable and is a primary key assume it is a sequence
			if column.HasDefault && column.IsPrimaryKeyColumn && !column.IsNullable {
				column.IsAutoIncrementColumn = true

				if column.DataType == "integer" {
					column.AutoIncrementDataType = "serial"
					table.SerialPrimaryKeyFieldType = "int"
				} else if column.DataType == "bigint" {
					column.AutoIncrementDataType = "bigserial"
					table.SerialPrimaryKeyFieldType = "int64"
				}

				table.SerialPrimaryKey = column.ColumnName
				table.SerialPrimaryKeyFieldName = column.FieldName
			}

			table.Columns = append(table.Columns, &column)
			if !column.IsAutoIncrementColumn {
				table.InsertColumns = append(table.InsertColumns, &column)
				table.UpdateColumns = append(table.UpdateColumns, &column)
			}
		}

		if len(table.PrimaryKeyColumns) > 0 {
			table.HasPrimaryKeys = true
		}
		tables = append(tables, table)
	}

	return tables, err
}
