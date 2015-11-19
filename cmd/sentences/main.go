package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/neurosnap/sentences/english"
	"github.com/spf13/cobra"
)

var VERSION string
var COMMITHASH string

var ver bool
var fname string
var delim string

var sentencesCmd = &cobra.Command{
	Use:   "sentences",
	Short: "Sentence tokenizer",
	Long:  "A utility that will break up a blob of text into sentences.",
	Run: func(cmd *cobra.Command, args []string) {
		var text []byte

		if ver {
			fmt.Println(VERSION)
			fmt.Println(COMMITHASH)
			return
		}

		if fname != "" {
			text, _ = ioutil.ReadFile(fname)
		} else {
			reader := bufio.NewReader(os.Stdin)
			text, _ = ioutil.ReadAll(reader)
		}

		tokenizer, err := english.NewSentenceTokenizer(nil)
		if err != nil {
			panic(err)
		}

		sentences := tokenizer.Tokenize(string(text))
		for _, s := range sentences {
			text := strings.Join(strings.Fields(s.Text), " ")

			text = strings.Join([]string{text, delim}, "")
			fmt.Printf("%s", text)
		}
	},
}

func main() {
	sentencesCmd.Flags().BoolVarP(&ver, "version", "v", false, "Get current version of sentences")
	sentencesCmd.Flags().StringVarP(&fname, "file", "f", "", "Read file as source input instead of stdin")
	sentencesCmd.Flags().StringVarP(&delim, "delimiter", "d", "\n", "Delimiter used to demarcate sentence boundaries")

	if err := sentencesCmd.Execute(); err != nil {
		fmt.Print(err)
	}
}