// different microcode operations
package edgeworth

type BinaryOperation func(Word, Word) Word
type TrinaryOperation func(Word, Word, Word) Word

var arithmeticOperations = []BinaryOperation{
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
	func(x, y Word) Word { return Word(x == y) },
}
