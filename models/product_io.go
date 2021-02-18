package models

type ProductForm struct {
    Code        string `form:"code"      valid:"Required"`
    Name        string `form:"name"      valid:"Required"`
    Number      string `form:"number"      valid:"Required"`
    Category    string `form:"category"      valid:"Required"`
    Image       string `form:"image"`
    Price       int    `form:"price"      valid:"Required"`
    PriceOnSale int    `form:"price_on_sale"     valid:"Required"`
    OnSale      bool   `form:"on_sale" valid:"Required"`
}

type ProductUpdateForm struct {
    Name        string `form:"name"       valid:"Required"`
    Image       string `form:"image"`
    Number      string `form:"number"      valid:"Required"`
    Price       int    `form:"price"     valid:"Required"`
    PriceOnSale int    `form:"price_on_sale"     valid:"Required"`
    OnSale      bool   `form:"on_sale" valid:"Required"`
}
