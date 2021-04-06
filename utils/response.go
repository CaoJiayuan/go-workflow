package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, data interface{}) error {
	var (
		b []byte
		e error
	)
	if s, ok := data.(string); ok {
		b = []byte(s)
	} else if bb, ok := data.([]byte); ok {
		b = bb
	} else {
		b, e = json.Marshal(data)
		if e != nil {
			return e
		}
	}

	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(b)
	return err
}

func PageJSONString(datas interface{}, count int64, pageIndex, pageSize int) (string, error) {
	data, err := PageJson(datas, count, pageIndex, pageSize)

	return string(data), err
}

func PageJson(datas interface{}, count int64, pageIndex, pageSize int) ([]byte, error) {
	result := map[string]interface{}{
		"rows":     datas,
		"pageSize": pageSize,
		"total":    count,
		"page":     pageIndex,
	}

	return json.Marshal(result)
}
