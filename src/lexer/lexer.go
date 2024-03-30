package lexer

type Lexer struct {
	input        string
	position     int // current position
	readPosition int // after current position
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}

	return l
}
