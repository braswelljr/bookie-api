package middleware

import "strings"

const (
	RoleSuperAdmin = "superadmin"
	RoleAdmin      = "admin"
	RoleUser       = "user"
)

// HasRole - is a function that checks if a user has a role.
//
//	@param ctx - context.Context
//	@param role - string
//	@return bool
func (data *DataI) HasRole(roles ...string) bool {
	// check and loop through the data roles
	for _, dataRole := range data.Roles {
		// loop through the roles
		for _, role := range roles {
			// check if the role is valid
			if strings.EqualFold(dataRole, role) {
				return true
			}
		}
	}

	return false
}
