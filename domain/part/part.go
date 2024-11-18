package part

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Part struct {
	ID          uuid.UUID
	Description string
	CreatedAt   time.Time
	Complete    bool
}

type List struct {
	parts []Part
}

func (s *List) Add(description string) {
	s.parts = append(s.parts, Part{
		ID:          uuid.New(),
		Description: description,
		CreatedAt:   time.Now(),
	})
}

func (s *List) Rename(id uuid.UUID, name string) Part {
	i := s.indexOf(id)
	s.parts[i].Description = name
	return s.parts[i]
}

func (s *List) Parts() []Part {
	return s.parts
}

func (s *List) ToggleDone(id uuid.UUID) Part {
	i := s.indexOf(id)
	s.parts[i].Complete = !s.parts[i].Complete
	return s.parts[i]
}

func (s *List) Delete(id uuid.UUID) {
	i := s.indexOf(id)
	s.parts = append(s.parts[:i], s.parts[i+1:]...)
}

func (s *List) ReOrder(ids []string) {
	var uuids []uuid.UUID
	for _, id := range ids {
		uuids = append(uuids, uuid.MustParse(id))
	}

	var newList []Part
	for _, id := range uuids {
		newList = append(newList, s.parts[s.indexOf(id)])
	}

	s.parts = newList
}

func (s *List) Search(search string) []Part {
	search = strings.ToLower(search)
	var results []Part
	for _, part := range s.parts {
		if strings.Contains(strings.ToLower(part.Description), search) {
			results = append(results, part)
		}
	}
	return results
}

func (s *List) Get(id uuid.UUID) Part {
	return s.parts[s.indexOf(id)]
}

func (s *List) Empty() {
	s.parts = []Part{}
}

func (s *List) indexOf(id uuid.UUID) int {
	return slices.IndexFunc(s.parts, func(part Part) bool {
		return part.ID == id
	})
}
