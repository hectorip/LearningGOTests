package word

import (
	"testing"
	"math/rand"
	"time"
)

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("detartrated") {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

func TestAccentedPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false`)
	}
}

func TestPalindromeWithPunctuation(t *testing.T){
	input_words := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input_words) {
		t.Errorf(`IsPalindrome(%q) = false`, input_words)
	}
}

func TestSeveralExamples(t *testing.T){
	var tests = []struct {
		input string
		result bool
	} {
		{"", true},
		{"a", true},
		{"aaa", true},
		{"aab", false}, // no es un palíndromo
		{"Anita lava la tina", true},
		{"yes", false}, // No es un palíndromo
	}
	for _, test := range tests {
		if !IsPalindrome(test.input) == test.result {
			t.Errorf(`Error IsPalindrome(%q) == %v`, test.input, test.result)
		}
	}
}

func getRandomPalindrome(rng *rand.Rand) (string, bool){
	n := rng.Intn(25) // Número aleatóreo hasta 24,
	runes := make([]rune, n)
	want_palindrome := rng.Intn(2) != 0 || n < 2 // 0 -> No palíndromo, 1 -> Palíndromo

	for i := 0; i < (n + 1) / 2 ; i++ {
		r :=  rune(rng.Intn(0x1000))
		runes[i] = r
		if want_palindrome {
			runes[ n - 1 - i ] = r
		} else {
			diff_r := rune(0x1000)
			runes[ n - 1 - i ] = diff_r
		}
	}
	if !want_palindrome  {
		runes = append(runes, rune(0x1001))
	}
	return string(runes), want_palindrome
}
func getRandomExampleWithPunctuationAndSpaces(rng *rand.Rand) (string, bool){
	example, is_palindrome := getRandomPalindrome(rng)
	// add spaces and punctuation
	special_char := string(rune(rng.Intn(0x0010) + 0x0020))
	example = example + special_char
	//fmt.Println(special_char)
	return example, is_palindrome
}
func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	//var palindrome string
	for i := 0; i < 1000; i++ {
		palindrome, expected_palindrome := getRandomExampleWithPunctuationAndSpaces(rng)
		if i % 2 == 0 {
			palindrome, expected_palindrome = getRandomPalindrome(rng)
		}
		if IsPalindrome(palindrome) != expected_palindrome {
			t.Errorf(`IsPalindrome(%q) != %v`, palindrome, expected_palindrome)
		}
	}
}