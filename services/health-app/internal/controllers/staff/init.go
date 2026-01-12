package staff

import staffsv "health-app/internal/services/staff"

type StaffHandler struct {
	service staffsv.Services
}

func NewStaffHandler(service staffsv.Services) *StaffHandler {
	return &StaffHandler{
		service: service,
	}
}
