package synonymcli

import (
	"fmt"

	wordscli "github.com/tatoalonso/synant-cli/internal"

	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const (
	wordFlag   = "word"
	typeOfWord = "synonym"
)

// InitSynonymCmd initialize beers command
func InitSynonymCmd(repository wordscli.WordsRepo, repositoryStorage wordscli.WordsRepoStorage) *cobra.Command {
	sinonymCmd := &cobra.Command{
		Use:   "synonym",
		Short: "Provide synomnims of the input",
		Run:   runSynonymsFn(repository, repositoryStorage),
	}

	sinonymCmd.Flags().StringP(wordFlag, "s", "", "synonyms of the word")
	sinonymCmd.MarkFlagRequired(wordFlag)

	return sinonymCmd
}

func runSynonymsFn(repository wordscli.WordsRepo, repositoryStorage wordscli.WordsRepoStorage) CobraFn {
	return func(cmd *cobra.Command, args []string) {

		word, _ := cmd.Flags().GetString(wordFlag)

		synonyms, _ := repository.GetWords(typeOfWord, word)

		fmt.Println(synonyms.Sinonimos)

		repositoryStorage.SaveWords(typeOfWord, word, synonyms)

	}
}
