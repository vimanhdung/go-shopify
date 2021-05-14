package goshopify

import (
	"fmt"
)

const usersBasePath = "users"

// UserService is an interface for interacting with the users
// endpoints of the Shopify API.
// See https://shopify.dev/docs/admin-api/rest/reference/plus/user
type UserService interface {
	List(interface{}) ([]User, error)
	Get(int64, interface{}) (*User, error)
}

// UserServiceOp handles communication with the user related methods of the
// Shopify API.
type UserServiceOp struct {
	client *Client
}

// User represents a Shopify user.
type User struct {
	ID                   int64    `json:"id"`
	FirstName            string   `json:"first_name"`
	Email                string   `json:"email"`
	URL                  string   `json:"url"`
	Im                   string   `json:"im"`
	ScreenName           string   `json:"screen_name"`
	Phone                string   `json:"phone"`
	LastName             string   `json:"last_name"`
	AccountOwner         bool     `json:"account_owner"`
	ReceiveAnnouncements int      `json:"receive_announcements"`
	Bio                  string   `json:"bio"`
	Permissions          []string `json:"permissions"`
	Locale               string   `json:"locale"`
	UserType             string   `json:"user_type"`
	AdminGraphqlAPIID    string   `json:"admin_graphql_api_id"`
	TfaEnabled           *bool    `json:"tfa_enabled?"`
}

// UserResource represents the result from the redirects/X.json endpoint
type UserResource struct {
	User *User `json:"user"`
}

// UsersResource represents the result from the redirects.json endpoint
type UsersResource struct {
	Users []User `json:"users"`
}

// List redirects
func (s *UserServiceOp) List(options interface{}) ([]User, error) {
	path := fmt.Sprintf("%s.json", usersBasePath)
	resource := new(UsersResource)
	err := s.client.Get(path, resource, options)
	return resource.Users, err
}

// Get individual user
func (s *UserServiceOp) Get(userID int64, options interface{}) (*User, error) {
	path := fmt.Sprintf("%s/%d.json", usersBasePath, userID)
	resource := new(UserResource)
	err := s.client.Get(path, resource, options)
	return resource.User, err
}
