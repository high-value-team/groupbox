package models

type SadException struct {
	Err error
}

func (e *SadException) Message() string {
	return e.Err.Error()
}

type SuprisingException struct {
	Err error
}

func (e *SuprisingException) Message() string {
	return e.Err.Error()
}

