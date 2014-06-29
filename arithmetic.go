// Arithmetic operations
package edgeworth

type ArithmeticOperation func(Word, Word) Word

var ArithmeticOperations = []ArithmeticOperation{
	func(x, _ Word) Word { return x },
	func(_, y Word) Word { return y },
	func(x, y Word) Word { return x + y },
	func(x, y Word) Word { return x - y },
	func(x, y Word) Word { return x * y },
	func(x, y Word) Word { return x / y },
	func(x, y Word) Word { return x % y },
	func(x, y Word) Word { return x << y },
	func(x, y Word) Word { return x >> y },
	func(x, y Word) Word { return x & y },
	func(x, y Word) Word { return x | y },
	func(x, _ Word) Word { return ^x },
	func(x, y Word) Word { return x ^ y },
}
