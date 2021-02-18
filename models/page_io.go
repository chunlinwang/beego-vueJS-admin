package models

type PageForm struct {
    Title          string `form:"title"   valid:"Required"`
    Code           string `form:"code"    valid:"Required"`
    Content        string `form:"content" valid:"Required"`
    SeoTitle       string `form:"seo_title"`
    SeoDescription string `form:"seo_description"`
}

type PageUpdateForm struct {
    Title          string `form:"title"   valid:"Required"`
    Content        string `form:"content" valid:"Required"`
    SeoTitle       string `form:"seo_title"`
    SeoDescription string `form:"seo_description"`
}
