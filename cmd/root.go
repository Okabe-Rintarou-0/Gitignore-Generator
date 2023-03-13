package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"igen/api"
	"igen/constants"
	"os"
	"path"
	"strings"
)

var (
	outPath string
	target  string
)

func handleError(err error) {
	fmt.Println(err.Error())
	os.Exit(-1)
}

var rootCmd = &cobra.Command{
	Use:   "igen",
	Short: "Generate .gitignore from github/gitignore",
	Long:  "Generate .gitignore from github/gitignore, for more info, please check https://github.com/github/gitignore",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			ok       bool
			likely   []string
			file     *os.File
			content  string
			language string
			err      error
		)
		if ok, language, likely = checkLanguage(target); ok {
			if content, err = api.FetchGitignore(language); err != nil {
				handleError(err)
			}
			if err = fixOutPath(); err != nil {
				handleError(err)
			}
			if file, err = os.Create(outPath); err != nil {
				handleError(err)
			}
			defer file.Close()
			if _, err = file.WriteString(content); err != nil {
				handleError(err)
			}
			fmt.Println("Done.")
		} else {
			fmt.Printf("%s not found in language set, did you mean:\n", target)
			for _, l := range likely {
				fmt.Println(l)
			}
		}
	},
}

func checkLanguage(target string) (bool, string, []string) {
	var (
		likely            []string
		targetLowercase   string
		languageLowercase string
	)

	for _, language := range constants.LanguageSet {
		if language == target {
			return true, language, nil
		}

		targetLowercase = strings.ToLower(target)
		languageLowercase = strings.ToLower(language)

		if targetLowercase == languageLowercase {
			return true, language, nil
		}

		if strings.Contains(languageLowercase, targetLowercase) || strings.Contains(targetLowercase, languageLowercase) {
			likely = append(likely, language)
		}
	}

	return false, "", likely
}

func existsFile(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func isFile(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

func fixOutPath() error {
	var (
		filename string
		dir      string
		err      error
	)
	filename = path.Base(outPath)
	if filename == ".gitignore" {
		dir = path.Dir(outPath)
	}
	dir = outPath
	if !existsFile(dir) {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	} else if isFile(dir) {
		return fmt.Errorf("%s shouldn't be a file", dir)
	}

	outPath = path.Join(dir, ".gitignore")
	return nil
}

func init() {
	rootCmd.Flags().StringVarP(&target, "target", "t", "", "target language")
	rootCmd.Flags().StringVarP(&outPath, "out", "o", "./", "output dir")
	_ = rootCmd.MarkFlagRequired("target")
	rootCmd.AddCommand(listCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}
