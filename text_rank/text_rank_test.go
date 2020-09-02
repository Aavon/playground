package text_rank

import (
	"fmt"
	"github.com/DavidBelicza/TextRank"
	"github.com/DavidBelicza/TextRank/convert"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_a(t *testing.T) {
	rawText := "Go语言中文网,中国 Golang 社区,Go语言学习园地,致力于构建完善的 Golang 中文社区,Go语言爱好者的学习家园。分享 Go 语言知识,交流使用经验"
	// TextRank object
	tr := textrank.NewTextRank()
	// Default Rule for parsing.
	rule := textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.NewDefaultLanguage()
	convert.NewLanguage()
	// Default algorithm for ranking text.
	algorithmDef := textrank.NewDefaultAlgorithm()

	// zh
	zh_words := []string{}
	err := filepath.Walk("./raw", func(path string, info os.FileInfo, err error) error {
		fmt.Println(info.Name())
		if info.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		for _, w := range strings.Split(string(data), "\n") {
			zh_words = append(zh_words, strings.TrimSpace(w))
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("words: ", len(zh_words))
	language.SetWords("zh", zh_words)
	language.SetActiveLanguage("zh")

	// Add text.
	tr.Populate(rawText, language, rule)
	// Run the ranking.
	tr.Ranking(algorithmDef)

	// Get all phrases by weight.
	rankedPhrases := textrank.FindPhrases(tr)

	signleWords := textrank.FindSingleWords(tr)
	fmt.Println(len(signleWords))
	words := []string{}
	for _, p := range signleWords {
		words = append(words, p.Word)
	}

	fmt.Println(strings.Join(words, " "))

	// Most important phrase.
	fmt.Println(rankedPhrases[0])
	// Second important phrase.
	fmt.Println(rankedPhrases[1])
}
