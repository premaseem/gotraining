// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/app"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/user"
)

// UserList returns all the existing users in the system.
// 200 Success, 404 Not Found, 500 Internal
func UserList(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(app.KeyValues).(*app.Values)

	u, err := user.List(ctx, v.TraceID, v.DB)
	if err != nil {
		return err
	}

	app.Respond(w, v.TraceID, u, http.StatusOK)
	return nil
}

// UserRetrieve returns the specified user from the system.
// 200 Success, 400 Bad Request, 404 Not Found, 500 Internal
func UserRetrieve(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(app.KeyValues).(*app.Values)

	u, err := user.Retrieve(ctx, v.TraceID, v.DB, params["id"])
	if err != nil {
		return err
	}

	app.Respond(w, v.TraceID, u, http.StatusOK)
	return nil
}

// UserCreate inserts a new user into the system.
// 200 OK, 400 Bad Request, 500 Internal
func UserCreate(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(app.KeyValues).(*app.Values)

	var u user.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return err
	}

	if err := user.Create(ctx, v.TraceID, v.DB, &u); err != nil {
		return err
	}

	app.Respond(w, v.TraceID, u, http.StatusCreated)
	return nil
}

// UserUpdate updates the specified user in the system.
// 200 Success, 400 Bad Request, 500 Internal
func UserUpdate(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(app.KeyValues).(*app.Values)

	var u user.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return err
	}

	if err := user.Update(ctx, v.TraceID, v.DB, params["id"], &u); err != nil {
		return err
	}

	app.Respond(w, v.TraceID, nil, http.StatusNoContent)
	return nil
}

// UserDelete removed the specified user from the system.
// 200 Success, 400 Bad Request, 500 Internal
func UserDelete(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
	v := ctx.Value(app.KeyValues).(*app.Values)

	u, err := user.Retrieve(ctx, v.TraceID, v.DB, params["id"])
	if err != nil {
		return err
	}

	if err := user.Delete(ctx, v.TraceID, v.DB, params["id"]); err != nil {
		return err
	}

	app.Respond(w, v.TraceID, u, http.StatusOK)
	return nil
}
