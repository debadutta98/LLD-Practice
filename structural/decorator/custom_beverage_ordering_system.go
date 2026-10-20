package decorator

import "fmt"

type Beverage interface {
	GetDescription() string
	GetCost() float64
}

type Espresso struct{}

func NewEspresso() Beverage {
	return new(Espresso)
}

func (e *Espresso) GetDescription() string {
	return "Espresso"
}

func (e *Espresso) GetCost() float64 {
	return 2.50
}

type Tea struct{}

func NewTea() Beverage {
	return new(Tea)
}

func (t *Tea) GetDescription() string {
	return "Tea"
}

func (t *Tea) GetCost() float64 {
	return 1.50
}

type MilkDecorator struct {
	beverage Beverage
}

func NewMilkDecorator(beverage Beverage) Beverage {
	return &MilkDecorator{beverage}
}

func (m *MilkDecorator) GetDescription() string {
	return fmt.Sprintf("%s, %s", m.beverage.GetDescription(), "Milk")
}

func (m *MilkDecorator) GetCost() float64 {
	return m.beverage.GetCost() + 0.50
}

type MochaDecorator struct {
	beverage Beverage
}

func NewMochaDecorator(beverage Beverage) Beverage {
	return &MochaDecorator{beverage}
}

func (m *MochaDecorator) GetDescription() string {
	return fmt.Sprintf("%s, %s", m.beverage.GetDescription(), "Mocha")
}

func (m *MochaDecorator) GetCost() float64 {
	return m.beverage.GetCost() + 0.75
}

type WhipDecorator struct {
	beverage Beverage
}

func NewWhipDecorator(beverage Beverage) Beverage {
	return &WhipDecorator{beverage}
}

func (w *WhipDecorator) GetDescription() string {
	return fmt.Sprintf("%s, %s", w.beverage.GetDescription(), "Whip")
}

func (w *WhipDecorator) GetCost() float64 {
	return w.beverage.GetCost() + 0.40
}

type OatMilkDecorator struct {
	beverage Beverage
}

func NewOatMilkDecorator(beverage Beverage) Beverage {
	return &OatMilkDecorator{beverage}
}

func (o *OatMilkDecorator) GetDescription() string {
	return fmt.Sprintf("%s, %s", o.beverage.GetDescription(), "Oat Milk")
}

func (o *OatMilkDecorator) GetCost() float64 {
	return o.beverage.GetCost() + 0.80
}
