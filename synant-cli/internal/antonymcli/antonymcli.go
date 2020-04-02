package antonymcli

import (
	"fmt"

	wordscli "github.com/tatoalonso/synant-cli/internal"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const (
	wordFlag   = "word"
	typeOfWord = "antonym"
)

// InitAntonymCmd initialize beers command
func InitAntonymCmd(repository wordscli.WordsRepo, repositoryStorage wordscli.WordsRepoStorage) *cobra.Command {
	antonymCmd := &cobra.Command{
		Use:   "antonym",
		Short: "Provide antonyms of the input",
		Run:   runAntonymsFn(repository, repositoryStorage),
	}

	antonymCmd.Flags().StringP(wordFlag, "a", "", "antonyms of the word")
	antonymCmd.MarkFlagRequired(wordFlag)

	return antonymCmd
}

func runAntonymsFn(repository wordscli.WordsRepo, repositoryStorage wordscli.WordsRepoStorage) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		word, _ := cmd.Flags().GetString(wordFlag)

		antonyms, _ := repository.GetWords(typeOfWord, word)

		fmt.Println(antonyms.Antonimos)

		repositoryStorage.SaveWords(typeOfWord, word, antonyms)
	}
}
