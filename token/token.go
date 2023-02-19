package token

import "fmt"

// TokenType 用于区分语法单元
type TokenType string

// Token is a struct
type Token struct {
	Type    TokenType
	Literal string
}

// Token 类型
const (
	// ILLEGAL 无效的
	ILLEGAL = "ILLEGAL"
	// EOF 文件结束
	EOF = "EOF"

	// IDENT 标识符
	IDENT = "IDENT"
	// INT 整型
	INT = "INT"

	// 运算法
	ASSIGN = "="
	PLUS   = "+"

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIndent find the token type of the identifier
func LookupIdent(ident string) TokenType {
	fmt.Println("LookupIdent: ", ident)
	// if ident is in the keywords map, return the token type
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	//otherwise, return IDENT
	return IDENT
}
