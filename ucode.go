// different microcode operations
package edgeworth

type ConditionalOperation func(Word, Word) bool
type BinaryOperation func(Word, Word) Word
type TrinaryOperation func(Word, Word, Word) Word

func BitToWord(val Bit) Word {
	if val {
		return 1
	} else {
		return 0
	}
}

var Operations = []BinaryOperation{
	func(x, y Word) Word { return x + y },
	func(x, y Word) Word { return x - y },
	func(x, y Word) Word { return x * y },
	func(x, y Word) Word { return x / y },
	func(x, y Word) Word { return x % y },
	func(x, y Word) Word { return x << y },
	func(x, y Word) Word { return x >> y },
	func(x, y Word) Word { return x & y },
	func(x, y Word) Word { return x | y },
	func(x, y Word) Word { return x &^ y },
	func(x, _ Word) Word { return ^x },
	func(x, y Word) Word { return BitToWord(x == y) },
	func(x, y Word) Word { return BitToWord(x != y) },
	func(x, y Word) Word { return BitToWord(x > y) },
	func(x, y Word) Word { return BitToWord(x < y) },
	func(x, y Word) Word { return BitToWord(x >= y) },
	func(x, y Word) Word { return BitToWord(x <= y) },
}
