package lib

import (
	"fmt"

	"github.com/rivo/tview"
)

/*
func welcomePageLayout() *tview.Flex {
	WelcomePage := tview.NewFlex()
	image := tview.NewImage()
	logoBin := Load_file("./image/logo.png")
	logo, imageDecodeErr := png.Decode(bytes.NewReader(logoBin))
	if imageDecodeErr != nil {
		panic(imageDecodeErr)
	}
	image.SetImage(logo)
	WelcomePage.SetDirection(tview.FlexColumnCSS).
		AddItem(image, 0, 1, false)
	return WelcomePage
}
*/

func setLangPageLayout() *tview.Flex {
	langs := map[]type
	SetLocalePage := tview.NewFlex()
	content := tview.NewList()
	for i, v := range langs {
		content.AddItem()
	}
	content.SetBorder(true).SetTitle("Select Language")
	SetLocalePage.SetDirection(tview.FlexRowCSS).
		AddItem(content, 0, 1, false)
	return SetLocalePage
}

/*
func setKeymapPageLayout() *tview.Flex {
	SetKeymapPage := tview.NewFlex()
	return SetKeymapPage
}

func setLocaltimePageLayout() *tview.Flex {
	SetLocalePage := tview.NewFlex()
	return SetLocalePage
}
*/

func CreateLayout(config *ConfigData) *tview.Flex {
	Layout := tview.NewFlex()
	pages := tview.NewPages()
	margin := tview.NewBox()
	//welcomePage := welcomePageLayout()
	setLocalePage := setLangPageLayout()
	//setKeymapPage := setKeymapPageLayout()
	//setLocaltimePage := setLocaltimePageLayout()
	pages. //AddPage("WelcomePage", welcomePage, true, true)
		AddPage("SetLocalePage", setLocalePage, true, true)
	//AddPage("SetKeymapPage", setKeymapPage, true, true).
	//AddPage("SetLocaltimePage", setLocaltimePage, true, true)
	Layout.SetDirection(tview.FlexColumnCSS).
		AddItem(margin, 2, 0, false).
		AddItem(pages, 0, 1, false)
	Layout.SetTitle(fmt.Sprintf(" %s Installer ", *config.DistroName)).SetBorder(true)
	return Layout
}
