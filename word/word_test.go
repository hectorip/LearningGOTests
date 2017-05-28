package word

import "testing"

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
		t.Error(`IsPalindrome("eté") = false`)
	}
}

func TestPalindromeWithPunctuation(t *testing.T){
	input_words := "A man, a plan, a canal: Panamá"
	if !IsPalindrome(input_words) {
		t.Errorf(`IsPalindrome(%q) = false`, input_words)
	}
}
