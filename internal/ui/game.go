package ui

import (
	"time"

	"github.com/rivo/tview"
)

// NewGameUI ゲームUIを初期化する
func NewGameUI(app *tview.Application) tview.Primitive {
	// レーン生成
	lane1 := NewLane()
	lane2 := NewLane()
	lane3 := NewLane()
	lane4 := NewLane()

	// レーン群のレイアウト（横並び）
	notes := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(lane1.Frame, 0, 1, false).
		AddItem(lane2.Frame, 0, 1, false).
		AddItem(lane3.Frame, 0, 1, false).
		AddItem(lane4.Frame, 0, 1, false)

	// 外枠（左右の余白確保）
	wrapper := tview.NewFrame(notes).SetBorders(1, 1, 0, 0, 2, 2)

	// アニメーション（落下）
	position := 0
	go func() {
		for {
			position += 1

			app.QueueUpdateDraw(func() {
				lane1.FallTo(position)
				lane2.FallTo(position)
				lane3.FallTo(position)
				lane4.FallTo(position)
			})

			if isPositionBottom(wrapper, position) {
				break
			}

			time.Sleep(10 * time.Millisecond)
		}
	}()

	return wrapper
}

// isPositionBottom 下部か判定する
func isPositionBottom(root *tview.Frame, position int) bool {
	_, _, _, screenHeight := root.GetRect()
	return position > screenHeight-3
}
