package lexer

import (
	symbol "lolang/token"
)

type Lexer struct {
	input        string
	cursor       int
	nextPosition int
	char         byte
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}

	lexer.readChar()

	return lexer
}

func (lexer *Lexer) nextToken() symbol.Token {
	var token symbol.Token

	lexer.eatWhitespace()

	if symbol.IsLetter(lexer.char) {
		token.Literal = lexer.readIdentifier()

		token.Type = symbol.LookupIdentifier(token.Literal)
	} else if symbol.IsDigit(lexer.char) {
		token.Literal = lexer.readNumber()

		token.Type = symbol.INTEGER
	} else if symbol.IsOperator(string(lexer.char)) {
		token.Literal = lexer.readOperator()

		token.Type = symbol.LookupOperator(token.Literal)
	} else if symbol.IsEOF(lexer.char) {
		token.Literal = ""

		token.Type = symbol.EOF

		lexer.readChar()
	} else {
		token.Literal = string(lexer.char)

		token.Type = symbol.ILLEGAL

		lexer.readChar()
	}

	return token
}

func (lexer *Lexer) eatWhitespace() {
	for lexer.char == ' ' || lexer.char == '\t' || lexer.char == '\n' || lexer.char == '\r' {
		lexer.readChar()
	}
}

func (lexer *Lexer) readIdentifier() string {
	cursor := lexer.cursor

	for symbol.IsLetter(lexer.char) {
		lexer.readChar()
	}

	return lexer.input[cursor:lexer.cursor]
}

func (lexer *Lexer) readNumber() string {
	cursor := lexer.cursor

	for symbol.IsDigit(lexer.char) {
		lexer.readChar()
	}

	return lexer.input[cursor:lexer.cursor]
}

func (lexer *Lexer) readOperator() string {
	cursor := lexer.cursor
	identifier := string(lexer.char)

	for symbol.IsOperator(identifier) {
		identifier += string(lexer.peek())

		lexer.readChar()
	}

	return lexer.input[cursor:lexer.cursor]
}

func (lexer *Lexer) readChar() {
	if lexer.nextPosition >= len(lexer.input) {
		lexer.char = 0
	} else {
		lexer.char = lexer.input[lexer.nextPosition]
	}

	lexer.cursor = lexer.nextPosition

	lexer.nextPosition += 1
}

func (lexer Lexer) peek() byte {
	if lexer.nextPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.nextPosition]
	}
}
