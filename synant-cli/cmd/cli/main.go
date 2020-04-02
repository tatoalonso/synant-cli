package main

import (
	wordscli "github.com/tatoalonso/synant-cli/internal"
	"github.com/tatoalonso/synant-cli/internal/antonymcli"
	"github.com/tatoalonso/synant-cli/internal/storage/api"
	csvstorage "github.com/tatoalonso/synant-cli/internal/storage/csv"
	"github.com/tatoalonso/synant-cli/internal/synonymcli"

	"github.com/spf13/cobra"
)

func main() {

	var repo wordscli.WordsRepo
	var repoStorage wordscli.WordsRepoStorage

	repo = api.NewAPIRepository()
	repoStorage = csvstorage.NewCsvRepository()

	rootCmd := &cobra.Command{
		Use:   "synant-cli",
		Short: "synant is a tool which provides synonyms and antonyms",
		Long:  "synant is a cli tool which provides synonyms and antonyms in spanish,Complete documentation is available at :",
	}

	rootCmd.AddCommand(synonymcli.InitSynonymCmd(repo, repoStorage))
	rootCmd.AddCommand(antonymcli.InitAntonymCmd(repo, repoStorage))
	rootCmd.Execute()
}
