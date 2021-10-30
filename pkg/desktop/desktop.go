package desktop

import (
	"reflect"

	"fyne.io/fyne"
)

type MoneyFlow struct {
	OperationFlow OperationServices
	InvestFlow    InvestActions
}

type OperationServices struct {
	Title            string
	IncomeAction     string
	ProcentByDeposit string
	OtherInput       string
	Payout           string
	AdminTitle       string
	AuditTitle       string
	BankTitle        string
	Payment          string
	OtherServices    string
}

type InvestActions struct {
}

func (o *OperationServices) OperServicesGetTitles() []string {
	operTest := OperationServices{
		IncomeAction:     "Поступления по операционной деятельности",
		ProcentByDeposit: "Проценты по депозитам",
		OtherInput:       "Прочие поступления",
		Payout:           "Выплаты по операционной деятельности",
		AdminTitle:       "Услуги администрирующей компании",
		AuditTitle:       "Аудит",
		BankTitle:        "Банковские комиссии",
		Payment:          "Заработная плата",
		OtherServices:    "Прочие услуги",
	}
	arrOfTitles := []string{}
	v := reflect.ValueOf(operTest)

	for i := 0; i < v.NumField(); i++ {
		pushElem := v.Field(i).String()
		arrOfTitles = append(arrOfTitles, pushElem)
	}
	return arrOfTitles
}

type AppMain interface {
	MainDesktopApp() fyne.Window
	GetTitles() []string
}
