package prototype

import (
	"maps"
)

type Cloneable interface {
	Clone() Cloneable
}

type Section struct {
	Section_title string
	Content_body  string
}

type Document struct {
	title       string
	header_text string
	footer_text string
	sections    []Section
	metadata    map[string]any
}

func (doc *Document) GetTitle() string {
	return doc.title
}

func (doc *Document) SetTitle(title string) {
	doc.title = title
}

func (doc *Document) GetHeaderText() string {
	return doc.header_text
}

func (doc *Document) SetHeaderText(text string) {
	doc.header_text = text
}

func (doc *Document) GetFooterText() string {
	return doc.footer_text
}

func (doc *Document) SetFooterText(text string) {
	doc.footer_text = text
}

func (doc *Document) GetSections() []Section {
	sections := make([]Section, len(doc.sections))
	copy(sections, doc.sections)
	return sections
}

func (doc *Document) SetSections(sections []Section) {
	doc.sections = make([]Section, len(sections))
	copy(doc.sections, sections)
}

func (doc *Document) GetMetadata() map[string]any {
	metadata := make(map[string]any)
	maps.Copy(metadata, doc.metadata)
	return metadata
}

func (doc *Document) SetMetadata(metadata map[string]any) {
	doc.metadata = make(map[string]any)
	maps.Copy(doc.metadata, metadata)
}

func NewDocument(
	title, header_text, footer_text string,
	sections []Section,
	metadata map[string]any) *Document {
	doc := &Document{
		title:       title,
		header_text: header_text,
		footer_text: footer_text,
		sections:    make([]Section, len(sections)),
		metadata:    make(map[string]any),
	}
	copy(doc.sections, sections)
	maps.Copy(doc.metadata, metadata)
	return doc
}

func (doc *Document) Clone() Cloneable {
	return NewDocument(
		doc.title,
		doc.header_text,
		doc.footer_text,
		doc.sections,
		doc.metadata,
	)
}
