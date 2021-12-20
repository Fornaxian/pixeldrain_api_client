package pixelapi

import (
	"net/url"
	"strconv"
	"time"

	"github.com/gocql/gocql"
)

// UserInfo contains information about the logged in user
type UserInfo struct {
	Username            string           `json:"username"`
	Email               string           `json:"email"`
	Subscription        SubscriptionType `json:"subscription"`
	StorageSpaceUsed    int64            `json:"storage_space_used"`
	IsAdmin             bool             `json:"is_admin"`
	BalanceMicroEUR     int64            `json:"balance_micro_eur"`
	Hotlinking          bool             `json:"hotlinking_enabled"`
	MonthlyTransferCap  int64            `json:"monthly_transfer_cap"`
	MonthlyTransferUsed int64            `json:"monthly_transfer_used"`
}

// UserSession is one user session
type UserSession struct {
	AuthKey      gocql.UUID `json:"auth_key"`
	CreationIP   string     `json:"creation_ip_address"`
	UserAgent    string     `json:"user_agent"`
	AppName      string     `json:"app_name"`
	CreationTime time.Time  `json:"creation_time"`
	LastUsedTime time.Time  `json:"last_used_time"`
}

// UserRegister registers a new user on the Pixeldrain server. username and
// password are always required. email is optional, but without it you will not
// be able to reset your password in case you forget it. captcha depends on
// whether reCaptcha is enabled on the Pixeldrain server, this can be checked
// through the GetRecaptcha function.
//
// The register API can return multiple errors, which will be stored in the
// Errors array. Check for len(Errors) == 0 to see if an error occurred. If err
// != nil it means a connection error occurred
func (p *PixelAPI) UserRegister(username, email, password, captcha string) (err error) {
	return p.form(
		"POST", "user/register",
		url.Values{
			"username":           {username},
			"email":              {email},
			"password":           {password},
			"recaptcha_response": {captcha},
		},
		nil,
	)
}

// PostUserLogin logs a user in with the provided credentials. The response will
// contain the returned API key. The app name is saved in the database and can
// be found on the user's API keys page.
func (p *PixelAPI) PostUserLogin(username, password, app string) (resp UserSession, err error) {
	return resp, p.form(
		"POST", "user/login",
		url.Values{
			"username": {username},
			"password": {password},
			"app_name": {app},
		},
		&resp,
	)
}

// GetUser returns information about the logged in user. Requires an API key
func (p *PixelAPI) GetUser() (resp UserInfo, err error) {
	return resp, p.jsonRequest("GET", "user", &resp)
}

// PostUserSession creates a new user sessions
func (p *PixelAPI) PostUserSession(app string) (resp UserSession, err error) {
	return resp, p.form("POST", "user/session", url.Values{"app": {app}}, &resp)
}

// GetUserSession lists all active user sessions
func (p *PixelAPI) GetUserSession() (resp []UserSession, err error) {
	return resp, p.jsonRequest("GET", "user/session", &resp)
}

// DeleteUserSession destroys an API key so it can no longer be used to perform
// actions
func (p *PixelAPI) DeleteUserSession(key string) (err error) {
	return p.jsonRequest("DELETE", "user/session", nil)
}

// FileInfoSlice a collection of files which belong to a user
type FileInfoSlice struct {
	Files []FileInfo `json:"files"`
}

// GetUserFiles gets files uploaded by a user
func (p *PixelAPI) GetUserFiles() (resp FileInfoSlice, err error) {
	return resp, p.jsonRequest("GET", "user/files", &resp)
}

// ListInfoSlice is a collection of lists which belong to a user
type ListInfoSlice struct {
	Lists []ListInfo `json:"lists"`
}

// GetUserLists gets lists created by a user
func (p *PixelAPI) GetUserLists() (resp ListInfoSlice, err error) {
	return resp, p.jsonRequest("GET", "user/lists", &resp)
}

type UserTransaction struct {
	Time               time.Time `json:"time"`
	NewBalance         int64     `json:"new_balance"`
	DepositAmount      int64     `json:"deposit_amount"`
	SubscriptionCharge int64     `json:"subscription_charge"`
	StorageCharge      int64     `json:"storage_charge"`
	StorageUsed        int64     `json:"storage_used"`
	BandwidthCharge    int64     `json:"bandwidth_charge"`
	BandwidthUsed      int64     `json:"bandwidth_used"`
}

func (p *PixelAPI) GetUserTransactions() (resp []UserTransaction, err error) {
	return resp, p.jsonRequest("GET", "user/transactions", &resp)
}

// PutUserPassword changes the user's password
func (p *PixelAPI) PutUserPassword(oldPW, newPW string) (err error) {
	return p.form(
		"PUT", "user/password",
		url.Values{"old_password": {oldPW}, "new_password": {newPW}},
		nil,
	)
}

// PutUserEmailReset starts the e-mail change process. An email will be sent to
// the new address to verify that it's real. Once the link in the e-mail is
// clicked the key it contains can be sent to the API with UserEmailResetConfirm
// and the change will be applied
func (p *PixelAPI) PutUserEmailReset(email string, delete bool) (err error) {
	return p.form(
		"PUT", "user/email_reset",
		url.Values{"new_email": {email}, "delete": {strconv.FormatBool(delete)}},
		nil,
	)
}

// PutUserEmailResetConfirm finishes process of changing a user's e-mail address
func (p *PixelAPI) PutUserEmailResetConfirm(key string) (err error) {
	return p.form(
		"PUT", "user/email_reset_confirm",
		url.Values{"key": {key}},
		nil,
	)
}

// PutUserPasswordReset starts the password reset process. An email will be sent
// the user to verify that it really wanted to reset the password. Once the link
// in the e-mail is clicked the key it contains can be sent to the API with
// UserPasswordResetConfirm and a new password can be set
func (p *PixelAPI) PutUserPasswordReset(email string, recaptchaResponse string) (err error) {
	return p.form(
		"PUT", "user/password_reset",
		url.Values{"email": {email}, "recaptcha_response": {recaptchaResponse}},
		nil,
	)
}

// PutUserPasswordResetConfirm finishes process of resetting a user's password.
// If the key is valid the new_password parameter will be saved as the new
// password
func (p *PixelAPI) PutUserPasswordResetConfirm(key string, newPassword string) (err error) {
	return p.form(
		"PUT", "user/password_reset_confirm",
		url.Values{"key": {key}, "new_password": {newPassword}},
		nil,
	)
}

// PutUserUsername changes the user's username.
func (p *PixelAPI) PutUserUsername(username string) (err error) {
	return p.form(
		"PUT", "user/username",
		url.Values{"new_username": {username}},
		nil,
	)
}
