package lexer

import (
	"fmt"
	"jping/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	for l.ch == '\n' || l.ch == ' ' || l.ch == '\t' {
		l.readChar()
	}

	var tok token.Token
	switch l.ch {
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: string(l.ch)}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: string(l.ch)}
	case '[':
		tok = token.Token{Type: token.LBRACK, Literal: string(l.ch)}
	case ']':
		tok = token.Token{Type: token.RBRACK, Literal: string(l.ch)}
	case ':':
		tok = token.Token{Type: token.COLON, Literal: string(l.ch)}
	case '"':
		l.readChar()
		tok.Literal = l.readString()
		tok.Type = token.STRING
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: string(l.ch)}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.Keywords[tok.Literal]
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.realNumber()
			return tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
			fmt.Println("Got illigal token")
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) realNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	if isLetter(l.ch) {
		fmt.Println("A number shouldn't be as a first character in a identifier name")
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// func (l *Lexer) peekChar() byte {
// 	if l.readPosition >= len(l.input) {
// 		return 0
// 	} else {
// 		return l.input[l.readPosition]
// 	}
// }

func (l *Lexer) readString() string {
	var str string
	for l.ch != '"' {
		str += string(l.ch)
		l.readChar()
	}
	return str
}
