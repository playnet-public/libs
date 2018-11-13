package problems

import (
	"fmt"
)

//DefaultType implements the default type content described in RFC 7807
const DefaultType = "about:blank"

//Problem implements the RFC 7807 "problem"/error standard
//Additional error field can be defined in custom struct implementing the Problem struct
type Problem struct {
	Title    string `json:"title"`
	Detail   string `json:"detail,omitempty"`
	Type     string `json:"type,omitempty"`
	Instance string `json:"instance,omitempty"`
	Status   int    `json:"status,omitempty"`
}

//New creates a new Problem with all object details described in RFC 7807
func New(title, detail string) *Problem {
	return &Problem{
		Title:  title,
		Detail: detail,
		Type:   DefaultType,
	}
}

//Error returns the error of the problem
func (p *Problem) Error() string {
	return fmt.Sprintf("Title: %s; Detail: %s", p.Title, p.Detail)
}

//SetTitle sets the title field of the problem, specified in RFC 7807
func (p *Problem) SetTitle(title string) *Problem {
	p.Title = title
	return p
}

//SetDetail sets the detail field of the problem, specified in RFC 7807
func (p *Problem) SetDetail(detail string) *Problem {
	p.Detail = detail
	return p
}

//SetType sets the detail field of the problem, specified in RFC 7807
func (p *Problem) SetType(t string) *Problem {
	if t == "" {
		p.Type = DefaultType
	} else {
		p.Type = t
	}

	return p
}

//SetInstance sets the instance field of the problem, specified in RFC 7807
func (p *Problem) SetInstance(instance string) *Problem {
	p.Instance = instance
	return p
}

//SetStatus sets the status field of the problem, specified in RFC 7807
func (p *Problem) SetStatus(status int) *Problem {
	p.Status = status
	return p
}
