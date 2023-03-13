package api

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/wangluozhe/requests"
	"github.com/wangluozhe/requests/models"
	"io/ioutil"
)

const (
	listUrl      = "https://github.com/github/gitignore"
	gitignoreUrl = "https://raw.githubusercontent.com/github/gitignore/main/"
)

func ListAvailableLanguages() ([]string, error) {
	var (
		res       *models.Response
		doc       *goquery.Document
		languages []string
		err       error
	)
	if res, err = requests.Get(listUrl, nil); err != nil {
		return nil, err
	}

	defer res.Body.Close()
	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	doc.Find(`span [title$=".gitignore"][data-turbo-frame="repo-content-turbo-frame"]`).Each(func(i int, selection *goquery.Selection) {
		if title, ok := selection.Attr("title"); ok && len(title) > 10 {
			language := title[:len(title)-10]
			languages = append(languages, language)
		}
	})

	return languages, nil
}

func FetchGitignore(language string) (string, error) {
	var (
		res     *models.Response
		content []byte
		url     string
		err     error
	)
	url = fmt.Sprintf("%s/%s.gitignore", gitignoreUrl, language)

	if res, err = requests.Get(url, nil); err != nil {
		return "", err
	}

	defer res.Body.Close()
	content, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
