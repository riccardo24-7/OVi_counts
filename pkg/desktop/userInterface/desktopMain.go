package userInterface

import (
	"fmt"
	"ovi/pkg/desktop"
	createwidget "ovi/pkg/desktop/userInterface/createWidget"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type mainInterface struct {
	homePage string
	planPage string
	factPage string
}

// func countOperServices(arrList []string) []float64 {
// 	var counts []float64
// 	for i := 0; i < len(arrList); i++ {
// 		counts = append(counts, 1)
// 	}
// 	return counts
// }

func CreateContainersPattern(arrList []string, arrCounts []float64) ([]*widget.Label, []*widget.Entry) {
	var labels []*widget.Label
	var entries []*widget.Entry

	for i := 0; i < len(arrList); i++ {
		widgetLabel := widget.NewLabel(arrList[i])
		widgetEntry := widget.NewEntry()

		if i < len(arrCounts) {
			countStr := fmt.Sprint(arrCounts[i])
			widgetEntry.SetPlaceHolder(countStr)
		}
		labels = append(labels, widgetLabel)
		entries = append(entries, widgetEntry)
	}

	return labels, entries
}

func MainDesktopApp() fyne.Window {
	a := app.New()
	w := a.NewWindow("OVi_counts")
	var operStruct desktop.OperationServices

	mainPanels := mainInterface{
		homePage: "Главная",
		planPage: "План",
		factPage: "Факт",
	}

	homeLabels := createwidget.CreateLabels(mainPanels.homePage, mainPanels.planPage, mainPanels.factPage)
	arrTitles, arrCounts := operStruct.OperServicesGetTitles()
	fmt.Println("Это arrTitles", arrTitles)
	fmt.Println("Это arrCounts", arrCounts)

	// arrList, arrEntry := CreateContainersPattern(arrTitles, arrCounts)
	listTitles := createwidget.CreateListLabels(arrTitles)
	listEntries := createwidget.CreateListEntry(len(arrTitles), arrCounts)
	planContainer := container.New(layout.NewGridLayout(2), listTitles, listEntries)

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
