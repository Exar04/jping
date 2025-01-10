package lexer

import "jping/token"

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
	for l.ch == '\n' {
		l.readChar()
	}

	var tok token.Token
	switch l.ch {
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: string(l.ch)}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: string(l.ch)}
	case ':':
		tok = token.Token{Type: token.COLON, Literal: string(l.ch)}
	case '"':
		l.readChar()
		tok.Literal = l.readString()
		tok.Type = token.STRING
	case ',':
		tok = token.Token{Type: token.LBRACE, Literal: string(l.ch)}
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readString() string {
	var str string
	for l.ch != '"' {
		str += string(l.ch)
		l.readChar()
	}
	return str
}
