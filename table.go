package main

import (
	"regexp"
	"strings"
)

type table struct {
	TableName string
	Columns   []tableColumn
}

type tableColumn struct {
	ColumnName string
	ColumnType string
	NotNull    bool
}

type tableParser struct {
	tableRegex  *regexp.Regexp
	columnRegex *regexp.Regexp
}

func newTableParser() *tableParser {
	tableRegex, err := regexp.Compile(`(CREATE TABLE)(.*)(\((.|\n)*?)\);`)
	if err != nil {
		panic(err.Error())
	}

	columnRegex, err := regexp.Compile(`(.\n|\(\n)((.|\n)*)`)
	if err != nil {
		panic(err.Error())
	}

	return &tableParser{
		tableRegex:  tableRegex,
		columnRegex: columnRegex,
	}
}

func (t *tableParser) ParseString(input string) []table {
	tables := []table{}

	strTables := t.tableRegex.FindAllString(input, -1)

	for _, strTable := range strTables {
		subMatches := t.tableRegex.FindStringSubmatch(strTable)

		tableName := subMatches[2]
		tableName = strings.Trim(tableName, " ")
		tableName = strings.Trim(tableName, "\"")

		columnsString := subMatches[3]
		columnMatch := t.columnRegex.FindAllStringSubmatch(columnsString, -1)
		columns := strings.Split(columnMatch[0][2], ",\n")

		tableColumns := []tableColumn{}
		for _, column := range columns {
			tableColumn := constructColumn(column)
			tableColumns = append(tableColumns, tableColumn)
		}

		tables = append(tables, table{
			TableName: tableName,
			Columns:   tableColumns,
		})
	}

	return tables
}

func constructColumn(strColumn string) tableColumn {
	tableColumn := tableColumn{}

	column := strings.Trim(strColumn, " ")
	fields := strings.SplitN(column, " ", 3)

	if len(fields) == 2 {
		columnName := fields[0]
		columnType := fields[1]

		tableColumn.ColumnName = strings.Trim(columnName, "\"")
		tableColumn.ColumnType = columnType
	} else if len(fields) == 3 {
		columnName := fields[0]
		columnType := fields[1]
		columnDetail := fields[2]

		tableColumn.ColumnName = strings.Trim(columnName, "\"")
		tableColumn.ColumnType = columnType

		if strings.Contains(columnDetail, "NOT NULL") {
			tableColumn.NotNull = true
		}
	}

	return tableColumn
}
