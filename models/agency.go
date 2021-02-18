package models

import (
    "app/models/mymongo"
    "github.com/astaxie/beego"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type Agency struct {
    ID        bson.ObjectId `bson:"_id"       json:"_id,omitempty"`
    Address   string        `bson:"address"   json:"address,omitempty"`
    City      string        `bson:"city"      json:"city,omitempty"`
    ZipCode   string        `bson:"zip_code"  json:"zip_code,omitempty"`
    Phone     string        `bson:"phone"     json:"phone,omitempty"`
    Email     string        `bson:"email"     json:"email,omitempty"`
    Schedules []Schedule    `bson:"schedules" json:"schedules,omitempty"`
    IsOpen    bool          `bson:"is_open"   json:"is_open,true"`
}

type Schedule struct {
    Day       int       `json:"day,omitempty"`
    BeginHour time.Time `json:"begin_hour,omitempty"`
    EndHour   time.Time `json:"end_hour,omitempty"`
}

var days = [...]string{
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
    "Sunday",
}


func InitAgency() (code int, err error) {

    agency := Agency{
        ID:        bson.NewObjectId(),
        Address:   "108 rue du vieux pont de Sévres",
        City:      "Boulogne-Billancourt",
        ZipCode:   "92100",
        Phone:     "01 58 17 04 04",
        Email:     "",
        IsOpen:    true}

    if count, _ := agency.Count(); count == 0 {
        return agency.Insert()
    }

    return 0, nil
}

// Insert insert a document to collection.
func (a *Agency) Insert() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableAgency)

    if count, err := c.FindId(a.ID).Count(); count == 0 || err != nil {
        err = c.Insert(a)
        if err != nil {
            if mgo.IsDup(err) {
                code = ErrDupRows
            } else {
                code = ErrDatabase
            }
        } else {
            code = 0
        }
    } else {
        return -1, nil
    }

    return
}

func (a *Agency) Count() (count int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableAgency)

    count, err = c.FindId(nil).Count()

    return
}

func (a *Agency) FirstAgency() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableAgency)

    err = c.Find(nil).One(&a)

    return 0, err
}

func (a *Agency) AgencyUpdate(f *AgencyForm) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableAgency)

    beego.Debug(f.Schedules)
    err = c.Update(bson.M{"_id": a.ID}, bson.M{"$set": bson.M{
        "address":   f.Address,
        "city":      f.City,
        "zip_code":  f.ZipCode,
        "phone":     f.Phone,
        "email":     f.Email,
        "schedules": f.Schedules,
        "is_open":   f.IsOpen}})

    if err != nil {
        if err == mgo.ErrNotFound {
            return ErrNotFound, err
        }

        return ErrDatabase, err
    }

    return 0, nil
}
