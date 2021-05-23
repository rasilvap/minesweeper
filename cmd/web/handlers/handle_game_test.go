package handlers_test

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"minesweeper-API/cmd/web/handlers"
	"minesweeper-API/cmd/web/handlers/mocks"
	"minesweeper-API/models/dto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createGETTest(t *testing.T, url string) (*http.Request, *httptest.ResponseRecorder) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal("failed to build request")
	}

	return req, httptest.NewRecorder()
}

func Test_Get(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		e := new(mocks.Game)
		e.On("Get", 123).Return(&dto.GameResponse{
			Rows:       3,
			Columns:    3,
			MineAmount: 1,
		}, nil)
		game := handlers.NewGame(e)
		r := mux.NewRouter()
		handlers.SetupRoutes(r, game)

		req, rr := createGETTest(t, "/v1/games/123")

		// act
		r.ServeHTTP(rr, req)

		assert.Equal(t, 200, rr.Code)
		assert.JSONEq(t, `{"columns":3, "mineAmount":1, "rows":3}`, rr.Body.String())
	})

	t.Run("Error", func(t *testing.T) {
		e := new(mocks.Game)
		game := handlers.NewGame(e)
		r := mux.NewRouter()
		handlers.SetupRoutes(r, game)

		req, rr := createGETTest(t, "/v1/games/xyz")

		// act
		r.ServeHTTP(rr, req)

		assert.Equal(t, 404, rr.Code)
	})
}
