package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Lane 1レーン分のコンポーネント
type Lane struct {
	Box   *tview.Box
	Flex  *tview.Flex
	Frame *tview.Frame
}

// NewLane レーンを初期化する
func NewLane() Lane {
	b := newBox()
	f := newFlex(b)
	fr := newFrame(f)
	return Lane{Box: b, Flex: f, Frame: fr}
}

// newBox Boxを初期化する
func newBox() *tview.Box {
	return tview.NewBox().SetBackgroundColor(tcell.ColorWhite)
}

// newFlex Flexを初期化する
func newFlex(box *tview.Box) *tview.Flex {
	note := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(box, 1, 1, false)
	return note
}

// newFrame Frameを初期化する
func newFrame(note *tview.Flex) *tview.Frame {
	return tview.NewFrame(note).SetBorders(0, 0, 0, 0, 1, 1)
}

// FallTo 落下する
func (l *Lane) FallTo(offset int) {
	// 上部にスペースを入れて擬似的に落下しているように見せている
	l.Flex.Clear()
	l.Flex.AddItem(tview.NewBox(), offset, 1, false)
	l.Flex.AddItem(l.Box, 1, 1, false)
}
