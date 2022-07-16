package rest

import (
	"io"
	"net/http"
	"strconv"
)

func PostData(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No body set"))
		return
	}

	if len(d) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Empty body"))
		return
	}

	w.Write([]byte("I got message:\n" + string(d)))
	w.WriteHeader(http.StatusOK)
}

func PostHeader(w http.ResponseWriter, r *http.Request) {
	header1 := r.Header.Get("a")
	header2 := r.Header.Get("b")

	if len(header1) == 0 || len(header2) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No headers found"))
		return
	}

	val1, err := strconv.Atoi(header1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid val for header a"))
		return
	}

	val2, err := strconv.Atoi(header2)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid val for header b"))
		return
	}

	result := val1 + val2

	w.Header().Set("a+b", strconv.Itoa(result))
	w.WriteHeader(http.StatusOK)
}
