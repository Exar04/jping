package parser

import (
	"fmt"
	"jping/lexer"
	"jping/token"
	"strconv"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func (p *Parser) ParseJson() map[string]interface{} {
	values := make(map[string]interface{})

	if p.curToken.Type != token.LBRACE {
		return nil
	}
	p.nextToken()
	p.parseKeyVal(&values)
	// fmt.Println(values)

	return values
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseKeyVal(values *map[string]interface{}) {
	switch p.curToken.Type {
	// case token.LBRACE:
	// nextedValues := make(map[string]interface{})
	// p.parseKeyVal(&nextedValues)

	case token.COMMA:
		p.nextToken()
		p.parseKeyVal(values)
	case token.STRING:
		key := p.curToken.Literal
		if p.peekToken.Type != token.COLON {
			fmt.Println("Error while parsing! Expected COLON got ", p.peekToken)
			return
		}
		p.nextToken() // now we have colon as curtoken
		p.nextToken() // now we must be having value as currtoken
		if p.curToken.Type == token.BOOL {
			if p.curToken.Literal == "true" {
				(*values)[key] = true
			} else {
				(*values)[key] = false
			}
		} else if p.curToken.Type == token.INT {
			i, e := strconv.Atoi(p.curToken.Literal)
			if e == nil {
				(*values)[key] = i
			}
		} else if p.curToken.Type == token.LBRACK {
			p.nextToken() // skipping the left bracket [
			var getArray func() interface{}
			getArray = func() interface{} {
				if p.curToken.Type == token.INT {
					return p.createIntArray()
				} else if p.curToken.Type == token.BOOL {
					return p.createBoolArray()
				} else if p.curToken.Type == token.STRING {
					return p.createStrArray()
				} else if p.curToken.Type == token.LBRACK {
					p.nextToken()
					var ark []interface{}
					arr := getArray()
					p.nextToken() // skipping ] bracket
					ark = append(ark, arr)

					// fmt.Println(ark)

					for p.curToken.Type == token.COMMA {
						// fmt.Println(p.curToken)
						p.nextToken()
						if p.curToken.Type == token.LBRACK {
							p.nextToken()
						} else {
							fmt.Println("Expected [ got ", p.curToken.Literal)
						}
						arr := getArray()
						p.nextToken()
						ark = append(ark, arr)
					}
					return ark
				} else if p.curToken.Type == token.RBRACK {
					p.nextToken()
				}
				return nil
			}
			(*values)[key] = getArray()
		} else if p.curToken.Type == token.LBRACE {
			nextedValues := make(map[string]interface{})
			p.nextToken()
			p.parseKeyVal(&nextedValues)
			(*values)[key] = nextedValues
		} else {
			(*values)[key] = p.curToken.Literal
		}
		if p.peekToken.Type != token.RBRACE {
			p.nextToken()
			p.parseKeyVal(values)
		}
	}
	return
}

func (p *Parser) createIntArray() []int {
	var arr []int
	for p.curToken.Type != token.RBRACK {
		i, e := strconv.Atoi(p.curToken.Literal)
		if e == nil {
			arr = append(arr, i)
		}
		p.nextToken()
		if p.curToken.Type == token.COMMA {
			if p.peekToken.Type != token.INT {
				fmt.Println("error while parsing array. Expected Integer token got ", p.peekToken)
				return nil
			}
			p.nextToken()
		}
	}
	return arr
}

func (p *Parser) createBoolArray() []bool {
	var arr []bool
	for p.curToken.Type != token.RBRACK {
		if p.curToken.Literal == "true" {
			arr = append(arr, true)
		} else {
			arr = append(arr, false)
		}
		p.nextToken()
		if p.curToken.Type == token.COMMA {
			if p.peekToken.Type != token.BOOL {
				fmt.Println("error while parsing array. Expected Integer token got ", p.peekToken)
				return nil
			}
			p.nextToken()
		}
	}
	return arr
}

func (p *Parser) createStrArray() []string {
	var arr []string
	for p.curToken.Type != token.RBRACK {
		arr = append(arr, p.curToken.Literal)
		p.nextToken()
		if p.curToken.Type == token.COMMA {
			if p.peekToken.Type != token.STRING {
				fmt.Println("error while parsing array. Expected Integer token got ", p.peekToken)
				return nil
			}
			p.nextToken()
		}
	}
	return arr
}
