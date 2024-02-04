package lexer

// TokenType represents the type of a token.

type TokenType int

const (
	EOF     TokenType = -1
	NEWLINE TokenType = 0
	NUMBER  TokenType = 1
	IDENT   TokenType = 2
	STRING  TokenType = 3
	// Keywords.
	LABEL    TokenType = 101
	GOTO     TokenType = 102
	PRINT    TokenType = 103
	INPUT    TokenType = 104
	LET      TokenType = 105
	IF       TokenType = 106
	THEN     TokenType = 107
	ENDIF    TokenType = 108
	WHILE    TokenType = 109
	REPEAT   TokenType = 110
	ENDWHILE TokenType = 111
	// Operators.
	EQ       TokenType = 201
	PLUS     TokenType = 202
	MINUS    TokenType = 203
	ASTERISK TokenType = 204
	SLASH    TokenType = 205
	EQEQ     TokenType = 206
	NOTEQ    TokenType = 207
	LT       TokenType = 208
	LTEQ     TokenType = 209
	GT       TokenType = 210
	GTEQ     TokenType = 211
)

// TokenInfo es una estructura que contiene información asociada con cada tipo de token.
type TokenInfo struct {
	TokenType TokenType
	Name      string
}

// Resto del código ...

// TokenInfoMap es un mapa que asocia cada TokenType con su información correspondiente.
var TokenInfoMap = map[TokenType]TokenInfo{
	EOF:     {EOF, "EOF"},
	NEWLINE: {NEWLINE, "NEWLINE"},
	NUMBER:  {NUMBER, "NUMBER"},
	IDENT:   {IDENT, "IDENT"},
	STRING:  {STRING, "STRING"},
	// Palabras clave.
	LABEL:    {LABEL, "LABEL"},
	GOTO:     {GOTO, "GOTO"},
	PRINT:    {PRINT, "PRINT"},
	INPUT:    {INPUT, "INPUT"},
	LET:      {LET, "LET"},
	IF:       {IF, "IF"},
	THEN:     {THEN, "THEN"},
	ENDIF:    {ENDIF, "ENDIF"},
	WHILE:    {WHILE, "WHILE"},
	REPEAT:   {REPEAT, "REPEAT"},
	ENDWHILE: {ENDWHILE, "ENDWHILE"},
	// Operadores.
	EQ:       {EQ, "EQ"},
	PLUS:     {PLUS, "PLUS"},
	MINUS:    {MINUS, "MINUS"},
	ASTERISK: {ASTERISK, "ASTERISK"},
	SLASH:    {SLASH, "SLASH"},
	EQEQ:     {EQEQ, "EQEQ"},
	NOTEQ:    {NOTEQ, "NOTEQ"},
	LT:       {LT, "LT"},
	LTEQ:     {LTEQ, "LTEQ"},
	GT:       {GT, "GT"},
	GTEQ:     {GTEQ, "GTEQ"},
}

// String devuelve una representación en cadena del tipo de token.
func (t TokenInfo) String() string {
	return t.Name
}
