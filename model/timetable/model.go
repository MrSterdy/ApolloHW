package timetable

import (
	"time"

	"github.com/MrSterdy/ApolloHW/util"
)

type Timetable interface {
	GetOffset() int

	GetSubjects() Subjects
}

type Base struct {
	Offset int `bun:"type:smallint,notnull" json:"offset"`

	Subjects Subjects `bun:"type:longtext" json:"subjects"`
}

func (b Base) GetOffset() int {
	return b.Offset
}

func (b Base) GetSubjects() Subjects {
	return b.Subjects
}

type Date struct {
	Base

	Date *util.Date `bun:",pk" json:"date"`

	Note string `bun:",nullzero" json:"note,omitempty"`
}

type Weekday struct {
	Base

	Day time.Weekday `bun:",pk" json:"day"`

	SubjectLength int `bun:",notnull" json:"subject_length"`
	SubjectBreak  int `bun:",notnull" json:"subject_break"`
}

type Subject struct {
	Name string `bun:",pk,notnull" json:"name,omitempty"`

	Length int `json:"length" bun:"-"`
	Break  int `json:"break" bun:"-"`

	Teacher string `bun:",nullzero" json:"teacher,omitempty"`

	Classroom string `bun:",nullzero" json:"classroom,omitempty"`

	Homework *Homework `json:"homework,omitempty" bun:"-"`
}

type Subjects []*Subject

func (ss Subjects) Diff(other Subjects) (added, removed Subjects) {
Remove:
	for _, s := range ss {
		for i, rs := range removed {
			if rs.Name == s.Name && s.Homework != nil {
				removed[i] = s

				continue Remove
			}
		}

		for _, os := range other {
			if s.Name == os.Name {
				continue Remove
			}
		}

		removed = append(removed, s)
	}

Add:
	for _, os := range other {
		for i, as := range added {
			if as.Name == os.Name && os.Homework != nil {
				added[i] = os

				continue Add
			}
		}

		for _, s := range ss {
			if os.Name == s.Name {
				continue Add
			}
		}

		added = append(added, os)
	}

	return
}

type Homework struct {
	Text string `json:"text,omitempty"`

	Images []string `json:"images,omitempty"`
}
