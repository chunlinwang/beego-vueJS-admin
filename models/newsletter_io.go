package models

type NewsletterForm struct {
    Content string `form:"content" valid:"Required"`
}
