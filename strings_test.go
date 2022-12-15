package main

import "fmt"

func ExampleContains() {
	fmt.Println(Contains("Hallo", 'a'))
	fmt.Println(Contains("Hallo", 'b'))
	fmt.Println(Contains("Hallo", 'H'))
	fmt.Println(Contains("Hallo", 'h'))

	// Output:
	// true
	// false
	// true
	// false
}

func ExampleCountChar() {
	fmt.Println(CountChar("Hallo", 'a'))
	fmt.Println(CountChar("Hallo", 'l'))
	fmt.Println(CountChar("Hallo", 'o'))
	fmt.Println(CountChar("Hallo", 'x'))
	fmt.Println(CountChar("", 'a'))

	// Output:
	// 1
	// 2
	// 1
	// 0
	// 0

}

func ExamplePositionOf() {
	fmt.Println(PositionOf("Hallo", 'H'))
	fmt.Println(PositionOf("Hallo", 'l'))
	fmt.Println(PositionOf("Hallo", 'o'))
	fmt.Println(PositionOf("Hallo", 'x'))
	fmt.Println(PositionOf("", 'a'))

	// Output:
	// 0
	// 2
	// 4
	// 5
	// 0
}

func ExampleContainsSubstring() {
	fmt.Println(ContainsSubstring("Hallo", "Ha"))
	fmt.Println(ContainsSubstring("Hallo", "llo"))
	fmt.Println(ContainsSubstring("Hallo", "Ho"))

	// Output:
	// true
	// true
	// false
}

func ExampleCheckParentheses() {
	fmt.Println(CheckParentheses("()"))
	fmt.Println(CheckParentheses("(())"))
	fmt.Println(CheckParentheses("()()"))
	fmt.Println(CheckParentheses("(a)b(ac)cb"))
	fmt.Println(CheckParentheses("())"))
	fmt.Println(CheckParentheses("(()"))
	fmt.Println(CheckParentheses("()("))

	// Output:
	// true
	// true
	// true
	// true
	// false
	// false
	// false
}

func ExampleConcatN() {
	fmt.Println(ConcatN("A", "", 5))
	fmt.Println(ConcatN("A", "B", 5))
	fmt.Println(ConcatN("AB", "C", 5))

	// Output:
	// AAAAA
	// ABABABABA
	// ABCABCABCABCAB
}

func ExampleZip() {
	fmt.Println(Zip("AAA", "BBB"))
	fmt.Println(Zip("ABABCD", "CDCDAB"))

	// Output:
	// ABABAB
	// ACBDACBDCADB
}

func ExampleZip_different_length() {
	fmt.Println(Zip("A", "BBB"))
	fmt.Println(Zip("AAA", "B"))

	// Output:
	// ABBB
	// ABAA
}
