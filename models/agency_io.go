package models

type AgencyForm struct {
    Address   string     `form:"address"`
    City      string     `form:"city"`
    ZipCode   string     `form:"zip_code"`
    Phone     string     `form:"phone"`
    Email     string     `form:"email"`
    Schedules []Schedule `form:"schedules"`
    IsOpen    bool       `form:"is_open"`
}
