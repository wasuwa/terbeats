package ui

import (
	"time"

	"github.com/rivo/tview"
)

const (
	// laneCount 画面に表示するレーンの数
	laneCount = 4
)

// NewGame ゲームを初期化する
func NewGame(app *tview.Application) tview.Primitive {
	lanes := newLanes(laneCount)

	// アニメーション（落下）
	position := 0
	go func() {
		for {
			position += 1

			app.QueueUpdateDraw(func() {
				lanes.FallTo(position)
			})

			if isPositionBottom(lanes.frame, position) {
				break
			}

			time.Sleep(10 * time.Millisecond)
		}
	}()

	return lanes.frame
}

// isPositionBottom 下部か判定する
func isPositionBottom(root *tview.Frame, position int) bool {
	_, _, _, screenHeight := root.GetRect()
	return position > screenHeight-2
}
