package models

import "time"

//LoginForm :
type LoginForm struct {
	UserName string `json:"u" valid:"Required"`
	Password string `json:"p" valid:"Required"`
}

// RegisterForm :
type RegisterForm struct {
	ClientName       string    `json:"client_name" valid:"Required"`
	Address          string    `json:"address,omitempty"`
	PostCd           string    `json:"post_cd,omitempty"`
	TelephoneNo      string    `json:"telephone_no,omitempty"`
	EmailAddr        string    `json:"email_addr,omitempty" valid:"Email"`
	ContactPerson    string    `json:"contact_person,omitempty"`
	ClientType       string    `json:"client_type,omitempty"`
	JoiningDate      time.Time `json:"joining_date,omitempty"`
	StartBillingDate time.Time `json:"start_billing_date,omitempty"`
	ExpiryDate       time.Time `json:"expiry_date,omitempty"`
	CreatedBy        string    `json:"created_by" valid:"Required"`
}

// ForgotForm :
type ForgotForm struct {
	EmailAddr string `json:"email,omitempty" valid:"Required;Email"`
}

// ResetPasswd :
type ResetPasswd struct {
	TokenEmail    string `json:"token_email" valid:"Required"`
	Passwd        string `json:"p" valid:"Required"`
	ConfirmPasswd string `json:"cp" valid:"Required"`
}
