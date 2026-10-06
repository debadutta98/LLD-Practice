package composite

import (
	"fmt"
	"slices"
	"strings"
)

type Entity interface {
	Display(indent int)
	GetSalary() float64
}

type Composite interface {
	Entity
	Add(Entity)
	Remove(Entity)
}

type Employee struct {
	name   string
	title  string
	salary float64
}

func NewEmployee(name, title string, salary float64) Entity {
	return &Employee{
		name:   name,
		title:  title,
		salary: salary,
	}
}

func (e *Employee) Display(indent int) {
	padding := strings.Repeat("  ", indent)
	fmt.Printf("%s- Employee: %s | Title: %s | Salary: $%.2f\n", padding, e.name, e.title, e.salary)
}

func (e *Employee) GetSalary() float64 {
	return e.salary
}

type GroupType string

const (
	TypeDepartment   GroupType = "Department"
	TypeBusinessUnit GroupType = "Business Unit"
)

type OrganizationGroup struct {
	name      string
	groupType GroupType
	entities  []Entity
}

func NewGroup(name string, groupType GroupType, children ...Entity) Composite {
	return &OrganizationGroup{
		name:      name,
		groupType: groupType,
		entities:  children,
	}
}

func (g *OrganizationGroup) Display(indent int) {
	padding := strings.Repeat("  ", indent)
	fmt.Printf("%s+ %s: %s (Total Budget: $%.2f)\n", padding, g.groupType, g.name, g.GetSalary())

	for _, child := range g.entities {
		child.Display(indent + 1)
	}
}

func (g *OrganizationGroup) GetSalary() float64 {
	total := 0.0
	for _, child := range g.entities {
		total += child.GetSalary()
	}
	return total
}

func (g *OrganizationGroup) Add(e Entity) {
	g.entities = append(g.entities, e)
}

func (g *OrganizationGroup) Remove(e Entity) {
	for i, entity := range g.entities {
		if entity == e {
			g.entities = slices.Delete(g.entities, i, i+1)
			break
		}
	}
}

func Move(from Composite, to Composite, e Entity) {
	from.Remove(e)
	to.Add(e)
}
