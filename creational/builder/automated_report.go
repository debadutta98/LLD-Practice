package builder

import (
	"errors"
	"fmt"
)

type Theme string

const (
	LIGHT     Theme = "light"
	DARK      Theme = "dark"
	CORPORATE Theme = "corporate"
)

type Report struct {
	title               string
	headerText          string
	footerText          string
	themeText           Theme
	sections            []string
	includePageNumber   bool
	includeTableContent bool
}

func (r *Report) Title() string         { return r.title }
func (r *Report) HeaderText() string    { return r.headerText }
func (r *Report) FooterText() string    { return r.footerText }
func (r *Report) Theme() Theme          { return r.themeText }
func (r *Report) PageNumber() bool      { return r.includePageNumber }
func (r *Report) TableOfContents() bool { return r.includeTableContent }

func (r *Report) Sections() []string {
	s := make([]string, len(r.sections))
	copy(s, r.sections)
	return s
}

func (r *Report) Print() {
	fmt.Printf("Report: %+v\n", r)
}

type ReportBuilder interface {
	SetTitle(string) ReportBuilder
	SetHeaderText(string) ReportBuilder
	SetFooterText(string) ReportBuilder
	SetTheme(Theme) ReportBuilder
	AddSection(string) ReportBuilder
	SetIncludePageNumber(bool) ReportBuilder
	SetIncludeTableContents(bool) ReportBuilder
	GetReport() (*Report, error)
}

type ReportLibraryBuilder struct {
	title               string
	headerText          string
	footerText          string
	themeText           Theme
	sections            []string
	includePageNumber   bool
	includeTableContent bool
}

func NewReportBuilder() ReportBuilder {
	return &ReportLibraryBuilder{
		themeText: LIGHT,
	}
}

func (r *ReportLibraryBuilder) SetTitle(val string) ReportBuilder {
	r.title = val
	return r
}

func (r *ReportLibraryBuilder) SetHeaderText(val string) ReportBuilder {
	r.headerText = val
	return r
}

func (r *ReportLibraryBuilder) SetFooterText(val string) ReportBuilder {
	r.footerText = val
	return r
}

func (r *ReportLibraryBuilder) SetTheme(val Theme) ReportBuilder {
	r.themeText = val
	return r
}

func (r *ReportLibraryBuilder) AddSection(val string) ReportBuilder {
	r.sections = append(r.sections, val)
	return r
}

func (r *ReportLibraryBuilder) SetIncludePageNumber(val bool) ReportBuilder {
	r.includePageNumber = val
	return r
}

func (r *ReportLibraryBuilder) SetIncludeTableContents(val bool) ReportBuilder {
	r.includeTableContent = val
	return r
}

func (r *ReportLibraryBuilder) GetReport() (*Report, error) {
	if r.title == "" {
		return nil, errors.New("report title is required")
	}

	sectionsCopy := make([]string, len(r.sections))
	copy(sectionsCopy, r.sections)

	return &Report{
		title:               r.title,
		headerText:          r.headerText,
		footerText:          r.footerText,
		themeText:           r.themeText,
		sections:            sectionsCopy,
		includePageNumber:   r.includePageNumber,
		includeTableContent: r.includeTableContent,
	}, nil
}

type Reporter struct {
	builder ReportBuilder
}

func NewReporter(builder ReportBuilder) *Reporter {
	return &Reporter{builder: builder}
}

func (d *Reporter) BuildExecutiveSummaryReport(title, header, footer string, sections []string) (*Report, error) {
	b := d.builder.
		SetTitle(title).
		SetHeaderText(header).
		SetFooterText(footer).
		SetTheme(CORPORATE).
		SetIncludePageNumber(true).
		SetIncludeTableContents(true)

	for _, sec := range sections {
		b.AddSection(sec)
	}

	return b.GetReport()
}

func (d *Reporter) BuildQuickDraftReport(title string, sections []string) (*Report, error) {
	b := d.builder.
		SetTitle(title).
		SetTheme(LIGHT).
		SetIncludePageNumber(false).
		SetIncludeTableContents(false)

	for _, sec := range sections {
		b.AddSection(sec)
	}

	return b.GetReport()
}
