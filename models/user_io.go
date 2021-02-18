package models

// RegisterForm definition.
type RegisterForm struct {
    Phone      string `form:"phone"       valid:"Required"`
    Email      string `form:"email"       valid:"Required"`
    Password   string `form:"password"    valid:"Required"`
    LastName   string `form:"last_name"   valid:"Required"`
    FirstName  string `form:"first_name"  valid:"Required"`
    Civility   string `form:"civility"    valid:"Required"`
    Address    string `form:"address"     valid:"Required"`
    SupAddress string `form:"sup_address" `
    ZipCode    string `form:"zip_code"    valid:"Required"`
    City       string `form:"city"        valid:"Required"`
    Admin      bool    `form:"admin"`
    Roles      []string `form:"roles"`
}

type UpdateUserForm struct {
    Phone      string `form:"phone"       valid:"Required"`
    LastName   string `form:"last_name"   valid:"Required"`
    FirstName  string `form:"first_name"  valid:"Required"`
    Civility   string `form:"civility"    valid:"Required"`
    Address    string `form:"address"     valid:"Required"`
    SupAddress string `form:"sup_address" `
    ZipCode    string `form:"zip_code"    valid:"Required"`
    City       string `form:"city"        valid:"Required"`
    Admin      bool    `form:"admin"`
    Roles      []string `form:"roles"`
}

// LoginForm definition.
type LoginForm struct {
    Email    string `form:"username" valid:"Required"`
    Password string `form:"password" valid:"Required"`
}

// LoginInfo definition.
type LoginInfo struct {
    Code     int   `json:"code"`
    UserInfo *User `json:"user"`
}

// LogoutForm defintion.
type LogoutForm struct {
    Email string `form:"email" valid:"Required;Email"`
}

// PasswdForm definition.
type PasswdForm struct {
    Email   string `form:"email"        valid:"Required;Email"`
    OldPass string `form:"old_password" valid:"Required"`
    NewPass string `form:"new_password" valid:"Required"`
}

// UploadsForm definition.
type UploadsForm struct {
    Email string `form:"email" valid:"Required;Email"`
}

type RequestPassWdForm struct {
    Email string `form:"email" valid:"Required;Email"`
}

// RoleAuthInfo definition.
type RoleAuthInfo struct {
    Token string `json:"myAuthToken"`
}

// AdminRoleAuthInfo definition.
type AdminRoleAuthInfo struct {
    Token   string   `json:"myAuthToken"`
    IsAdmin bool     `json:"isAdmin"`
    Roles   []string `json:"roles"`
}
