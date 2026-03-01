package util

import (
	"context"
	"fmt"
	"housebook/consts"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func TestPathID(t *testing.T) {
	wanted := 1

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/{id}", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(consts.IDPathValue, strconv.Itoa(wanted))
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	func(w http.ResponseWriter, r *http.Request) {
		id, err := PathID(consts.IDPathValue, r)
		if err != nil {
			t.Fatal(err)
		}
		if id != int64(wanted) {
			t.Errorf("PathID returned %d; wanted %d", id, wanted)
		}
	}(w, r)
}

func TestPathUUID(t *testing.T) {
	var wanted pgtype.UUID
	wanted.Scan("00000000-0000-0000-0000-000000000000")

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/{uuid}", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(consts.UUIDPathValue, wanted.String())
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	func(w http.ResponseWriter, r *http.Request) {
		uuid, err := PathUUID(consts.UUIDPathValue, r)
		if err != nil {
			t.Fatal(err)
		}
		if uuid != wanted {
			t.Errorf("PathUUID returned %s; wanted %s", uuid.String(), wanted.String())
		}
	}(w, r)
}

func TestGetSearchPagination(t *testing.T) {
	t.Parallel()
	tests := []struct {
		limit, offset             string
		wantedLimit, wantedOffset int
	}{
		{"1", "1", 1, 1},
		{"5", "5", 5, 5},
		{"", "1", 10, 1},
		{"1", "", 1, 0},
		{"", "", 10, 0},
	}
	for _, c := range tests {
		name := fmt.Sprintf("%s,%s", c.limit, c.offset)
		t.Run(name, func(t *testing.T) {
			u, err := url.ParseQuery(fmt.Sprintf("limit=%s&offset=%s", c.limit, c.offset))
			if err != nil {
				t.Fatal(err)
			}
			limit, offset, err := GetSearchPagination(u)
			if err != nil {
				t.Fatal(err)
			}
			if limit != int32(c.wantedLimit) || offset != int32(c.wantedOffset) {
				t.Errorf(
					"GetSearchPagination returned %d, %d; wanted %d, %d",
					limit, offset,
					c.wantedLimit, c.wantedOffset,
				)
			}
		})
	}
}
