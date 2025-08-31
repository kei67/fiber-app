package user

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetUsers(t *testing.T) {
	app := fiber.New()
	app.Get("/users", GetUsers)

	req := httptest.NewRequest("GET", "/users", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()

	var users []User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}
}

func TestGetUser(t *testing.T) {
	app := fiber.New()
	app.Get("/users/:id", GetUser)

	req := httptest.NewRequest("GET", "/users/1", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if user.ID != 1 {
		t.Fatalf("expected user ID 1, got %d", user.ID)
	}
}

func TestCreateUser(t *testing.T) {
	app := fiber.New()
	app.Post("/users", CreateUser)

	payload := `{"id":3,"name":"鈴木","age":40}`
	req := httptest.NewRequest("POST", "/users", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if user.Name != "鈴木" || user.Age != 40 {
		t.Fatalf("unexpected user data: %+v", user)
	}
}
