package helpers

import (
	"io"
	"math/rand"
	"net/url"
	"regexp"
	"strings"

	"github.com/alexsergivan/transliterator"
)

// TruncateStringToRune takes a string, limits it by byte length, preserving UTF-8 characters and direction (e.g. safe for Hebrew)
func TruncateStringToRune(str string, byteLength int) string {
	a := strings.ToValidUTF8(str, "")
	if len(a) > byteLength {
		a = strings.ToValidUTF8(a[0:byteLength], "")
	}
	return a
}

func BufferToString(buf io.Reader) string {
	resp, _ := io.ReadAll(buf)
	return string(resp)
}

func FirstStr(input ...string) string {
	return Coalesce(input...)
}

func TruncateString(input string, length int) string {
	if len(input) > length {
		return input[:length]
	}
	return input
}

func QueryEscape(k, v string) string {
	return url.QueryEscape(k) + "=" + url.QueryEscape(v)
}

func FilterEmptyStrings(input []string) []string {
	var out []string
	for _, s := range input {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func GetLanguageFromLocale(locale string) string {
	if len(locale) < 2 {
		return ""
	}
	return locale[:2]
}

var localeRegex = regexp.MustCompile("[a-z]{2}-[A-Z]{2}")

//func GetLocale(request psp.BaseTransactionRequest) string {
//	locale := FirstStr(request.Meta.Device.Language, string(request.BillPayer.Language))
//	if locale == "" {
//		locale = "en-US"
//	}
//	if localeRegex.MatchString(locale) {
//		return locale
//	}
//	return fmt.Sprintf("%s-%s", locale, strings.ToUpper(string(request.Meta.BillingAddress.Country)))
//}

func GetCountryFromLocale(locale string) string {
	if len(locale) < 5 {
		return ""
	}
	return locale[3:]
}

var translit = transliterator.NewTransliterator(nil)

func NormalizeString(s string) string {
	return strings.TrimSpace(translit.Transliterate(s, ""))
}

var cardholderRegex = regexp.MustCompile("[^A-Za-z0-0.' -]+")

func NormalizeCardholderName(s string) string {
	s = NormalizeString(s)
	return cardholderRegex.ReplaceAllString(s, "")
}

var alphanumeric = regexp.MustCompile(`[^a-zA-Z0-9]`)

// AlphanumericOnly removes non-alphanumeric chars from string
func AlphanumericOnly(input string) (output string) {
	return alphanumeric.ReplaceAllString(input, "")
}

func Pattern(pattern string) string {
	var result strings.Builder
	for _, ch := range pattern {
		char := ch
		switch char {
		case '!':
			match := []rune{'X', '0'}
			char = match[rand.Intn(2)]
		case '?':
			match := []rune{'x', '0'}
			char = match[rand.Intn(2)]
		case '*':
			match := []rune{'x', '0', 'X'}
			char = match[rand.Intn(3)]
		}

		switch char {
		case 'X':
			result.WriteByte(byte(rand.Intn(26) + 65)) // A-Z
		case 'x':
			result.WriteByte(byte(rand.Intn(26) + 97)) // a-z
		case '0':
			result.WriteByte(byte(rand.Intn(10) + 48)) // 0-9
		default:
			if char >= '1' && char <= '9' {
				maxDigit := int(char - '0')
				result.WriteByte(byte(rand.Intn(maxDigit+1) + 48))
			} else {
				result.WriteRune(char)
			}
		}
	}
	return result.String()
}

func VerifyPattern(template, pattern string) bool {
	replacer := strings.NewReplacer(
		"X", "[A-Z]",
		"x", "[a-z]",
		"0", "[0-9]",
		"5", "[0-5]",
		"?", "[a-z0-9]",
		"!", "[A-Z0-9]",
		"*", "[a-zA-Z0-9]",
	)
	regexPattern := replacer.Replace(template)
	re := regexp.MustCompile("^" + regexPattern + "$")
	return re.MatchString(pattern)
}
