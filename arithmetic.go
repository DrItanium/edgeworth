//arithmetic operations
package edgeworth

type ArithmeticOperation func(Word, Word) Word

var BasicArithmeticOperations = [OperationPerGroupMax]ArithmeticOperation{
	func(x, y Word) Word { return x + y },  // add
	func(x, y Word) Word { return x - y },  // sub
	func(x, y Word) Word { return x * y },  // mul
	func(x, y Word) Word { return x / y },  // div
	func(x, y Word) Word { return x % y },  // remainder
	func(x, y Word) Word { return x << y }, // shift left
	func(x, y Word) Word { return x >> y }, // shift right
	func(x, y Word) Word { return x & y },  // bitwise and
	func(x, y Word) Word { return x | y },  // bitwise or
	func(x, _ Word) Word { return ^x },     // bitwise not
	func(x, y Word) Word { return x &^ y }, //bitwise clear
}
