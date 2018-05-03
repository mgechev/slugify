package slugify

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// Transform returns a slugified string
func Transform(str string) string {
	str = strings.ToLower(str)
	result := ""

	var lengths []int
	for _, word := range multicharmap {
		currentLen := len(word)
		found := false
		for _, ex := range lengths {
			if ex == currentLen {
				found = true
			}
		}
		if !found {
			lengths = append(lengths, currentLen)
		}
	}
	for i := 0; i < utf8.RuneCountInString(str); i++ {
		char := string([]rune(str)[i])
		foundWord := false
		for _, currentLen := range lengths {
			if i+currentLen >= len(str) {
				continue
			}
			word := string([]rune(str)[i : i+currentLen])
			if replacement, ok := multicharmap[word]; ok {
				foundWord = true
				i += currentLen - 1
				char = replacement
			}
		}
		if foundWord {
			result += char
			continue
		}
		if replacement, ok := charmap[char]; ok {
			char = replacement
		}
		char = invalidChars.ReplaceAllString(char, "")
		result += char
	}
	result = notWhiteSpace.ReplaceAllString(result, "")
	result = whiteSpace.ReplaceAllString(result, "-")
	result = trailingReplacements.ReplaceAllString(result, "")
	return strings.ToLower(result)
}

var invalidChars = regexp.MustCompile(`[^\w\s\-\.\_~]`)
var notWhiteSpace = regexp.MustCompile(`^\s+|\s+$`)
var whiteSpace = regexp.MustCompile(`[-\s]+`)
var trailingReplacements = regexp.MustCompile(`-$`)
var replacement = "-"
var symbols = true
var lower = false
var multicharmap = map[string]string{
	"<3": "love", "&&": "and", "||": "or", "w/": "with",
}
var charmap = map[string]string{
	// latin
	"ß": "ss", "à": "a", "á": "a",
	"â": "a", "ã": "a", "ä": "a", "å": "a", "æ": "ae", "ç": "c", "è": "e",
	"é": "e", "ê": "e", "ë": "e", "ì": "i", "í": "i", "î": "i", "ï": "i",
	"ð": "d", "ñ": "n", "ò": "o", "ó": "o", "ô": "o", "õ": "o", "ö": "o",
	"ő": "o", "ø": "o", "ù": "u", "ú": "u", "û": "u", "ü": "u", "ű": "u",
	"ý": "y", "þ": "th", "ÿ": "y", "ẞ": "SS",
	// greek
	"α": "a", "β": "b", "γ": "g", "δ": "d", "ε": "e", "ζ": "z", "η": "h", "θ": "8",
	"ι": "i", "κ": "k", "λ": "l", "μ": "m", "ν": "n", "ξ": "3", "ο": "o", "π": "p",
	"ρ": "r", "σ": "s", "τ": "t", "υ": "y", "φ": "f", "χ": "x", "ψ": "ps", "ω": "w",
	"ά": "a", "έ": "e", "ί": "i", "ό": "o", "ύ": "y", "ή": "h", "ώ": "w", "ς": "s",
	"ϊ": "i", "ΰ": "y", "ϋ": "y", "ΐ": "i",
	// turkish
	"ş": "s", "Ş": "S", "ı": "i", "İ": "I",
	"ğ": "g",
	// russian
	"а": "a", "б": "b", "в": "v", "г": "g", "д": "d", "е": "e", "ё": "yo", "ж": "zh",
	"з": "z", "и": "i", "й": "j", "к": "k", "л": "l", "м": "m", "н": "n", "о": "o",
	"п": "p", "р": "r", "с": "s", "т": "t", "у": "u", "ф": "f", "х": "h", "ц": "c",
	"ч": "ch", "ш": "sh", "щ": "sh", "ъ": "u", "ы": "y", "ь": "", "э": "e", "ю": "yu",
	"я": "ya",
	// ukranian
	"є": "ye", "і": "i", "ї": "yi", "ґ": "g",
	// czech
	"č": "c", "ď": "d", "ě": "e", "ň": "n", "ř": "r", "š": "s", "ť": "t", "ů": "u",
	"ž": "z",
	// polish
	"ą": "a", "ć": "c", "ę": "e", "ł": "l", "ń": "n", "ś": "s", "ź": "z",
	"ż": "z",
	// latvian
	"ā": "a", "ē": "e", "ģ": "g", "ī": "i", "ķ": "k", "ļ": "l", "ņ": "n",
	"ū": "u",
	// lithuanian
	"ė": "e", "į": "i", "ų": "u", "Ė": "E", "Į": "I", "Ų": "U",
	// romanian
	"ț": "t", "Ț": "T", "ţ": "t", "Ţ": "T", "ș": "s", "Ș": "S", "ă": "a", "Ă": "A",
	// vietnamese
	"ạ": "a", "ả": "a", "ầ": "a", "ấ": "a", "ậ": "a", "ẩ": "a",
	"ẫ": "a", "ằ": "a", "ắ": "a", "ặ": "a", "ẳ": "a", "ẵ": "a", "ẹ": "e",
	"ẻ": "e", "ẽ": "e", "ề": "e", "ế": "e", "ệ": "e", "ể": "e", "ễ": "e",
	"ị": "i", "ỉ": "i", "ĩ": "i", "ọ": "o", "ỏ": "o", "ồ": "o", "ố": "o",
	"ộ": "o", "ổ": "o", "ỗ": "o", "ơ": "o", "ờ": "o", "ớ": "o", "ợ": "o",
	"ở": "o", "ỡ": "o", "ụ": "u", "ủ": "u", "ũ": "u", "ư": "u", "ừ": "u",
	"ứ": "u", "ự": "u", "ử": "u", "ữ": "u", "ỳ": "y", "ỵ": "y", "ỷ": "y",
	"ỹ": "y", "đ": "d",
	// currency
	"€": "euro", "₢": "cruzeiro", "₣": "french franc", "£": "pound",
	"₤": "lira", "₥": "mill", "₦": "naira", "₧": "peseta", "₨": "rupee",
	"₩": "won", "₪": "new shequel", "₫": "dong", "₭": "kip", "₮": "tugrik",
	"₯": "drachma", "₰": "penny", "₱": "peso", "₲": "guarani", "₳": "austral",
	"₴": "hryvnia", "₵": "cedi", "¢": "cent", "¥": "yen", "元": "yuan",
	"円": "yen", "﷼": "rial", "₠": "ecu", "¤": "currency", "฿": "baht",
	"$": "dollar", "₹": "indian rupee",
	// symbols
	"©": "(c)", "œ": "oe", "Œ": "OE", "∑": "sum", "®": "(r)", "†": "+",
	"“": "'", "”": "'", "‘": "'", "’": "'", "∂": "d", "ƒ": "f", "™": "tm",
	"℠": "sm", "…": "...", "˚": "o", "º": "o", "ª": "a", "•": "*",
	"∆": "delta", "∞": "infinity", "♥": "love", "&": "and", "|": "or",
	"<": "less", ">": "greater",
}
