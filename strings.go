package main

func main() {}

// Erwartet einen String s und einen Buchstaben c.
// Prüft, ob c in s vorkommt.
func Contains(s string, c byte) bool {
	for i := range s {
		if s[i] == c {
			return true
		}
	}
	return false
}

// Erwartet einen String s und einen Buchstaben c.
// Zählt, wie oft c in s vorkommt.
func CountChar(s string, c byte) int {
	count := 0
	for i := range s {
		if s[i] == c {
			count++
		}
	}
	return count
}

// Erwartet einen String s und einen Buchstaben c.
// Liefert die Position, an der c in s vorkommt.
// Liefert die Länge von s, falls c nicht in s vorkommt.
// Kommt c mehrfach vor, soll die erste Position geliefert werden.
func PositionOf(s string, c byte) int {
	for i := range s {
		if s[i] == c {
			return i
		}
	}
	return len(s)
}

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(s, t string) bool {
	for i := 0; i < len(s)-len(t)+1; i++ {
		if s[i:i+len(t)] == t {
			return true
		}
	}
	return false
}

// Erwartet einen String und prüft, ob er korrekte Klammerpaare enthält.
// D.h. der Eingabestring enthält Klammern '(' und ')'.
// Die Funktion soll nun prüfen, ob der String für jede öffnende Klammer auch eine
// schließende Klammer enthält.
// Außerdem darf es keine schließenden Klammern geben, für die es nicht vorher eine
// passende öffnende Klammer gegeben hat.
// Die Funktion soll true liefern, falls der String korrekt geklammert ist.
func CheckParentheses(s string) bool {
	counter := 0
	for _, c := range s {
		if c == '(' {
			counter++
		}
		if c == ')' {
			counter--
			if counter < 0 {
				return false
			}
		}
	}
	return counter == 0
}

// Erwartet zwei Strings s und sep sowie eine Zahl n.
// Liefert einen String, der aus n Kopien von s besteht, die duch sep getrennt werden.
func ConcatN(s, sep string, n int) string {
	result := ""
	for i := 0; i < n-1; i++ {
		result += s + sep
	}
	return result + s
}

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip(s1, s2 string) string {
	result := ""

	min := len(s1)
	if len(s2) < min {
		min = len(s2)
	}
	for i := 0; i < min; i++ {
		result += string(s1[i]) + string(s2[i])
	}
	result += s1[min:]
	result += s2[min:]

	return result
}
