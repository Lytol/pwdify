package main

import "github.com/lytol/pwdify/pkg/pwdify"

type state struct {
	password string
	files    []string
	ch       chan pwdify.Status
	status   []pwdify.Status
}

func (s *state) PercentCompleted() float64 {
	if len(s.status) == 0 || len(s.files) == 0 {
		return 0.0
	}

	return float64(s.CompleteCount()) / float64(s.TotalCount())
}

func (s *state) CompleteCount() int {
	return len(s.status)
}

func (s *state) ErrorCount() int {
	errors := 0
	for _, status := range s.status {
		if status.Error != nil {
			errors++
		}
	}
	return errors
}

func (s *state) TotalCount() int {
	return len(s.files)
}

func (s *state) Completed() bool {
	return s.CompleteCount() == s.TotalCount()
}
