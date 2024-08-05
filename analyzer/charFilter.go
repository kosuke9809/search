package analyzer

import (
	"strings"
	"unicode"
)

type CharFilter interface {
	Filter(text string) string
}

// 大文字 -> 小文字変換
type LowercaseFilter struct{}

func (f LowercaseFilter) Filter(text string) string {
	return strings.ToLower(text)
}

// 空白の正規化 半角スペース化
type WhitespaceNormalizationFilter struct{}

func (f WhitespaceNormalizationFilter) Filter(text string) string {
	return strings.Join(strings.Fields(text), " ")
}

// 句読点の排除
type PunctuationRemovalFilter struct{}

func (f PunctuationRemovalFilter) Filter(text string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, text)
}

// 文字列置換
type MappingCharFilter struct {
	Mappings map[string]string
}

func (f MappingCharFilter) Filter(text string) string {
	for old, new := range f.Mappings {
		text = strings.ReplaceAll(text, old, new)
	}
	return text
}

// 全角 -> 半角変換
type FullWidthToHalfWidthFilter struct{}

func (f FullWidthToHalfWidthFilter) Filter(text string) string {
	return ""
}
