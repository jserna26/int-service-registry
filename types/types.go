package types

//StatusType Type
type StatusType string

const StatusEnabled StatusType = "TRUE"
const StatusDisabled StatusType = "FALSE"

//System Type
type NewSystemType struct {
	Name        string
	Description string
	Status      StatusType
}

type SystemType struct {
	ID string
	NewSystemType
}

type ErrorType struct {
	Message string
}

// TODO: No longer require SystemType. Remove all dependencies
