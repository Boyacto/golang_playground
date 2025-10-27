package lexer

type Lexer struct {
	input        string
	currentIndex int  // current index in input (points to current char)
	nextIndex    int  // current reading index in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextIndex >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextIndex]
	}
	l.currentIndex = l.nextIndex
	l.nextIndex += 1
}
