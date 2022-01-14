package app

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	integer int8
	string  string
	float   float32
}

func getData(writer http.ResponseWriter, _ *http.Request) {
	data := []Data{
		{2, "dummy text 0", 1.2},
		{4, "dummy text 1", 2.4},
		{8, "dummy text 2", 4.8},
	}

	writer.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(data)

	if err != nil {
		return
	}
}
