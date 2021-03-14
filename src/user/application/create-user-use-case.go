package user_application_create_user_use_case

import (
	"github.com/shinYeongHyeon/onlyOurs-api/core"
	"net/http"
)

type user struct {
	A int
	Name string
}

// Exec: user_application_create_user_use_case.exec()
func Exec(w http.ResponseWriter, r *http.Request) {
	core.WriteJSON(w, user {
		1,
		"신영현",
	})
}
