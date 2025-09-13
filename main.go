package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	// アプリケーションの初期化
	app := tview.NewApplication()

	// ノートの初期化
	box := newBox()
	note1 := newNote(box)
	note2 := newNote(box)
	note3 := newNote(box)
	note4 := newNote(box)

	// フレームの初期化
	frame1 := newFrame(note1)
	frame2 := newFrame(note2)
	frame3 := newFrame(note3)
	frame4 := newFrame(note4)

	// ノーツの初期化
	notes := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(frame1, 0, 1, false).
		AddItem(frame2, 0, 1, false).
		AddItem(frame3, 0, 1, false).
		AddItem(frame4, 0, 1, false)


	// スペースを設定
	wrapper := tview.NewFrame(notes).SetBorders(1, 1, 0, 0, 35, 35)

	// ボックスの初期位置
	position := 0

	// ゴルーチンを使用してボックスを下に動かす
	go func() {
		for {
			position += 1

			// アプリケーションの描画を更新
			app.QueueUpdateDraw(func() {
				note1.Clear()
				note2.Clear()
				note3.Clear()
				note4.Clear()

				// ノーツの前に空白のスペースを挿入
				note1.AddItem(tview.NewBox(), position, 1, false)
				note2.AddItem(tview.NewBox(), position, 1, false)
				note3.AddItem(tview.NewBox(), position, 1, false)
				note4.AddItem(tview.NewBox(), position, 1, false)

				note1.AddItem(box, 1, 1, false)
				note2.AddItem(box, 1, 1, false)
				note3.AddItem(box, 1, 1, false)
				note4.AddItem(box, 1, 1, false)
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

func newBox() *tview.Box {
	return tview.NewBox().SetBackgroundColor(tcell.ColorWhite)
}

func newNote(box *tview.Box) (*tview.Flex) {
	note := tview.NewFlex().
						SetDirection(tview.FlexRow).
						AddItem(box, 1, 1, false)
	return note
}

func newFrame(note *tview.Flex) *tview.Frame {
	return tview.NewFrame(note).SetBorders(0, 0, 0, 0, 2, 2)
}

func isPositionBottom(root *tview.Frame, position int) bool {
	// ターミナルの高さを取得
	_, _, _, screenHeight := root.GetRect()
	return position > screenHeight - 3
}
