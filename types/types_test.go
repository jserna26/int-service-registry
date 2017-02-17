package types

import "testing"

func TestTypes(t *testing.T) {
	ns := NewSystemType{
		Name:        "test",
		Description: "Test Desc",
		Status:      StatusEnabled,
	}

	ns.Status = StatusDisabled

	s := SystemType{
		ID: "1",
	}
	s.Name = "test"
	s.Description = "Test Desc"

	s.Status = StatusDisabled

}
