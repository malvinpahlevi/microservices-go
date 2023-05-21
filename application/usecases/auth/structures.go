// Package auth provides the use case for authentication
package auth

import "time"

// LoginUser is a struct that contains the request body for the login user
type LoginUser struct {
	Email    string
	Password string
}

// DataUserAuthenticated is a struct that contains the data for the authenticated user
type DataUserAuthenticated struct {
	UserName  string `json:"username" example:"UserName" gorm:"unique"`
	Email     string `json:"email" example:"some@mail.com" gorm:"unique"`
	FirstName string `json:"firstname" example:"John"`
	LastName  string `json:"lastname" example:"Doe"`
	Status    bool   `json:"status" example:"1"`
	Role      string `json:"role" example:"admin"`
	ID        int    `json:"id" example:"123"`
}

// DataSecurityAuthenticated is a struct that contains the security data for the authenticated user
type DataSecurityAuthenticated struct {
	JWTAccessToken            string    `json:"jwt_access_token" example:"SomeAccessToken"`
	JWTRefreshToken           string    `json:"jwt_refresh_token" example:"SomeRefreshToken"`
	ExpirationAccessDateTime  time.Time `json:"expiration_access_datetime" example:"2023-02-02T21:03:53.196419-06:00"`
	ExpirationRefreshDateTime time.Time `json:"expiration_refresh_datetime" example:"2023-02-03T06:53:53.196419-06:00"`
}

// SecurityAuthenticatedUser is a struct that contains the data for the authenticated user
type SecurityAuthenticatedUser struct {
	Data     DataUserAuthenticated     `json:"data"`
	Security DataSecurityAuthenticated `json:"security"`
}
