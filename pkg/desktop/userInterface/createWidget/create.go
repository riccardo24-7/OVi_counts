package createwidget

import (
	"fmt"

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

func CreateListLabels(labels []string) *widget.List {

	list := widget.NewList(
		func() int {
			return len(labels)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Поступления по операционной деятельности")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {

			o.(*widget.Label).SetText(labels[i])
		})

	return list
}

func CreateListEntry(length int, entries []float64) *widget.List {
	var dataStr []string
	for _, val := range entries {
		dataStr = append(dataStr, fmt.Sprint(val))
	}
	listEntries := widget.NewList(
		func() int {
			return length
		},
		func() fyne.CanvasObject {
			return widget.NewEntry()
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Entry).SetPlaceHolder(dataStr[i])
		})
	return listEntries
}
