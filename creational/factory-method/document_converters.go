package factorymethod

import "fmt"

type Format string

const (
	HTML     Format = "html"
	PDF      Format = "pdf"
	WORD     Format = "word"
	Markdown Format = "markdown"
)

type Converter interface {
	Convert(string) string
}

type markdownToPDF struct{}

func (m markdownToPDF) Convert(content string) string {
	return fmt.Sprintf("Converting Markdown to PDF: %s", content)
}

type markdownToWord struct{}

func (m markdownToWord) Convert(content string) string {
	return fmt.Sprintf("Converting Markdown to Word: %s", content)
}

type htmlToPDF struct{}

func (m htmlToPDF) Convert(content string) string {
	return fmt.Sprintf("Converting HTML to PDF: %s", content)
}

type htmlToWORD struct{}

func (m htmlToWORD) Convert(content string) string {
	return fmt.Sprintf("Converting HTML to Word: %s", content)
}

func getKey(src, tar Format) string {
	return fmt.Sprintf("%s:%s", src, tar)
}

var convertersMap = map[string]Converter{
	getKey(HTML, WORD):     htmlToWORD{},
	getKey(HTML, PDF):      htmlToPDF{},
	getKey(Markdown, PDF):  markdownToPDF{},
	getKey(Markdown, WORD): markdownToWord{},
}

func GetConverter(srcFormat, tarFormat Format) (Converter, error) {
	unsupportedErr := fmt.Errorf("%s to %s is unsupported", srcFormat, tarFormat)
	if val, ok := convertersMap[getKey(srcFormat, tarFormat)]; ok {
		return val, nil
	} else {
		return nil, unsupportedErr
	}
}
