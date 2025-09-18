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
		lane.fallTo(offset)
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
	l := Lane{}
	l.setBox()
	l.setFlex()
	l.setFrame()
	return l
}

// setBox Boxをセットする
func (l *Lane) setBox() {
	l.Box = tview.NewBox().SetBackgroundColor(tcell.ColorWhite)
}

// setFlex Flexをセットする
func (l *Lane) setFlex() {
	l.Flex = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(l.Box, 1, 1, false)
}

// setFrame Frameをセットする
func (l *Lane) setFrame() {
	l.Frame = tview.NewFrame(l.Flex).SetBorders(0, 0, 0, 0, 1, 1)
}

// fallTo Laneを落下させる
func (l Lane) fallTo(offset int) {
	// 上部にスペースを入れて擬似的に落下しているように見せている
	l.Flex.Clear()
	l.Flex.AddItem(tview.NewBox(), offset, 1, false)
	l.Flex.AddItem(l.Box, 1, 1, false)
}
