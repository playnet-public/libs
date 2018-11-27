package problems

import (
	"github.com/pkg/errors"
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
	err      error
}

//Wrap is an alias to github.com/pkg/errors Wrap function. If Problem pointer passed as error, it sets the error and a new message of Problem.
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}

	problem, ok := err.(*Problem)
	if !ok {
		return errors.Wrap(err, msg)
	}

	problem.err = errors.Wrap(problem.err, msg)
	return problem
}

//WithStack is an alias to github.com/pkg/errors WithStack function. If Problem pointer passed as error, it sets the error of Problem.
func WithStack(err error) error {
	if err == nil {
		return nil
	}

	problem, ok := err.(*Problem)
	if !ok {
		return errors.WithStack(err)
	}

	problem.err = errors.WithStack(problem.err)
	return problem
}

//WithMessage is an alias to github.com/pkg/errors WithMessage function. If Problem pointer passed as error, it sets the error, stacktrace and message of Problem.
func WithMessage(err error, msg string) error {
	if err == nil {
		return nil
	}

	problem, ok := err.(*Problem)
	if !ok {
		return errors.WithMessage(err, msg)
	}

	problem.err = errors.WithMessage(problem.err, msg)
	return problem
}

//New creates a new Problem with all object details described in RFC 7807
func New(title, detail string, status int) error {
	return newWithError(
		errors.New(detail),
		title,
		detail,
		status,
	)
}

func newWithError(err error, title, detail string, status int) *Problem {
	return &Problem{
		Type:   DefaultType,
		Title:  title,
		Detail: err.Error(),
		Status: status,
		err:    err,
	}
}

//Error returns the error of the problem
func (p Problem) Error() string {
	return p.err.Error()
}

//SetTitle sets the title field of the problem, specified in RFC 7807
func (p Problem) SetTitle(title string) *Problem {
	p.Title = title
	return &p
}

//SetDetail sets the detail field of the problem, specified in RFC 7807
func (p Problem) SetDetail(detail string) *Problem {
	p.Detail = detail
	return &p
}

//SetType sets the detail field of the problem, specified in RFC 7807
func (p Problem) SetType(t string) *Problem {
	if t == "" {
		p.Type = DefaultType
	} else {
		p.Type = t
	}

	return &p
}

//SetInstance sets the instance field of the problem, specified in RFC 7807
func (p Problem) SetInstance(instance string) *Problem {
	p.Instance = instance
	return &p
}

//SetStatus sets the status field of the problem, specified in RFC 7807
func (p Problem) SetStatus(status int) *Problem {
	p.Status = status
	return &p
}
