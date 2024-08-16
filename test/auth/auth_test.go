package auth_test

import (
	"antar-jemput/auth/routes"
	"antar-jemput/models"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate")
	}
	return db
}

func TestDeleteUser(t *testing.T) {
	e := echo.New()
	db := setupTestDB()

	// Create user first
	user := models.User{Name: "Test", Phone: "+123456789", Email: "example@test.com"}
	if err := user.SetPassword("testpass"); err != nil {
		t.Fatalf("failed to set password: %v", err)
	}
	db.Create(&user)

	req := httptest.NewRequest(http.MethodDelete, "/auth/user/1", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("db", db)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if err := routes.DeleteUser(c); err != nil {
		t.Fatalf("handler failed: %v", err)
	}

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}
}
