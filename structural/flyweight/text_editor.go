package flyweight

import (
	"fmt"
)

type CharacterStyle struct {
	fontFamily string
	fontSize   uint32
	color      string
}

func NewCharacterStyle(fontFamily string, fontSize uint32, color string) *CharacterStyle {
	return &CharacterStyle{
		fontFamily: fontFamily,
		fontSize:   fontSize,
		color:      color,
	}
}

func (s *CharacterStyle) GetFontFamily() string { return s.fontFamily }
func (s *CharacterStyle) GetFontSize() uint32   { return s.fontSize }
func (s *CharacterStyle) GetColor() string      { return s.color }

type StyleFactory struct {
	styles map[string]*CharacterStyle
}

func NewStyleFactory() *StyleFactory {
	return &StyleFactory{
		styles: make(map[string]*CharacterStyle), // Properly initialized map
	}
}

func (f *StyleFactory) GetStyle(fontFamily string, fontSize uint32, color string) *CharacterStyle {
	key := fmt.Sprintf("%s-%d-%s", fontFamily, fontSize, color)

	if style, exists := f.styles[key]; exists {
		return style // Return existing shared flyweight
	}

	style := NewCharacterStyle(fontFamily, fontSize, color)
	f.styles[key] = style
	return style
}

type Character struct {
	ch    rune
	style *CharacterStyle
}

func NewCharacter(ch rune, style *CharacterStyle) *Character {
	return &Character{
		ch:    ch,
		style: style,
	}
}

func (c *Character) Render() {
	fmt.Printf("Render Char '%c' [Color: %s, Font: %s, Size: %d]\n",
		c.ch, c.style.GetColor(), c.style.GetFontFamily(), c.style.GetFontSize())
}

type TextEditor struct {
	factory *StyleFactory
	chars   []*Character
}

func NewTextEditor() *TextEditor {
	return &TextEditor{
		factory: NewStyleFactory(),
		chars:   make([]*Character, 0),
	}
}

func (e *TextEditor) InsertCharacter(ch rune, fontFamily string, fontSize uint32, color string) {
	style := e.factory.GetStyle(fontFamily, fontSize, color)
	character := NewCharacter(ch, style)
	e.chars = append(e.chars, character)
}

func (e *TextEditor) RenderDocument() {
	for _, c := range e.chars {
		c.Render()
	}
}
