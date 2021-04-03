// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/shinYeongHyeon/onlyOurs-api/ent/schema"
	"github.com/shinYeongHyeon/onlyOurs-api/ent/users"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	usersFields := schema.Users{}.Fields()
	_ = usersFields
	// usersDescUserID is the schema descriptor for user_id field.
	usersDescUserID := usersFields[1].Descriptor()
	// users.DefaultUserID holds the default value on creation for the user_id field.
	users.DefaultUserID = usersDescUserID.Default.(string)
	// usersDescName is the schema descriptor for name field.
	usersDescName := usersFields[2].Descriptor()
	// users.DefaultName holds the default value on creation for the name field.
	users.DefaultName = usersDescName.Default.(string)
	// usersDescPassword is the schema descriptor for password field.
	usersDescPassword := usersFields[3].Descriptor()
	// users.DefaultPassword holds the default value on creation for the password field.
	users.DefaultPassword = usersDescPassword.Default.(string)
	// usersDescCreatedAt is the schema descriptor for created_at field.
	usersDescCreatedAt := usersFields[4].Descriptor()
	// users.DefaultCreatedAt holds the default value on creation for the created_at field.
	users.DefaultCreatedAt = usersDescCreatedAt.Default.(func() time.Time)
}
