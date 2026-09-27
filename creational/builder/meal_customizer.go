package builder

import "fmt"

type Size string

const (
	Small  Size = "Small"
	Medium Size = "Medium"
	Large  Size = "Large"
)

type MainItem struct {
	Name    string
	Protein string
	Price   float64
}

type SideItem struct {
	Name  string
	Size  Size
	Price float64
}

type Drink struct {
	Name   string
	Size   Size
	HasIce bool
	Price  float64
}

type Dessert struct {
	Name  string
	Price float64
}

type Meal struct {
	mainItem   MainItem
	sideItem   *SideItem
	drink      *Drink
	dessert    *Dessert
	totalPrice float64
}

func (m *Meal) DisplayMealSummary() {
	fmt.Printf("Main: %s (%s) - $%.2f\n", m.mainItem.Name, m.mainItem.Protein, m.mainItem.Price)
	if m.sideItem != nil {
		fmt.Printf("Side: %s [%s] - $%.2f\n", m.sideItem.Name, m.sideItem.Size, m.sideItem.Price)
	}
	if m.drink != nil {
		ice := "no ice"
		if m.drink.HasIce {
			ice = "with ice"
		}
		fmt.Printf("Drink: %s [%s, %s] - $%.2f\n", m.drink.Name, m.drink.Size, ice, m.drink.Price)
	}
	if m.dessert != nil {
		fmt.Printf("Dessert: %s - $%.2f\n", m.dessert.Name, m.dessert.Price)
	}
	fmt.Printf("Total: $%.2f\n\n", m.totalPrice)
}

type MealBuilder struct {
	meal *Meal
}

func NewMealBuilder(main MainItem) *MealBuilder {
	return &MealBuilder{
		meal: &Meal{
			mainItem:   main,
			totalPrice: main.Price,
		},
	}
}

func (b *MealBuilder) WithSide(side SideItem) *MealBuilder {
	b.meal.sideItem = &side
	b.meal.totalPrice += side.Price
	return b // Returns pointer to builder for chaining
}

func (b *MealBuilder) WithDrink(drink Drink) *MealBuilder {
	b.meal.drink = &drink
	b.meal.totalPrice += drink.Price
	return b
}

func (b *MealBuilder) WithDessert(dessert Dessert) *MealBuilder {
	b.meal.dessert = &dessert
	b.meal.totalPrice += dessert.Price
	return b
}

func (b *MealBuilder) Build() *Meal {
	return b.meal
}
