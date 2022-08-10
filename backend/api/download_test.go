package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jak103/usu-gdsf/db"
	"github.com/jak103/usu-gdsf/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	_db, _ = db.NewDatabaseFromEnv()

	download0 = models.Download{
		UserId:       "000",
		GameId:       "62dadce12fcdb47399118408",
		CreationDate: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	download1 = models.Download{
		UserId:       "001",
		GameId:       "62dadce12fcdb47399118408",
		CreationDate: time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC),
	}
)

func TestGetAllDownloads(t *testing.T) {
	e := echo.New()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/downloads", nil)
	c := e.NewContext(request, recorder)

	if assert.NoError(t, getAllDownloads(c)) {
		assert.Equal(t, http.StatusOK, recorder.Code)
	}
}

func TestAddGame(t *testing.T) {
	_, err := _db.AddDownload(download0)
	assert.Nil(t, err)
}