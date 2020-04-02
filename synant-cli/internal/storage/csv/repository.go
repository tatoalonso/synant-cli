package csvstorage

import (
	"encoding/csv"
	"fmt"
	"os"

	wordscli "github.com/tatoalonso/synant-cli/internal"
)

const (
	pathFile   = "../../data/"
	fileType   = ".csv"
	antonymPre = "antonym-"
	synonymPre = "synonym-"
)

type wordRepoStorage struct {
	path string
}

// NewCsvRepository save synonyms or antonyms to csv file
func NewCsvRepository() wordscli.WordsRepoStorage {

	return &wordRepoStorage{path: pathFile}
}

func (w *wordRepoStorage) SaveWords(typeOfWord string, word string, list wordscli.ListOfWords) (err error) {

	var pre string
	var arrayWords []string
	var arrayOfArrayWords [][]string

	switch typeOfWord {
	case "synonym":
		pre = synonymPre
		listSyn := list.Sinonimos

		for _, v := range listSyn {
			arrayWords = append(arrayWords, v.Sinonimo)
			arrayOfArrayWords = append(arrayOfArrayWords, arrayWords)
			arrayWords = nil
		}

	case "antonym":
		pre = antonymPre
		listAnt := list.Antonimos

		for _, v := range listAnt {
			arrayWords = append(arrayWords, v.Antonimo)
			arrayOfArrayWords = append(arrayOfArrayWords, arrayWords)
			arrayWords = nil
		}
	}

	file, err := os.Create(fmt.Sprintf("%v%v%v%v", w.path, pre, word, fileType))

	if err != nil {
		fmt.Println(err)
		return
	}

	writer := csv.NewWriter(file)

	err = writer.WriteAll(arrayOfArrayWords) // returns error
	if err != nil {
		fmt.Println("An error encountered ::", err)

		return

	}

	return
}
