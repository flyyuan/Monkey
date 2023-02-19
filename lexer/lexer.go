package lexer

import (
	"github.com/flyyuan/Monkey/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	// create a new lexer
	l := &Lexer{input: input}
	// read one char to initialize the lexer
	// this is necessary because the lexer does not have a current char until it reads one
	// the lexer will not be able to read the first char in the input without this
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// if we have reached the end of the input, set the current char to 0
	if l.readPosition >= len(l.input) {
		// 0 is the ASCII code for the null character
		l.ch = 0
	} else {
		// otherwise, set the current char to the next char in the input
		// readPosition is the index of the next char to be read
		// l.readPosition defaults to 0 as it is not defined explicitly, so the first char read will be at index 0
		l.ch = l.input[l.readPosition]
	}
	// increment the position counters
	l.position = l.readPosition
	// increment the read position
	l.readPosition += 1
}

// NextToken returns the next token in the input
// it does this by reading the current char and then calling the appropriate function to handle the char
// the function will then return the token
// the lexer will then read the next char and repeat the process
// this is the main function of the lexer
// it is called by the parser to get the next token
// the parser will then use the token to determine what to do next
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// if the current char is not a special char, it could be part of an identifier
		// identifiers are strings of letters and numbers
		// if the current char is a letter, it is part of an identifier
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			// check if the identifier is a keyword
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// readIdentifier reads the current char and then reads the next char until it reaches a non-letter char
// it then returns the string of letters that it read
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// isLetter returns true if the given char is a letter
// it is used to determine if the current char is part of an identifier
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// newToken creates a new token with the given type and literal
// it is used to create tokens for single character tokens
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
