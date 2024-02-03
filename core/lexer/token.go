package lexer

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
