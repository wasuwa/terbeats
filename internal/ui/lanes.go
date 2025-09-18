package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Lanes レーン一覧
type Lanes struct {
	lanes []Lane
	frame *tview.Frame
}

// newLanes Lanesを初期化する
func newLanes(laneCount int) Lanes {
	lanes := Lanes{}
	lanes.setLanes(laneCount)
	lanes.setFrame()
	return lanes
}

// setLanes Lanesをセットする
func (ll *Lanes) setLanes(laneCount int) {
	lanes := make([]Lane, laneCount)
	for i := range laneCount {
		lanes[i] = newLane()
	}
	ll.lanes = lanes
}

// setFrame Frameをセットする
func (ll *Lanes) setFrame() {
	flex := tview.NewFlex().SetDirection(tview.FlexColumn)
	for _, lane := range ll.lanes {
		flex = flex.AddItem(lane.Frame, 0, 1, false)
	}
	// 外枠（左右の余白確保）
	space := 10
	ll.frame = tview.NewFrame(flex).SetBorders(0, 0, 0, 0, space, space)
}

// FallTo Lanesを落下させる
func (ll Lanes) FallTo(offset int) {
	for _, lane := range ll.lanes {
		// 上部にスペースを入れて擬似的に落下しているように見せている
		lane.Flex.Clear()
		lane.Flex.AddItem(tview.NewBox(), offset, 1, false)
		lane.Flex.AddItem(lane.Box, 1, 1, false)
	}
}

// Lane レーン
type Lane struct {
	Box   *tview.Box
	Flex  *tview.Flex
	Frame *tview.Frame
}

// newLane Laneを初期化する
func newLane() Lane {
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
