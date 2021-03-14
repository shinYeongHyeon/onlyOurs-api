// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/shinYeongHyeon/onlyOurs-api/ent/schema"
	"github.com/shinYeongHyeon/onlyOurs-api/ent/users"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	usersFields := schema.Users{}.Fields()
	_ = usersFields
	// usersDescEmail is the schema descriptor for email field.
	usersDescEmail := usersFields[0].Descriptor()
	// users.DefaultEmail holds the default value on creation for the email field.
	users.DefaultEmail = usersDescEmail.Default.(string)
	// usersDescName is the schema descriptor for name field.
	usersDescName := usersFields[1].Descriptor()
	// users.DefaultName holds the default value on creation for the name field.
	users.DefaultName = usersDescName.Default.(string)
	// usersDescPassword is the schema descriptor for password field.
	usersDescPassword := usersFields[2].Descriptor()
	// users.DefaultPassword holds the default value on creation for the password field.
	users.DefaultPassword = usersDescPassword.Default.(string)
}
