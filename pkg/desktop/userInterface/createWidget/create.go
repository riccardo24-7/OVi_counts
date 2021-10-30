package createwidget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CreateLabels(str ...string) []*widget.Label {
	var newLabel []*widget.Label
	for _, elem := range str {
		newLabel = append(newLabel, widget.NewLabel(elem))
	}
	return newLabel
}

func CreateListsLabel(arrList []string) *widget.List {
	listTitles := widget.NewList(
		func() int {
			return len(arrList)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(arrList[i])
		})
	return listTitles
}

func CreateListsEntry(arrList []string) *widget.List {
	listEntries := widget.NewList(
		func() int {
			return len(arrList)
		},
		func() fyne.CanvasObject {
			return widget.NewEntry()
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Entry).Show()
		})
	return listEntries
}
