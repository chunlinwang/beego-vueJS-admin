package models

import (
    "app/models/mymongo"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Menu struct {
    Id          bson.ObjectId `bson:"_id"    json:"_id,omitempty"`
    Name        string        `bson:"name"  json:"name,omitempty"`
    Price       int           `bson:"price"  json:"price,omitempty"`
    PriceOnSale int           `bson:"price_on_sale" json:"price_on_sale,omitempty"`
    OnSale      bool          `bson:"on_sale"       json:"on_sale,omitempty"`
    Entries     []ProductItem `bson:"entries"  json:"entries,omitempty"`
    Plats       []ProductItem `bson:"plats"  json:"plats,omitempty"`
    Desserts    []ProductItem `bson:"desserts"  json:"desserts,omitempty"`
    Drinks      []ProductItem `bson:"drinks"  json:"drinks,omitempty"`
}

type ProductItem struct {
    Product    Product `bson:"product"  json:"product,omitempty"`
    ExtraPrice int     `bson:"extra_price"  json:"extra_price,omitempty"`
}

func init() () {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableMenu)

    index := mgo.Index{
        Key:        []string{"code"},
        Unique:     true,
        DropDups:   true,
        Background: true, // See notes.
        Sparse:     true,
    }

    _ = c.EnsureIndex(index)
}

func (m *Menu) FindByID(id string) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableMenu)
    err = c.FindId(bson.ObjectIdHex(id)).One(m)

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

// Insert insert a document to collection.
func (u *Menu) Insert() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableMenu)

    err = c.Insert(u)
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

func MenuListByQuery(category string, query string) (menus []Menu, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableMenu)

    err = c.Find(bson.M{
        "$or": []bson.M{bson.M{"code": bson.M{"$regex": query}},
            bson.M{"title": bson.M{"$regex": query}}}}).All(&menus)

    return menus, err
}

func MenuList() (menus []Menu, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableMenu)

    err = c.Find(nil).All(&menus)

    return menus, err
}

func (m *Menu) New() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableMenu)

    m.Id = bson.NewObjectId()

    if err = c.Insert(m); err != nil {
        return -1, err
    }

    return 0, nil
}

func (m *Menu) MenuUpdate() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableMenu)

    err = c.Update(bson.M{"_id": m.Id}, bson.M{"$set": bson.M{
        "name":          m.Name,
        "price":         m.Price,
        "price_on_sale": m.PriceOnSale,
        "on_sale":       m.OnSale,
        "entries":       m.Entries,
        "plats":         m.Plats,
        "desserts":      m.Desserts,
        "drinks":        m.Drinks,
    }})

    if err != nil {
        if err == mgo.ErrNotFound {
            return ErrNotFound, err
        }

        return ErrDatabase, err
    }

    return 0, nil
}

func MenuTotal() (total int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableMenu)

    total, err = c.Find(nil).Count()

    return
}
