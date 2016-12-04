package ui

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Text struct {
	text string
	font *Font
}

// func (t *Text) NewText(text string) {
// 	t.length = len(t.text)
// }

func (t *Text) SetText(text string) {
	t.text = text
}

func (t *Text) GetText() string {
	return t.text
}

func (t *Text) SetFont(font *Font) {
	t.font = font
}

func (t *Text) GetFont() *Font {
	return t.font
}

func (t *Text) TextToVec2() []mgl32.Vec2 {
	text_slc := make([]mgl32.Vec2, len(t.text))

	for v := range text_slc {
		text_slc[v] = t.font.GetVec2(string(t.text[v]))
	}

	// text_arr[0] = mgl32.Vec2{935221.0, 731292.0}
	return text_slc
}
