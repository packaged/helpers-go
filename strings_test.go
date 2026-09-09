package helpers

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncate(t *testing.T) {
	{
		str := "this long string"
		str = TruncateStringToRune(str, 9)
		if str != "this long" {
			t.Error("unexpected truncate")
		}
	}

	{
		str := "this😆emoji"
		str = TruncateStringToRune(str, 5)
		if str != "this" {
			t.Error("unexpected truncate")
		}
	}

	{
		str := "truncate longer"
		str = TruncateStringToRune(str, 30)
		if str != "truncate longer" {
			t.Error("unexpected truncate")
		}
	}
}

func TestBufferToString(t *testing.T) {
	str := "hello"
	result := BufferToString(strings.NewReader(str))
	if result != str {
		t.Errorf("BufferToString failed. Got: %s", result)
	}
}

func TestFirstStr(t *testing.T) {
	{
		result := FirstStr("hello", "world")
		if result != "hello" {
			t.Errorf("FirstStr failed. Got: %s", result)
		}
	}

	{
		result := FirstStr("", "world")
		if result != "world" {
			t.Errorf("FirstStr failed. Got: %s", result)
		}
	}

	{
		result := FirstStr("", "")
		if result != "" {
			t.Errorf("FirstStr failed. Got: %s", result)
		}
	}
}

func TestQueryEscape(t *testing.T) {
	result := QueryEscape("a key", "a value")
	if result != "a+key=a+value" {
		t.Errorf("QueryEscape failed. Got: %s", result)
	}
}

func TestTruncateString(t *testing.T) {
	tests := []struct {
		test   string
		len    int
		expect string
	}{
		{"abcd", 3, "abc"},
		{"abcd", 4, "abcd"},
		{"abcd", 5, "abcd"},
	}

	for _, test := range tests {
		t.Run(test.test, func(t *testing.T) {
			assert.Equal(t, test.expect, TruncateString(test.test, test.len))
		})
	}
}

//func TestGetLocale(t *testing.T) {
//	in := psp.BaseTransactionRequest{
//		Meta: psp.Meta{
//			BillingAddress: psp.Address{
//				Country: "US",
//			},
//		},
//	}
//	locale := GetLocale(in)
//	if locale != "en-US" {
//		t.Error("Expected en-US, got ", locale)
//	}
//	in.BillPayer.Language = "ru"
//	in.Meta.BillingAddress.Country = "RU"
//	locale = GetLocale(in)
//	if locale != "ru-RU" {
//		t.Error("Expected ru-RU, got ", locale)
//	}
//	in.Meta.Device.Language = "en-GB"
//	locale = GetLocale(in)
//	if locale != "en-GB" {
//		t.Error("Expected en-GB, got ", locale)
//	}
//	// only partial language, concat with billing country
//	in.Meta.Device.Language = "en"
//	locale = GetLocale(in)
//	if locale != "en-RU" {
//		t.Error("Expected en-RU, got ", locale)
//	}
//}

func TestGetCountryFromLocale(t *testing.T) {
	tests := []struct {
		locale string
		expect string
	}{
		{"en-US", "US"},
		{"en-GB", "GB"},
		{"en", ""},
	}

	for _, test := range tests {
		t.Run(test.locale, func(t *testing.T) {
			assert.Equal(t, test.expect, GetCountryFromLocale(test.locale))
		})
	}
}

func TestNormalizeString(t *testing.T) {
	assert.Equal(t, "AaEeIiOoUunNcssAEaeOEoeDdThth", NormalizeString("ÀáËëÏïÖöÜüñÑçßÆæŒœÐðÞþ"))
	assert.Equal(t, "Ni Hao", NormalizeString("你好"))
}

func TestNormalizeCardholderName(t *testing.T) {
	assert.Equal(t, "Bucky O'Hare", NormalizeCardholderName("Bucky O'Hare"))
	assert.Equal(t, "Cheng Long", NormalizeCardholderName("成龙"))
	assert.Equal(t, "Elena", NormalizeCardholderName("Елена"))
}

func TestAlphanumericOnly(t *testing.T) {
	assert.Equal(t, "BuckyOHare", AlphanumericOnly("Bucky O'Hare"))
	assert.Equal(t, "abc123", AlphanumericOnly("abc > 123"))
}

func TestNormalizeBrowserLanguage(t *testing.T) {
	tests := []struct {
		name   string
		tag    string
		expect string
	}{
		{"empty", "", ""},
		{"language only", "en", "en"},
		{"language and region", "en-GB", "en-GB"},
		{"region preserved", "pt-BR", "pt-BR"},
		{"at the limit", "ca-ES-va", "ca-ES-va"},
		{"variant dropped", "en-GB-oxendict", "en-GB"},
		{"valencian variant dropped", "ca-ES-valencia", "ca-ES"},
		{"script kept, region dropped", "zh-Hans-CN", "zh-Hans"},
		{"latin script kept", "sr-Latn-RS", "sr-Latn"},
		{"year variant dropped", "de-DE-1996", "de-DE"},
		{"long script drops to language", "abcdefg-Hant", "abcdefg"},
		{"oversized primary subtag truncated", "abcdefghij", "abcdefgh"},
		{"whitespace trimmed", "  en-GB-oxendict  ", "en-GB"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := NormalizeBrowserLanguage(test.tag)
			assert.Equal(t, test.expect, got)
			assert.LessOrEqual(t, len(got), maxBrowserLanguageLength)
		})
	}
}
