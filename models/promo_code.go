package models

import (
    "app/models/mymongo"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "time"
)

type PromoCode struct {
    Id         bson.ObjectId `bson:"_id"         json:"_id,omitempty"`
    Code       string        `bson:"code"        json:"code,omitempty"`
    Type       string        `bson:"type"        json:"type,omitempty"`
    Value      int           `bson:"value"       json:"value,omitempty"`
    Active     bool          `bson:"active"      json:"active,omitempty"`
    MinConsomm int           `bson:"min_consomm" json:"min_consomm,omitempty"`
    BeginDate  time.Time     `bson:"begin_date"  json:"begin_date,omitempty"`
    EndDate    time.Time     `bson:"end_date"    json:"end_date,omitempty"`
}

func init() () {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTablePromoCode)

    index := mgo.Index{
        Key:        []string{"code"},
        Unique:     true,
        DropDups:   true,
        Background: true, // See notes.
        Sparse:     true,
    }

    _ = c.EnsureIndex(index)
}

func (p *PromoCode) FindByID(id string) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTablePromoCode)

    err = c.FindId(id).One(p)
    if err != nil {
        if err == mgo.ErrNotFound {
            code = ErrNotFound
        } else {
            code = ErrDatabase
        }
    } else {
        code = 0
    }
    return
}

func NewPromoCode(f *PromoCodeForm) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTablePromoCode)

    promoCode := PromoCode{
        Id:         bson.NewObjectId(),
        Code:       f.Code,
        Type:       f.Type,
        Value:      f.Value,
        MinConsomm: f.MinConsomm,
        BeginDate:  f.BeginDate,
        EndDate:    f.EndDate}

    if err = c.Insert(promoCode); err != nil {
        return -1, err
    }

    return 0, nil
}

func (PromoCode *PromoCode) PromoCodeUpdate(p *PromoCodeUpdateForm) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTablePromoCode)

    err = c.Update(bson.M{"_id": PromoCode.Code}, bson.M{"$set": bson.M{
        "type":       p.Type,
        "value":      p.Value,
        "active":     p.Active,
        "begin_date": time.Date(p.BeginDate.Year(), p.BeginDate.Month(), p.BeginDate.Day(), 0, 0, 0, 0, time.Local),
        "end_date":   time.Date(p.EndDate.Year(), p.EndDate.Month(), p.EndDate.Day(), 23, 59, 59, 99, time.Local)}})

    if err != nil {
        if err == mgo.ErrNotFound {
            return ErrNotFound, err
        }

        return ErrDatabase, err
    }

    return 0, nil
}

func PromoCodeList() (promoCodes []PromoCode, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTablePromoCode)

    err = c.Find(nil).All(&promoCodes)

    return promoCodes, err
}

func PromoCodeTotal() (total int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTablePromoCode)

    total, err = c.Find(nil).Count()

    return
}
