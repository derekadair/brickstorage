package bookweb

import "net/http"

// Read godoc
//
//	@summary		Read book form
//	@description	Read book form
//	@tags			books
//	@accept			json
//	@produce		json
//	@success		200	{array}		DTO
//	@failure		500	{object}	err.Error
//	@router			/ [get]
func Read(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Book web form"))
}
