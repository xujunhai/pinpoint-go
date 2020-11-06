package sql

const (
	SEPARATOR     byte = ','
	SymbolReplace byte = '$'
	NumberReplace byte = '#'

	NextTokenNotExist   int = -1
	NormalizedSqlBuffer int = 32
)

var (
	NullObject = NormalizedSql{normalizedSql: "", parseParameter: ""}
)

type Parser struct {
}

func (Parser) normalizedSql(sql string) NormalizedSql {
	return NullObject
}
