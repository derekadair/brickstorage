package parthttp

import (
	"brickstorage/domain/part"
	"brickstorage/web/views"
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

var (
	//go:embed static
	static embed.FS
)

type PartHandler struct {
	http.Handler

	list      *part.List
	partView  *views.ModelView[part.Part]
	indexView *views.IndexView
}

func NewPartHandler(service *part.List, partView *views.ModelView[part.Part], indexView *views.IndexView) (*PartHandler, error) {
	router := chi.NewRouter()
	handler := &PartHandler{
		Handler:   router,
		list:      service,
		partView:  partView,
		indexView: indexView,
	}

	staticHandler, err := newStaticHandler()
	if err != nil {
		return nil, fmt.Errorf("problem making static resources handler: %w", err)
	}

	router.Route("/api/v1", func(router chi.Router) {
		router.Method(http.MethodGet, "/", handler.index)
		//r.Method(http.MethodGet, "/books", requestlog.NewHandler(bookAPI.List, l))
		//router.Get("/", handler.index).Methods(http.MethodGet)
		router.Post("/parts", handler.add).Methods(http.MethodPost)
		router.Get("/parts", handler.search).Methods(http.MethodGet)
		router.Post("/parts/sort", handler.reOrder).Methods(http.MethodPost)
		router.Get("/parts/{ID}/edit", handler.edit).Methods(http.MethodGet)
		router.Post("/parts/{ID}/toggle", handler.toggle).Methods(http.MethodPost)
		router.Delete("/parts/{ID}", handler.delete).Methods(http.MethodDelete)
		router.Patch("/parts/{ID}", handler.rename).Methods(http.MethodPatch)
		router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", staticHandler))
	})

	return handler, nil
}

func (t *PartHandler) index(w http.ResponseWriter, _ *http.Request) {
	t.indexView.Index(w, t.list.Parts())
}

func (t *PartHandler) add(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t.list.Add(r.FormValue("description"))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (t *PartHandler) toggle(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.URL.Query().Get("ID")) // chi.Vars(r)["ID"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.partView.View(w, t.list.ToggleDone(id))
}

func (t *PartHandler) delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.URL.Query().Get("ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.list.Delete(id)
}

func (t *PartHandler) reOrder(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t.list.ReOrder(r.Form["id"])
	t.partView.List(w, t.list.Parts())
}

func (t *PartHandler) search(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("search")
	results := t.list.Search(searchTerm)
	t.partView.List(w, results)
}

func (t *PartHandler) rename(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := uuid.Parse(r.URL.Query().Get("ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newName := r.Form["name"][0]
	t.partView.View(w, t.list.Rename(id, newName))
}

func (t *PartHandler) edit(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(r.URL.Query().Get("ID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item := t.list.Get(id)
	t.partView.Edit(w, item)
}

func newStaticHandler() (http.Handler, error) {
	lol, err := fs.Sub(static, "static")
	if err != nil {
		return nil, err
	}
	return http.FileServer(http.FS(lol)), nil
}
