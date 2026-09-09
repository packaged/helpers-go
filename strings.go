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
	for _, s := range input {
		if s != "" {
			return s
		}
	}
	return ""
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

// maxBrowserLanguageLength is the EMV 3DS 2.2 limit on the browserLanguage data
// element. The field widens to 35 characters in 3DS 2.3, but 8 is within both.
const maxBrowserLanguageLength = 8

// browserLanguageRegion matches a BCP 47 region subtag: two letters, or the
// three-digit UN M.49 form ("419" for Latin America). Script subtags are always
// four letters and variants five to eight, so neither can be mistaken for one.
var browserLanguageRegion = regexp.MustCompile(`^([A-Za-z]{2}|[0-9]{3})$`)

// NormalizeBrowserLanguage fits a browser-supplied BCP 47 language tag into the
// EMV 3DS browserLanguage element. Tags carrying a script or variant subtag
// ("en-GB-oxendict", "zh-Hans-CN", "ca-ES-valencia") are valid BCP 47 but exceed
// the limit, and a 3DS server rejects the whole AReq rather than ignoring the
// field.
//
// Reduce to language-region, dropping any script and variant in between, rather
// than simply cutting subtags off the end. The two differ only for tags carrying
// both a script and a region, where the end-cutting approach keeps the script
// ("zh-Hans") and this keeps the region ("zh-CN"). language-region is the shape
// browsers overwhelmingly report, so it is the shape an ACS is most likely to
// recognise when picking a challenge language, and for Chinese the region
// implies the script anyway. It also leaves the ACS a region to correlate
// against the billing country.
//
// Truncating to length instead would emit "en-GB-ox", trading a length error for
// a format one.
func NormalizeBrowserLanguage(tag string) string {
	tag = strings.TrimSpace(tag)
	if len(tag) <= maxBrowserLanguageLength {
		return tag
	}

	subtags := strings.Split(tag, "-")
	language := subtags[0]

	for _, subtag := range subtags[1:] {
		// A single-character subtag is an extension singleton; no region follows.
		if len(subtag) == 1 {
			break
		}
		if browserLanguageRegion.MatchString(subtag) {
			if candidate := language + "-" + subtag; len(candidate) <= maxBrowserLanguageLength {
				return candidate
			}
			break
		}
	}

	// An oversized primary subtag is not a valid language tag whatever we do, so
	// keep the field within length and let the server judge the value.
	return TruncateStringToRune(language, maxBrowserLanguageLength)
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
