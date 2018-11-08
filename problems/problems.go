package errors

import (
	"fmt"
)

//Problem implements the RFC7807 "problem"/error standard
//Additional error field can be defined in custom struct implementing the Problem struct
type Problem struct {
	Title    string `json:"title"`
	Detail   string `json:"detail,omitempty"`
	Type     string `json:"type,omitempty"`
	Instance string `json:"instance,omitempty"`
	Status   int    `json:"status,omitempty"`
}

//Error implements the error interface
func (p *Problem) Error() string {
	return fmt.Sprintf("Title: %s; Detail: %s", p.Title, p.Detail)
}

//DefaultType implements the default type content described in RFC7807
const DefaultType = "about:blank"

//SetTitle sets the title of the problem object
func (p *Problem) SetTitle(title string) *Problem {
	p.Title = title
	return p
}

//SetDetail sets the detail info of the problem object
func (p *Problem) SetDetail(detail string) *Problem {
	p.Detail = detail
	return p
}

//SetType sets the type of the problem object
func (p *Problem) SetType(typ string) *Problem {
	if typ == "" {
		p.Type = DefaultType
	} else {
		p.Type = typ
	}

	return p
}

//SetInstance sets the instance of the problem object
func (p *Problem) SetInstance(instance string) *Problem {
	p.Instance = instance
	return p
}

//SetStatus sets the HTTP status of the problem object HTTP request
func (p *Problem) SetStatus(status int) *Problem {
	p.Status = status
	return p
}

//New creates a new Problem with all object details described in RFC7807
func New() *Problem {
	return &Problem{
		Type: DefaultType,
	}
}

//NewFromError returns a new Problem object from an error interface
//The error message is replaced with the detail field
func NewFromError(err error) *Problem {
	return &Problem{
		Type:   DefaultType,
		Detail: err.Error(),
	}
}
