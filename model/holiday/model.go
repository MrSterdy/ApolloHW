package holiday

import "github.com/MrSterdy/ApolloHW/util"

type Holiday struct {
	StartDate *util.Date `bun:",notnull" json:"start_date"`
	EndDate   *util.Date `json:"end_date"`
}

func (h Holiday) IntersectsWith(date *util.Date) bool {
	if h.EndDate == nil {
		return h.StartDate.Equal(date.Time)
	}

	return !date.Before(h.StartDate.Time) && !date.After(h.EndDate.Time)
}

func (h Holiday) String() string {
	if h.EndDate == nil {
		return h.StartDate.String()
	}

	return h.StartDate.String() + " - " + h.EndDate.String()
}
