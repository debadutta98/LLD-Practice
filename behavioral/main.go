package main

import (
	. "behavioural/chainofresponsibilty"
)

func main() {
	teamLeadApprover := NewTeamLeadApprover()
	managerApprover := NewManagerApprover()
	teamLeadApprover.SetNext(managerApprover)
	directorApporver := NewDirectorApporver()
	managerApprover.SetNext(directorApporver)
	ceoApprover := NewCEOApporver()
	directorApporver.SetNext(ceoApprover)
	CanApprove(teamLeadApprover, 1000)
	CanApprove(teamLeadApprover, 51000)
	CanApprove(teamLeadApprover, 4000)
}
