package handler

import (
	"net/http/httptest"
	"testing"
	"time"
	"reflect"
	"encoding/json"
	"log"
)

func TestOmikujiHandler(t *testing.T) {
	type fields struct {
		month time.Month
		day int
	}

	tests := []struct {
		name   string
		fields fields
		want string
	}{
		{name:"1/1", fields:fields{month:1, day:1}, want:"大吉"},
		{name:"1/2", fields:fields{month:1, day:2}, want:"大吉"},
		{name:"1/3", fields:fields{month:1, day:3}, want:"大吉"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/", nil)
			loc, _ := time.LoadLocation("Asia/Tokyo")
			date := time.Date(2018, tt.fields.month, tt.fields.day, 0, 0, 0, 0, loc)
			OmikujiHandler{TimeProvider: ArbitraryTimeProvider{time: date}}.ServeHTTP(w, r)
			rw := w.Result()
			defer rw.Body.Close()
			decoder := json.NewDecoder(rw.Body)
			var result result
			if err := decoder.Decode(&result); err != nil {
				log.Fatal(err)
			}
			if !reflect.DeepEqual(result.Data, tt.want) {
				t.Errorf("omikujiHandler = %v, want %v", result.Data, tt.want)
			}
		})
	}
}
