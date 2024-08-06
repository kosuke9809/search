package main

import (
	"fmt"
	"search/analyzer"
)

func main() {
	text := "ＡＢＣ１２３　ｱｲｳｴｵ　カキクケコ　ﾊﾟﾋﾟﾌﾟﾍﾟﾎﾟ　、。！？　マツモトキヨシ"
	mappings := map[string]string{
		"マツモトキヨシ": "ドラッグストア",
		"１":       "一",
		"２":       "二",
		"３":       "三",
	}
	charFilter := analyzer.NewCompositeCharFilter(mappings)
	result := charFilter.Filter(text)
	fmt.Println(result)
}
