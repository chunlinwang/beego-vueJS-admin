package models

import (
    "app/models/mymongo"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type DeliveryCity struct {
    ID      bson.ObjectId `bson:"_id"       json:"_id,omitempty"`
    City    string        `bson:"city"      json:"city,omitempty"`
    ZipCode string        `bson:"zip_code"  json:"zip_code,omitempty"`
    Active  bool          `bson:"active"   json:"active,true"`
}

func (d *DeliveryCity) FindByID(id string) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableDeliveryCity)
    err = c.FindId(bson.ObjectIdHex(id)).One(&d)

    if err != nil {
        if mgo.IsDup(err) {
            code = ErrDupRows
        } else {
            code = ErrDatabase
        }
    }

    return 0, nil
}

// Insert insert a document to collection.
func (d *DeliveryCity) Insert() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableDeliveryCity)
    d.ID = bson.NewObjectId()

    err = c.Insert(d)
    if err != nil {
        if mgo.IsDup(err) {
            code = ErrDupRows
        } else {
            code = ErrDatabase
        }
    } else {
        code = 0
    }

    return
}

func (d *DeliveryCity) Count() (count int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableDeliveryCity)

    count, err = c.FindId(nil).Count()

    return
}

func (d *DeliveryCity) Update() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableDeliveryCity)

    err = c.Update(bson.M{"_id": d.ID}, bson.M{"$set": bson.M{
        "city":     d.City,
        "zip_code": d.ZipCode,
        "active":   d.Active}})

    if err != nil {
        if err == mgo.ErrNotFound {
            return ErrNotFound, err
        }

        return ErrDatabase, err
    }

    return 0, nil
}

func DeliveryCityTotal() (total int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableDeliveryCity)

    total, err = c.Find(nil).Count()

    return
}

func DeliveryCityList() (deliveryCities []DeliveryCity, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableDeliveryCity)

    err = c.Find(nil).All(&deliveryCities)

    return deliveryCities, err
}
