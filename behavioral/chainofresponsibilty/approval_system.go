package chainofresponsibilty

import (
	"fmt"
)

type ExpenseApprover interface {
	Approve(expense float64) bool
	SetNext(ExpenseApprover)
}

type BaseApprover struct {
	next ExpenseApprover
}

func (b *BaseApprover) SetNext(e ExpenseApprover) {
	b.next = e
}

func (b *BaseApprover) ApproveNext(expense float64) bool {
	if b.next != nil {
		return b.next.Approve(expense)
	}
	return false
}

type TeamLeadApprover struct {
	BaseApprover
}

func NewTeamLeadApprover() ExpenseApprover {
	return new(TeamLeadApprover)
}

func (t *TeamLeadApprover) Approve(expense float64) bool {
	if expense <= 1000 {
		fmt.Println("Expense is approved by team lead")
		return true
	}
	return t.ApproveNext(expense)
}

type ManagerApprover struct {
	BaseApprover
}

func NewManagerApprover() ExpenseApprover {
	return new(ManagerApprover)
}

func (m *ManagerApprover) Approve(expense float64) bool {
	if expense <= 5000 {
		fmt.Println("Expense is approved by Manager")
		return true
	}
	return m.ApproveNext(expense)
}

type DirectorApprover struct {
	BaseApprover
}

func NewDirectorApporver() ExpenseApprover {
	return new(DirectorApprover)
}

func (d *DirectorApprover) Approve(expense float64) bool {
	if expense <= 10_000 {
		fmt.Println("Expense is approved by Director")
		return true
	}
	return d.ApproveNext(expense)
}

type CEOApprover struct {
	BaseApprover
}

func NewCEOApporver() ExpenseApprover {
	return new(CEOApprover)
}

func (c *CEOApprover) Approve(expense float64) bool {
	if expense <= 50_000 {
		fmt.Println("Expense is approved by CEO")
		return true
	}
	return c.ApproveNext(expense)
}

func CanApprove(e ExpenseApprover, expense float64) {
	if isApproved := e.Approve(expense); !isApproved {
		fmt.Printf("Expense of $%0.2f can't be approve\n", expense)
	}
}
