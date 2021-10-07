package types

import (
	"github.com/antihax/optional"
	"time"
)

type CreateUserRequest struct {
	// The user name of the user.
	Username string `json:"username,omitempty"`
	// The identifier of the organization the user belongs to.
	OrgUUID string `json:"org_UUID,omitempty"`
	// The identifier of the user.
	UUID      string                     `json:"UUID,omitempty"`
	Name      *OrgShortNameUserName      `json:"name,omitempty"`
	Authority *OrgShortNameUserAuthority `json:"authority,omitempty"`
}

type CreateUserResponse struct {
	// Success description
	Message string `json:"message,omitempty"`
	Created *User  `json:"created,omitempty"`
}

type ListUsersOpts struct {
	Page optional.Int32
}

type ListUsersResponse struct {
	// Total user records found
	TotalCount int32 `json:"totalCount,omitempty"`
	// Number of user records in a page
	ItemsPerPage int32 `json:"itemsPerPage,omitempty"`
	// Total number of pages
	PageCount int32 `json:"pageCount,omitempty"`
	// Current page
	CurrentPage int32 `json:"currentPage,omitempty"`
	// Previous page
	PrevPage int32 `json:"prevPage,omitempty"`
	// Next page
	NextPage int32           `json:"nextPage,omitempty"`
	Users    *[]UserNoSecret `json:"users,omitempty"`
}

type OrgShortNameUserAuthority struct {
	ActiveRoles []string `json:"active_roles,omitempty"`
}

type OrgShortNameUserName struct {
	// The first name of the user.
	First string `json:"first,omitempty"`
	// The last name of the user.
	Last string `json:"last,omitempty"`
	// The middle name of the user, if any.
	Middle string `json:"middle,omitempty"`
	// The suffix of the user, if any.
	Suffix string `json:"suffix,omitempty"`
}

type UpdateUserOpts struct {
	Active            optional.Bool
	NewUsername       optional.String
	OrgShortname      optional.String
	NameFirst         optional.String
	NameLast          optional.String
	NameMiddle        optional.String
	NameSuffix        optional.String
	ActiveRolesAdd    optional.String
	ActiveRolesRemove optional.String
}

type UpdatedUserResponse struct {
	// Success description
	Message string        `json:"message,omitempty"`
	Updated *UserNoSecret `json:"updated,omitempty"`
}

type User struct {
	// The user name of the user.
	Username string `json:"username,omitempty"`
	// The identifier of the organization the user belongs to.
	OrgUUID string `json:"org_UUID,omitempty"`
	// The identifier of the user.
	UUID string `json:"UUID,omitempty"`
	// The user is an active user of the organization.
	Active bool                  `json:"active,omitempty"`
	Name   *OrgShortNameUserName `json:"name,omitempty"`
	// The API key of the user.
	Secret    string                     `json:"secret,omitempty"`
	Authority *OrgShortNameUserAuthority `json:"authority,omitempty"`
	Time      *UserTime                  `json:"time,omitempty"`
}

type UserNoSecret struct {
	// The user name of the user.
	Username string `json:"username,omitempty"`
	// The identifier of the organization the user belongs to.
	OrgUUID string `json:"org_UUID,omitempty"`
	// The identifier of the user.
	UUID string `json:"UUID,omitempty"`
	// The user is an active user of the organization.
	Active    bool                       `json:"active,omitempty"`
	Name      *OrgShortNameUserName      `json:"name,omitempty"`
	Authority *OrgShortNameUserAuthority `json:"authority,omitempty"`
	Time      *UserTime                  `json:"time,omitempty"`
}

type UserTime struct {
	// The time the user was created.
	Created time.Time `json:"created,omitempty"`
	// The last time the user was modified.
	Modified time.Time `json:"modified,omitempty"`
}
