package parthttp_test

import (
	"brickstorage/domain/part"
	"brickstorage/web/parthttp"
	"brickstorage/web/views"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alecthomas/assert/v2"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func TestNewPartHandler(t *testing.T) {
	partList := &part.List{}
	templates, err := views.NewTemplates()
	assert.NoError(t, err)
	handler, err := parthttp.NewPartHandler(partList, views.NewPartView(templates), views.NewIndexView(templates))
	assert.NoError(t, err)

	server := httptest.NewServer(handler)
	defer server.Close()

	launcher := launcher.New().Headless(true).MustLaunch()
	rod := rod.New().Timeout(20 * time.Second).ControlURL(launcher).MustConnect()

	partListPage := &PartPage{
		Rod:  rod,
		Page: rod.MustPage(server.URL),
		URL:  server.URL,
	}

	it := func(description string, f func(*testing.T)) {
		t.Helper()
		partList.Empty()
		partListPage.Home()
		t.Run(description, f)
	}

	it("can add some parts", func(t *testing.T) {
		partListPage.Add("Eat cheese")
		partListPage.Add("Drink port")

		assert.Equal(t, 2, len(partList.Parts()))
		assert.Equal(t, "Eat cheese", partList.Parts()[0].Description)
		assert.Equal(t, "Drink port", partList.Parts()[1].Description)
	})

	it("can edit a partview", func(t *testing.T) {
		partListPage.Add("Eat cheese")
		partListPage.Edit("Eat cheese", "Eat cheese and crackers")
		assert.Equal(t, "Eat cheese and crackers", partList.Parts()[0].Description)
	})

	it("can delete a partview", func(t *testing.T) {
		partListPage.Add("Eat cheese")
		partListPage.Add("Drink port")
		assert.Equal(t, 2, len(partList.Parts()))

		partListPage.Delete("Drink port")
		assert.Equal(t, 1, len(partList.Parts()))
		assert.Equal(t, "Eat cheese", partList.Parts()[0].Description)
	})

	it("can mark a partview as done", func(t *testing.T) {
		partListPage.Add("Mark this as done")
		assert.False(t, partList.Parts()[0].Complete)

		partListPage.Toggle("Mark this as done")
		assert.True(t, partList.Parts()[0].Complete)
		partListPage.Toggle("Mark this as done")
		assert.False(t, partList.Parts()[0].Complete)
	})

	//t.Run("partview: attempts at testing drag and drog", func(t *testing.T) {
	//	t.Skip("pft")
	//	portBox, _ := page.MustElement(`[data-description="Drink port"]`).Shape()
	//	log.Println(portBox.OnePointInside())
	//	cheeseBox, _ := page.MustElement(`[data-description="Eat cheese"]`).Shape()
	//	log.Println(cheeseBox.OnePointInside())
	//
	//	mouse := page.Mouse
	//
	//	assert.NoError(t, mouse.MoveTo(*portBox.OnePointInside()))
	//	mouse.MustDown(proto.InputMouseButtonLeft)
	//	mouse.MoveLinear(*cheeseBox.OnePointInside(), 3)
	//	mouse.MustUp(proto.InputMouseButtonLeft)
	//
	//	portBox, _ = page.MustElement(`[data-description="Drink port"]`).Shape()
	//	log.Println(portBox.OnePointInside())
	//
	//	page = rod.MustPage(server.URL)
	//	assert.Equal(t, "Drink port", partList.Parts()[0].Description)
	//	assert.Equal(t, "Eat cheese", partList.Parts()[1].Description)
	//})

}
