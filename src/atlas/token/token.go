package token

//type of token(string can hold many forms of types)
type TokenType string

//struct for the Token(has a Type and value of the Token)
type Token struct{
	Type TokenType
	Value string
}

//type of TokenTypes
const (
	//special cases
	ILLEGAL = "ILLEGAL" //Unknown character 
	EOF = "EOF"

	//Identifiers & Value
	IDENT = "IDENT" // function_name, variable_name...
	INT = "INT" //values of integers

	//Operators
	ASSIGN = "="
	PLUS = "+"

	//Delimiter
	COMMA = ","
	SEMICOLON = ";"
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	
	//Keyword
	FUNCTION = "FUNCTION"
	LET = "LET"
)

