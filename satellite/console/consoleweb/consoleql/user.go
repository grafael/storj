// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package consoleql

import (
	"github.com/graphql-go/graphql"

	"storj.io/storj/satellite/console"
)

const (
	// UserType is a graphql type for user
	UserType = "user"
	// UserInputType is a graphql type for user input
	UserInputType = "userInput"
	// FieldID is a field name for id
	FieldID = "id"
	// FieldEmail is a field name for email
	FieldEmail = "email"
	// FieldPassword is a field name for password
	FieldPassword = "password"
	// FieldFullName is a field name for "first name"
	FieldFullName = "fullName"
	// FieldShortName is a field name for "last name"
	FieldShortName = "shortName"
	// FieldCreatedAt is a field name for created at timestamp
	FieldCreatedAt = "createdAt"
	// FieldPartnerID is a field name for partnerID
	FieldPartnerID = "partnerId"
)

// base graphql config for user
func baseUserConfig() graphql.ObjectConfig {
	return graphql.ObjectConfig{
		Name: UserType,
		Fields: graphql.Fields{
			FieldID: &graphql.Field{
				Type: graphql.String,
			},
			FieldEmail: &graphql.Field{
				Type: graphql.String,
			},
			FieldFullName: &graphql.Field{
				Type: graphql.String,
			},
			FieldShortName: &graphql.Field{
				Type: graphql.String,
			},
			FieldCreatedAt: &graphql.Field{
				Type: graphql.DateTime,
			},
			FieldPartnerID: &graphql.Field{
				Type: graphql.String,
			},
		},
	}
}

// graphqlUser creates *graphql.Object type representation of satellite.User
// TODO: simplify
func graphqlUser() *graphql.Object {
	return graphql.NewObject(baseUserConfig())
}

// graphqlUserInput creates graphql.InputObject type needed to register/update satellite.User
func graphqlUserInput() *graphql.InputObject {
	return graphql.NewInputObject(graphql.InputObjectConfig{
		Name: UserInputType,
		Fields: graphql.InputObjectConfigFieldMap{
			FieldEmail: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			FieldFullName: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			FieldShortName: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			FieldPassword: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			FieldPartnerID: &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	})
}

func fromMapCreateUser(args map[string]interface{}) (user console.CreateUser) {
	user.Email, _ = args[FieldEmail].(string)
	user.FullName, _ = args[FieldFullName].(string)
	user.ShortName, _ = args[FieldShortName].(string)
	user.Password, _ = args[FieldPassword].(string)
	user.PartnerID, _ = args[FieldPartnerID].(string)
	return
}

// fillUserInfo fills satellite.UserInfo from satellite.User and input args
func fillUserInfo(user *console.User, args map[string]interface{}) console.UserInfo {
	info := console.UserInfo{
		FullName:  user.FullName,
		ShortName: user.ShortName,
	}

	for fieldName, fieldValue := range args {
		value, ok := fieldValue.(string)
		if !ok {
			continue
		}

		switch fieldName {
		case FieldFullName:
			info.FullName = value
			user.FullName = value
		case FieldShortName:
			info.ShortName = value
			user.ShortName = value
		}
	}

	return info
}
