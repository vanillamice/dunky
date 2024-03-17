package token

type TokenType string

type Token struct{
	Type TokenType
	Literal string
}

const(
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//Identifiers + Literals
	IDENT = "IDENT" //add,foo,x,y
	INT = "INT" //1,2,3,41,0

	//Operators
	ASSIGN = "="
	PLUS = "+"

	//Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)

