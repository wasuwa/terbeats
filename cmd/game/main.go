package main

import (
	"time"

	"github.com/rivo/tview"
	ui "github.com/wasuwa/terbeats/internal/ui"
)

func main() {
	// アプリケーションの初期化
	app := tview.NewApplication()

	// レーンの初期化（box/flex/frame を束ねる）
	lane1 := ui.NewLane()
	lane2 := ui.NewLane()
	lane3 := ui.NewLane()
	lane4 := ui.NewLane()

	// レーン群のレイアウト
	notes := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(lane1.Frame, 0, 1, false).
		AddItem(lane2.Frame, 0, 1, false).
		AddItem(lane3.Frame, 0, 1, false).
		AddItem(lane4.Frame, 0, 1, false)

	// スペースを設定（左右の余白を縮小して各レーンの幅を確保）
	wrapper := tview.NewFrame(notes).SetBorders(1, 1, 0, 0, 2, 2)

	// ボックスの初期位置
	position := 0

	// ゴルーチンを使用してボックスを下に動かす
	go func() {
		for {
			position += 1

			// アプリケーションの描画を更新
			app.QueueUpdateDraw(func() {
				lane1.Flex.Clear()
				lane2.Flex.Clear()
				lane3.Flex.Clear()
				lane4.Flex.Clear()

				// ノーツの前に空白のスペースを挿入
				lane1.Flex.AddItem(tview.NewBox(), position, 1, false)
				lane2.Flex.AddItem(tview.NewBox(), position, 1, false)
				lane3.Flex.AddItem(tview.NewBox(), position, 1, false)
				lane4.Flex.AddItem(tview.NewBox(), position, 1, false)

				lane1.Flex.AddItem(lane1.Box, 1, 1, false)
				lane2.Flex.AddItem(lane2.Box, 1, 1, false)
				lane3.Flex.AddItem(lane3.Box, 1, 1, false)
				lane4.Flex.AddItem(lane4.Box, 1, 1, false)
			})

			// 画面の一番下である場合は停止
			if isPositionBottom(wrapper, position) {
				break
			}

			// スリープ
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// アプリケーションを起動
	if err := app.SetRoot(wrapper, true).Run(); err != nil {
		panic(err)
	}
}

func isPositionBottom(root *tview.Frame, position int) bool {
	// ターミナルの高さを取得
	_, _, _, screenHeight := root.GetRect()
	return position > screenHeight-3
}
