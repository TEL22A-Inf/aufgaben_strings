package main

func main() {}

// Erwartet einen String s und einen Buchstaben c.
// Prüft, ob c in s vorkommt.
func Contains(s string, c byte) bool {
	// Hinweis: Prüfen Sie in einer Schleife für jedes Zeichen in s, ob es gleich c ist.
	// Falls ja, können Sie true liefern ("Early Out"), ansonsten liefern Sie am Ende false.
	// TODO
	return false
}

// Erwartet einen String s und einen Buchstaben c.
// Zählt, wie oft c in s vorkommt.
func CountChar(s string, c byte) int {
	count := 0
	// Hinweis: Gehen Sie wie bei Contains vor. Aber statt des Early-Out erhöhen Sie
	// bei jedem Vorkommen von c einen Zähler, den Sie am Ende liefern.
	// TODO
	return count
}

// Erwartet einen String s und einen Buchstaben c.
// Liefert die Position, an der c in s vorkommt.
// Liefert die Länge von s, falls c nicht in s vorkommt.
// Kommt c mehrfach vor, soll die erste Position geliefert werden.
func PositionOf(s string, c byte) int {
	// Hinweis: Gehen Sie wie bei Contains vor, aber führen Sie zusätzlich einen
	// Schleifenzähler mit, dessen Wert Sie im Fall einer Fundstelle von c liefern.
	// TODO
	return len(s)
}

// Erwartet zwei Strings s und t und prüft, ob t in s als Teilstring vorkommt.
func ContainsSubstring(s, t string) bool {
	// Hinweis: Verwenden Sie eine "Slice", um einen Teilstring von s zu gewinnen.
	// Beispiel: s[3:6] würde Ihnen die Zeichen aus s von 3 bis 5 liefern.
	// Vergleichen Sie solche Teilstrings in einer Schleife mit t.
	for i := 0; i < len(s)-len(t)+1; i++ {
		// TODO: Vergleichen Sie das s-Teilstück von i bis i+len(t) mit t.
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
	// Hinweis: Laufen Sie in einer Schleife über s.
	// Erhöhen Sie einen Klammer-Zähler um 1, wenn Sie eine öffnende Klammer sehen,
	// und verringern Sie ihn, wenn Sie eine schließende Klammer sehen.
	// Falls der Zähler jemals negativ wird oder am Ende nicht 0 ist,
	// liefern Sie false.
	counter := 0
	// TODO
	return counter == 0
}

// Erwartet zwei Strings s und sep sowie eine Zahl n.
// Liefert einen String, der aus n Kopien von s besteht, die duch sep getrennt werden.
func ConcatN(s, sep string, n int) string {
	// Hinweis: Schreiben Sie eine Schleife mit n-1 Durchläufen, die jedes Mal s und sep
	// an das Ergebnis anhängt. Am Ende müssen Sie noch ein weiteres Mal s anhängen.
	result := ""
	// TODO
	return result + s
}

// Erwartet zwei strings s1 und s2.
// Liefert einen String, der abwechselnd aus den Buchstaben des einen und des anderen
// Strings besteht.
func Zip(s1, s2 string) string {
	// Hinweis: Es ist Code vorgegeben, der die minimale Länge von s1 und s2 bestimmt.
	// Schreiben Sie eine Schleife, die bis zu diesem min läuft und in jedem Durchlauf
	// jeweils ein Zeichen aus s1 und eines aus s2 ans Ergebnis anhängt.
	// Hängen Sie anschließend noch die beiden restlichen Strings an.
	result := ""

	min := len(s1)
	if len(s2) < min {
		min = len(s2)
	}
	// TODO

	return result
}
