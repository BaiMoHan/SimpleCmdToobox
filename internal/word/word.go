package word

import (
	"strings"
	"unicode"
)

// 单词全部转换为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// 单词全部转换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// 下画线单词转大写驼峰单词
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// 下划线单词转小写驼峰单词
func UnderscoreToLowerCamelCase(s string) string {
	// 复用转大写驼峰
	s = UnderscoreToUpperCamelCase(s)
	// 字符串第一位取出来 其他的拼接
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// 驼峰单词转下划线单词
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		// 从第二个单词开始加下划线 并转小写
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
