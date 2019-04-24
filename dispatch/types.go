package dispatch

type DispatchID string

type Status string

type DateTime interface {
	Start() string
	End() string
	Duration() string
}

func NewDateTime() DateTime {
	return dateTime{}
}

type dateTime struct{}

func (dt dateTime) Start() string {
	return ""
}

func (dt dateTime) End() string {
	return ""
}

func (dt dateTime) Duration() string {
	return ""
}
