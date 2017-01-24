package ui

import (
	"gde/render"
	// "github.com/go-gl/mathgl/mgl32"
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

func (t *Text) TextToVec4() []render.Vector4 {
	text_slc := make([]render.Vector4, len(t.text))

	for v := range text_slc {
		text_slc[v] = t.font.GetVec4(string(t.text[v]))
	}

	// text_arr[0] = mgl32.Vec2{935221.0, 731292.0}
	return text_slc
}
