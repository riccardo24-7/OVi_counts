package userInterface

import (
	"ovi/pkg/desktop"
	createwidget "ovi/pkg/desktop/userInterface/createWidget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
)

type mainInterface struct {
	homePage string
	planPage string
	factPage string
}

func MainDesktopApp() fyne.Window {
	a := app.New()
	w := a.NewWindow("OVi_counts")
	var operStruct desktop.OperationServices

	arrOfTitles := operStruct.OperServicesGetTitles()

	mainPanels := mainInterface{
		homePage: "Главная",
		planPage: "План",
		factPage: "Фак",
	}

	homeLabels := createwidget.CreateLabels(mainPanels.homePage, mainPanels.planPage, mainPanels.factPage)
	listTitles := createwidget.CreateListsLabel(arrOfTitles)
	listEntry := createwidget.CreateListsEntry(arrOfTitles)

	planContainer := container.New(layout.NewGridLayout(2), listTitles, listEntry)
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon(mainPanels.homePage, theme.HomeIcon(), homeLabels[0]),
		container.NewTabItem(mainPanels.planPage, planContainer),
		container.NewTabItem(mainPanels.factPage, homeLabels[2]),
	)

	tabs.SetTabLocation(container.TabLocationTrailing)
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(1000, 800))
	w.CenterOnScreen()

	return w
}
