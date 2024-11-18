package part_test

import (
	"brickstorage/domain/part"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestService(t *testing.T) {
	t.Run("can add partview and toggle completion", func(t *testing.T) {
		service := part.List{}
		assert.Equal(t, 0, len(service.Parts()))

		someTask := "kill react"
		service.Add(someTask)
		service.Add("blah blah 1")
		service.Add("blah blah 2")
		service.Add("blah blah 3")

		parts := service.Parts()
		assert.False(t, parts[0].Complete)
		id := parts[0].ID

		service.ToggleDone(id)

		parts = service.Parts()
		assert.True(t, parts[0].Complete)

		service.ToggleDone(id)

		parts = service.Parts()
		assert.False(t, parts[0].Complete)
	})

	t.Run("rename", func(t *testing.T) {
		service := part.List{}
		service.Add("kill react")
		service.Add("blah blah 1")

		parts := service.Parts()
		assert.Equal(t, "kill react", parts[0].Description)
		id := parts[0].ID

		service.Rename(id, "kill react and redux")

		parts = service.Parts()
		assert.Equal(t, "kill react and redux", parts[0].Description)
		assert.Equal(t, "blah blah 1", parts[1].Description)
	})

	t.Run("delete", func(t *testing.T) {
		service := part.List{}
		service.Add("kill react")
		service.Add("blah blah 1")
		service.Add("blah blah 2")
		service.Add("blah blah 3")

		parts := service.Parts()
		assert.Equal(t, 4, len(parts))
		id := parts[0].ID

		service.Delete(id)

		parts = service.Parts()
		assert.Equal(t, 3, len(parts))
		assert.Equal(t, "blah blah 1", parts[0].Description)
		assert.Equal(t, "blah blah 2", parts[1].Description)
		assert.Equal(t, "blah blah 3", parts[2].Description)
	})

	t.Run("reorder", func(t *testing.T) {
		service := part.List{}
		service.Add("kill react")
		service.Add("blah blah 1")
		service.Add("blah blah 2")
		service.Add("blah blah 3")

		parts := service.Parts()
		assert.Equal(t, 4, len(parts))
		assert.Equal(t, "kill react", parts[0].Description)

		// reorder
		service.ReOrder([]string{
			parts[3].ID.String(),
			parts[2].ID.String(),
			parts[1].ID.String(),
			parts[0].ID.String(),
		})

		parts = service.Parts()
		assert.Equal(t, 4, len(parts))
		assert.Equal(t, "blah blah 3", parts[0].Description)
		assert.Equal(t, "blah blah 2", parts[1].Description)
		assert.Equal(t, "blah blah 1", parts[2].Description)
		assert.Equal(t, "kill react", parts[3].Description)
	})

	t.Run("search", func(t *testing.T) {
		service := part.List{}
		service.Add("kill react")
		service.Add("blah blah 1")
		service.Add("blah blah 2")
		service.Add("blah blah 3")

		parts := service.Search("blah")
		assert.Equal(t, 3, len(parts))
		assert.Equal(t, "blah blah 1", parts[0].Description)
		assert.Equal(t, "blah blah 2", parts[1].Description)
		assert.Equal(t, "blah blah 3", parts[2].Description)
	})

	t.Run("empty the list", func(t *testing.T) {
		service := part.List{}
		service.Add("kill react")
		service.Add("blah blah 1")
		service.Add("blah blah 2")
		service.Add("blah blah 3")

		parts := service.Parts()
		assert.Equal(t, 4, len(parts))

		service.Empty()

		parts = service.Parts()
		assert.Equal(t, 0, len(parts))
	})

}
