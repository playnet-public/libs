package problem_test

import (
	"net/http"
	"testing"

	"github.com/TheMysteriousVincent/libs/problem"
	"github.com/pkg/errors"
)

func TestWrap(t *testing.T) {
	if problem.Wrap(nil, "") != nil {
		t.Fatal("the resulting error has to be nil")
	}

	if _, ok := (problem.Wrap(errors.New("test message"), "this is appended")).(problem.Problem); ok {
		t.Fatal("the returned error has to not implement the problem interface")
	}

	if _, ok := (problem.Wrap(problem.New("test title", "test detail", http.StatusBadRequest), "some appended message")).(problem.Problem); !ok {
		t.Fatal("the returned error has to implement the Problem interface")
	}
}

func TestWithStack(t *testing.T) {
	if problem.WithStack(nil) != nil {
		t.Fatal("the resulting error has to be nil")
	}

	if _, ok := (problem.WithStack(errors.New("test message"))).(problem.Problem); ok {
		t.Fatal("the returned error has to not implement the problem interface")
	}

	if _, ok := (problem.WithStack(problem.New("test title", "test detail", http.StatusBadRequest))).(problem.Problem); !ok {
		t.Fatal("the returned error has to implement the Problem interface")
	}
}

func TestWithMessage(t *testing.T) {
	if problem.WithMessage(nil, "") != nil {
		t.Fatal("the resulting error has to be nil")
	}

	if _, ok := (problem.WithMessage(errors.New("test message"), "this is appended")).(problem.Problem); ok {
		t.Fatal("the returned error has to not implement the problem interface")
	}

	if _, ok := (problem.WithMessage(problem.New("test title", "test detail", http.StatusBadRequest), "some appended message")).(problem.Problem); !ok {
		t.Fatal("the returned error has to implement the Problem interface")
	}
}

func TestNew(t *testing.T) {
	p := problem.New("test title", "test detail", http.StatusBadRequest)

	if p.Detail != "test detail" || p.Title != "test title" || p.Status != http.StatusBadRequest {
		t.Fatal("the problem object was not returned correctly")
	}
}

func TestProblemCause(t *testing.T) {
	p := problem.New("test title", "test detail", http.StatusBadRequest)

	if p.Cause().Error() != "test detail" {
		t.Fatal("wrong error given")
	}
}

func TestProblemError(t *testing.T) {
	p := problem.New("test title", "test detail", http.StatusBadRequest)

	if p.Error() != "test detail" {
		t.Fatal("wrong error message given")
	}
}

func TestProblemSetTitle(t *testing.T) {
	p := problem.New("", "", 0)
	p.SetTitle("title")

	if p.Title != "title" {
		t.Fatal("the title has to be changed")
	}
}

func TestProblemSetDetail(t *testing.T) {
	p := problem.New("", "", 0)
	p.SetDetail("detail")

	if p.Detail != "detail" {
		t.Fatal("the detail has to be changed")
	}
}

func TestProblemSetType(t *testing.T) {
	p := problem.New("", "", 0)
	p.SetType("type")

	if p.Type != "type" {
		t.Fatal("the status has to be changed")
	}

	p.SetType("")

	if p.Type != problem.DefaultType {
		t.Fatal("the type ")
	}
}

func TestProblemSetInstance(t *testing.T) {
	p := problem.New("", "", 0)
	p.SetInstance("test instance")

	if p.Instance != "test instance" {
		t.Fatal("the status has to be changed")
	}
}

func TestProblemSetStatus(t *testing.T) {
	p := problem.New("", "", 0)
	p.SetStatus(http.StatusBadRequest)

	if p.Status != http.StatusBadRequest {
		t.Fatal("the status has to be changed")
	}
}
