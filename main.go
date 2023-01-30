package main

import (
	"main/lib"

	"github.com/mattn/go-runewidth"
	"github.com/rivo/tview"
)

func main() {
	runewidth.DefaultCondition = &runewidth.Condition{EastAsianWidth: false}
	//設定ファイル読み込み
	config := lib.LoadConfig("config.toml")
	//fmt.Println(config)
	//メイン処理
	app := tview.NewApplication()
	layout := lib.CreateLayout(&config)
	app.SetRoot(layout, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

/*
package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("List item 1", "Some explanatory text", 'a', nil).
		AddItem("List item 2", "Some explanatory text", 'b', nil).
		AddItem("List item 3", "Some explanatory text", 'c', nil).
		AddItem("List item 4", "Some explanatory text", 'd', nil).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
*/
