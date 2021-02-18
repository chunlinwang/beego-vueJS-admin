package models

import (
    "app/models/mymongo"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Product struct {
    Id          bson.ObjectId `bson:"_id"           json:"_id,omitempty"`
    Number      string        `bson:"number"        json:"number,omitempty"`
    Category    string        `bson:"category"        json:"category,omitempty"`
    Name        string        `bson:"name"          json:"name,omitempty"`
    Image       string        `bson:"image"         json:"image,omitempty"`
    Price       int           `bson:"price"         json:"price,omitempty"`
    PriceOnSale int           `bson:"price_on_sale" json:"price_on_sale,omitempty"`
    OnSale      bool          `bson:"on_sale"       json:"on_sale,omitempty"`
}

type ProductListResponse struct {
    Items []Product `json:"items,omitempty"`
    Total int       `json:"total,0"`
}

func init() () {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableProduct)

    index := mgo.Index{
        Key:        []string{"code"},
        Unique:     true,
        DropDups:   true,
        Background: true, // See notes.
        Sparse:     true,
    }

    _ = c.EnsureIndex(index)
}

// Insert insert a document to collection.
func (u *Product) Insert() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableProduct)

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

func (p *Product) New() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableProduct)

    p.Id = bson.NewObjectId()

    if err = c.Insert(p); err != nil {
        return -1, err
    }

    return 0, nil
}

func NewProduct(f *ProductForm) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableProduct)

    product := Product{
        Id:          bson.NewObjectId(),
        Name:        f.Name,
        Number:      f.Number,
        Category:    f.Category,
        Image:       f.Image,
        Price:       f.Price,
        PriceOnSale: f.PriceOnSale,
        OnSale:      f.OnSale}

    if err = c.Insert(product); err != nil {
        return -1, err
    }

    return 0, nil
}

func (p *Product) FindByID(id string) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableProduct)
    err = c.FindId(bson.ObjectIdHex(id)).One(p)

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

func (p *Product) GetPrice() (price int) {
    if (p.OnSale) {
        return p.PriceOnSale
    }

    return p.Price
}

func ProductList() (products []Product, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableProduct)

    err = c.Find(nil).All(&products)

    return products, err
}

func ProductListByQuery(category string, query string) (products []Product, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableProduct)

    err = c.Find(bson.M{
        "category": category,
        "$or": []bson.M{ bson.M{"code" : bson.M{"$regex": query} },
        bson.M{"name" : bson.M{"$regex": query}},
        bson.M{"number" : bson.M{"$regex": query}},
    }}).All(&products)

    return products, err
}

func (p *Product) ProductUpdate() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableProduct)

    err = c.Update(bson.M{"_id": p.Id}, bson.M{"$set": bson.M{
        "name":          p.Name,
        "number":        p.Number,
        "category":      p.Category,
        "image":         p.Image,
        "price":         p.Price,
        "price_on_sale": p.PriceOnSale,
        "on_sale":       p.OnSale}})

    if err != nil {
        if err == mgo.ErrNotFound {
            return ErrNotFound, err
        }

        return ErrDatabase, err
    }

    return 0, nil
}

func ProductTotal() (total int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableProduct)

    total, err = c.Find(nil).Count()

    return
}
