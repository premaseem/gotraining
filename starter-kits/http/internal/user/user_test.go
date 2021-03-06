// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/ardanlabs/gotraining/starter-kits/http/internal/platform/middleware"
	"github.com/ardanlabs/gotraining/starter-kits/http/internal/user"
)

const (

	// Succeed is the Unicode codepoint for a check mark.
	Succeed = "\u2713"

	// Failed is the Unicode codepoint for an X mark.
	Failed = "\u2717"
)

// TestUsers validates a user can be created, retrieved and
// then removed from the system.
func TestUsers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := middleware.NewMGOSession()
	defer db.Close()

	now := time.Now()
	traceID := "traceid"

	u := user.User{
		UserType:     1,
		FirstName:    "Bill",
		LastName:     "Kennedy",
		Email:        "bill@ardanstugios.com",
		Company:      "Ardan Labs",
		DateModified: &now,
		DateCreated:  &now,
		Addresses: []user.Address{
			{
				Type:         1,
				LineOne:      "12973 SW 112th ST",
				LineTwo:      "Suite 153",
				City:         "Miami",
				State:        "FL",
				Zipcode:      "33172",
				Phone:        "305-527-3353",
				DateModified: &now,
				DateCreated:  &now,
			},
		},
	}

	t.Log("Given the need to add a new user, retrieve and remove that user from the system.")
	{
		if err := u.Validate(); err != nil {
			t.Fatal("\tShould be able to validate the user data.", Failed)
		}
		t.Log("\tShould be able to validate the user data.", Succeed)

		if err := user.Create(ctx, traceID, db, &u); err != nil {
			t.Fatal("\tShould be able to create a user in the system.", Failed)
		}
		t.Log("\tShould be able to create a user in the system.", Succeed)

		if u.UserID == "" {
			t.Fatal("\tShould have an UserID for the user.", Failed)
		}
		t.Log("\tShould have an UserID for the user.", Succeed)

		ur, err := user.Retrieve(ctx, traceID, db, u.UserID)
		if err != nil {
			t.Fatal("\tShould be able to retrieve the user back from the system.", Failed)
		}
		t.Log("\tShould be able to retrieve the user back from the system.", Succeed)

		if ur == nil || u.UserID != ur.UserID {
			t.Fatal("\tShould have a match between the created user and the one retrieved.", Failed)
		}
		t.Log("\tShould have a match between the created user and the one retrieved.", Succeed)

		if err := user.Delete(ctx, traceID, db, u.UserID); err != nil {
			t.Fatal("\tShould be able to remove the user from the system.", Failed)
		}
		t.Log("\tShould be able to remove the user from the system", Succeed)

		if _, err := user.Retrieve(ctx, traceID, db, u.UserID); err == nil {
			t.Fatal("\tShould NOT be able to retrieve the user back from the system.", Failed)
		}
		t.Log("\tShould NOT be able to retrieve the user back from the system.", Succeed)
	}
}
