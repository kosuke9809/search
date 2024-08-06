package analyzer

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
	"golang.org/x/text/width"
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
	return width.Narrow.String(text)
}

type FullWidthSpaceToHalfWidthFilter struct{}

func (f FullWidthSpaceToHalfWidthFilter) Filter(text string) string {
	return strings.ReplaceAll(text, "　", " ")
}

type KatakanaNormalizationFilter struct{}

func (f KatakanaNormalizationFilter) Filter(text string) string {
	return norm.NFKC.String(text)
}

type DakutenNormalizationFilter struct{}

func (f DakutenNormalizationFilter) Filter(text string) string {
	return normalizeDakuten(text)
}

func normalizeDakuten(text string) string {
	n := []rune{}
	for i, r := range text {
		if i > 0 && (r == '゛' || r == '゜') {
			prev := n[len(n)-1]
			combined := norm.NFKC.String(string([]rune{prev, r}))
			n[len(n)-1] = []rune(combined)[0]
		} else {
			n = append(n, r)
		}
	}
	return string(n)
}

type CompositeCharFilter struct {
	Filters []CharFilter
}

func (f CompositeCharFilter) Filter(text string) string {
	for _, filter := range f.Filters {
		text = filter.Filter(text)
	}
	return text
}

func NewCompositeCharFilter(mappings map[string]string) CompositeCharFilter {
	return CompositeCharFilter{
		Filters: []CharFilter{
			MappingCharFilter{Mappings: mappings},
			FullWidthToHalfWidthFilter{},
			FullWidthSpaceToHalfWidthFilter{},
			WhitespaceNormalizationFilter{},
			LowercaseFilter{},
			PunctuationRemovalFilter{},
			KatakanaNormalizationFilter{},
			DakutenNormalizationFilter{},
		},
	}
}
