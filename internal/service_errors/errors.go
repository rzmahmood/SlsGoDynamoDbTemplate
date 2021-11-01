package service_errors

import "fmt"

type MissingQueryParameter struct{
	MissingParam string
}

func (m *MissingQueryParameter) Error() string {
	return fmt.Sprintf("Missing Query Parameter: %v", m.MissingParam)
}

type InvalidHttpMethod struct{
	CurrentHttpMethod string
	CorrectHttpMethod string
}

func (m *InvalidHttpMethod ) Error() string {
	return fmt.Sprintf("Invalid Http Method '%v', expected Http Method: %v", m.CurrentHttpMethod, m.CorrectHttpMethod)
}
