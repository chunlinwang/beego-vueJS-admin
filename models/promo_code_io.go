package models

import (
    "time"
)

type PromoCodeForm struct {
    Code       string    `form:"code"      valid:"Required"`
    Type       string    `form:"type"       valid:"Required"`
    Value      int       `form:"value"      valid:"Required"`
    MinConsomm int       `form:"min_consomm" `
    Active     bool      `form:"active"     valid:"Required"`
    BeginDate  time.Time `form:"begin_date" valid:"Required"`
    EndDate    time.Time `form:"end_date"   valid:"Required"`
}

type PromoCodeUpdateForm struct {
    Type       string    `form:"type"       valid:"Required"`
    Value      int       `form:"value"      valid:"Required"`
    Active     bool      `form:"active"     valid:"Required"`
    MinConsomm int       `form:"min_consomm" `
    BeginDate  time.Time `form:"begin_date" valid:"Required"`
    EndDate    time.Time `form:"end_date"   valid:"Required"`
}
