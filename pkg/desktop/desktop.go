package desktop

import (
	"reflect"

	"fyne.io/fyne"
)

type MoneyFlow struct {
	OperationFlow OperationServices
	InvestFlow    InvestServices
}

type OperationServices struct {
	Title        map[string]float64
	IncomeAction map[string]float64
	Payout       map[string]float64
}

type InvestServices struct {
	IncomeAction string
	Payout       string
}

type GetFlows interface {
	OperationServicesGet()
}

func (o *OperationServices) OperationServicesGet() OperationServices {
	titleMap := make(map[string]float64)
	incomeAction := make(map[string]float64)
	payout := make(map[string]float64)
	var countIncomeActions float64
	var countPayout float64

	incomeAction["itest"] = 20.1
	incomeAction["itest2"] = 22.3
	payout["ptest"] = 30.3
	payout["ptest2"] = 43.2

	for _, value := range incomeAction {
		countIncomeActions = countIncomeActions + value
	}
	for _, value := range payout {
		countPayout = countPayout + value
	}
	titleMap["Поступления по операционной деятельности"] = countIncomeActions
	titleMap["Выплаты по операционной деятельности"] = countPayout
	operStruct := OperationServices{
		Title:        titleMap,
		IncomeAction: incomeAction,
		Payout:       payout,
	}
	return operStruct
}

func (o *OperationServices) OperServicesGetTitles() ([]string, []float64) {
	var arrOfCounts []float64
	var arrOfTitles []string
	var pushElem reflect.Value

	operTest := o.OperationServicesGet()
	v := reflect.ValueOf(operTest.Title)

	for i := 0; i < v.Len(); i++ {
		pushElem = v.MapIndex(v.MapKeys()[i])
		arrOfCounts = append(arrOfCounts, pushElem.Float())

	}
	for _, val := range v.MapKeys() {
		arrOfTitles = append(arrOfTitles, val.String())
	}

	return arrOfTitles, arrOfCounts
}

type AppMain interface {
	MainDesktopApp() fyne.Window
}
