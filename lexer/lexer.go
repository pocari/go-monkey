package lexer

import (
	"bytes"
	"fmt"
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置(現在の文字)
	readPosition int  // これから読み込む位置(現在の文字の次)
	ch           byte // 現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}
func (l *Lexer) dump() string {
	return fmt.Sprintf(`
input: %s
position: %d
readPosition: %d
ch: %q`, l.input, l.position, l.readPosition, string(l.ch))
}

func (l *Lexer) NextToken() *token.Token {
	var tok *token.Token

	l.skipWhitespace()

	switch l.ch {
	case '"':
		tok = &token.Token{
			Type:    token.STRING,
			Literal: l.readString(),
		}
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = &token.Token{
				Type:    token.EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '[':
		tok = newToken(token.LBRACKET, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = &token.Token{
				Type:    token.NOT_EQ,
				Literal: literal,
			}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok = newEofToken()
	default:
		if isLetter(l.ch) {
			return newIdentiferToken(l.readIdentifier())
		} else if isDigit(l.ch) {
			return newIntToken(l.readNumber())
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func newToken(tokenType token.TokenType, ch byte) *token.Token {
	return &token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func newEofToken() *token.Token {
	return &token.Token{
		Type:    token.EOF,
		Literal: "",
	}
}

func newIntToken(digits string) *token.Token {
	return &token.Token{
		Type:    token.INT,
		Literal: digits,
	}
}

func newIdentiferToken(literal string) *token.Token {
	return &token.Token{
		Type:    token.LookupIdent(literal),
		Literal: literal,
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// end of input
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition > len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readEscapeChar(position int, buffer *bytes.Buffer) int {
	l.readChar()
	switch l.ch {
	case 'a':
		buffer.WriteString(fmt.Sprintf("%s\a", l.input[position:l.position-1]))
	case 'b':
		buffer.WriteString(fmt.Sprintf("%s\b", l.input[position:l.position-1]))
	case 'f':
		buffer.WriteString(fmt.Sprintf("%s\f", l.input[position:l.position-1]))
	case 'n':
		buffer.WriteString(fmt.Sprintf("%s\n", l.input[position:l.position-1]))
	case 'r':
		buffer.WriteString(fmt.Sprintf("%s\r", l.input[position:l.position-1]))
	case 't':
		buffer.WriteString(fmt.Sprintf("%s\t", l.input[position:l.position-1]))
	case 'v':
		buffer.WriteString(fmt.Sprintf("%s\v", l.input[position:l.position-1]))
	case '\\':
		buffer.WriteString(fmt.Sprintf("%s\\", l.input[position:l.position-1]))
	case '\'':
		buffer.WriteString(fmt.Sprintf("%s'", l.input[position:l.position-1]))
	case '"':
		buffer.WriteString(fmt.Sprintf("%s\"", l.input[position:l.position-1]))
	default:
		// それ以外はバックスラッシュ無視してその文字自体にする
		buffer.WriteString(fmt.Sprintf("%s%s", l.input[position:l.position-1], string(l.ch)))
	}
	return l.position + 1
}

func (l *Lexer) readString() string {
	var b bytes.Buffer
	p := l.position + 1

exit_loop:
	for {
		l.readChar()
		switch l.ch {
		case '\\':
			p = l.readEscapeChar(p, &b)
		default:
			if l.ch == '"' || l.ch == 0 {
				break exit_loop
			}
		}
	}
	b.WriteString(l.input[p:l.position])
	return b.String()
}
