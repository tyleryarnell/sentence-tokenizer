package main

import (
	"io/ioutil"
	"strings"
	"testing"

	td "github.com/neurosnap/sentences/data"
	"github.com/neurosnap/sentences/punkt"
)

func loadTokenizer(data string) *punkt.DefaultSentenceTokenizer {
	b, err := td.Asset(data)
	if err != nil {
		panic(err)
	}

	training, err := punkt.LoadTraining(b)
	if err != nil {
		panic(err)
	}

	return punkt.NewTokenizer(training)
}

func readFile(fname string) string {
	content, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func getFileLocation(prefix, original, expected string) []string {
	orig_text := strings.Join([]string{prefix, original}, "")
	expected_text := strings.Join([]string{prefix, expected}, "")
	return []string{orig_text, expected_text}
}

func TestEnglish(t *testing.T) {
	t.Log("Starting test suite ...")

	tokenizer := loadTokenizer("data/english.json")

	prefix := "test_files/english/"

	test_files := [][]string{
		getFileLocation(prefix, "carolyn.txt", "carolyn_s.txt"),
		getFileLocation(prefix, "ecig.txt", "ecig_s.txt"),
		getFileLocation(prefix, "foul_ball.txt", "foul_ball_s.txt"),
		getFileLocation(prefix, "fbi.txt", "fbi_s.txt"),
		getFileLocation(prefix, "dre.txt", "dre_s.txt"),
		getFileLocation(prefix, "dr.txt", "dr_s.txt"),
		getFileLocation(prefix, "quotes.txt", "quotes_s.txt"),
		getFileLocation(prefix, "kiss.txt", "kiss_s.txt"),
		getFileLocation(prefix, "kentucky.txt", "kentucky_s.txt"),
		getFileLocation(prefix, "iphone6s.txt", "iphone6s_s.txt"),
		getFileLocation(prefix, "lebanon.txt", "lebanon_s.txt"),
		getFileLocation(prefix, "duma.txt", "duma_s.txt"),
		getFileLocation(prefix, "demolitions.txt", "demolitions_s.txt"),
		getFileLocation(prefix, "qa.txt", "qa_s.txt"),
	}

	for _, f := range test_files {
		actual_text := readFile(f[0])
		expected_text := readFile(f[1])
		expected := strings.Split(expected_text, "{{sentence_break}}")

		t.Log(f[0])
		sentences := tokenizer.NTokenize(actual_text)
		for index, s := range sentences {
			if strings.TrimSpace(s) != strings.TrimSpace(expected[index]) {
				t.Logf("Actual  : %q", s)
				t.Log("--------")
				t.Logf("Expected: %q", strings.TrimSpace(expected[index]))
				t.Fatalf("%s line %d: Actual sentence does not match expected sentence", f[0], index)
			}
		}
	}

}
