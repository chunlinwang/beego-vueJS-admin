package models

import (
    "app/models/mymongo"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "strings"
    "unicode"
)

type Page struct {
    Id             bson.ObjectId `bson:"_id"   json:"_id,omitempty"`
    Title          string        `bson:"title"   json:"title,omitempty"`
    Code           string        `bson:"code"   json:"code,omitempty"`
    Content        string        `bson:"content"   json:"content,omitempty"`
    SeoTitle       string        `bson:"seo_title"   json:"seo_title,omitempty"`
    SeoDescription string        `bson:"seo_description"   json:"seo_description,omitempty"`
}

type PageListResponse struct {
    Items []Page `json:"items,omitempty"`
    Total int    `json:"total,0"`
}

func init() () {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTablePage)

    index := mgo.Index{
        Key:        []string{"code"},
        Unique:     true,
        DropDups:   true,
        Background: true, // See notes.
        Sparse:     true,
    }

    _ = c.EnsureIndex(index)
}

func (p *Page) FindByID(id string) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTablePage)
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

func PageTotal() (total int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTablePage)

    total, err = c.Find(nil).Count()

    return
}

func NewPage(f *PageForm) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTablePage)
    page := Page{
        Id:             bson.NewObjectId(),
        Code:           f.Code,
        Title:          f.Title,
        Content:        f.Content,
        SeoTitle:       f.SeoTitle,
        SeoDescription: f.SeoDescription}

    if err = c.Insert(page); err != nil {
        return -1, err
    }

    return 0, nil
}

func (p *Page) New() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTablePage)
    page := Page{
        Id:             bson.NewObjectId(),
        Code:           strings.ToLowerSpecial(unicode.AzeriCase, strings.Replace(p.Title, " ", "-", -1)),
        Title:          p.Title,
        Content:        p.Content,
        SeoTitle:       p.SeoTitle,
        SeoDescription: p.SeoDescription}

    if err = c.Insert(page); err != nil {
        return -1, err
    }

    return 0, nil
}

func (p *Page) PageUpdate() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTablePage)

    err = c.Update(bson.M{"_id": p.Id}, bson.M{"$set": bson.M{
        "title":           p.Title,
        "content":         p.Content,
        "seo_title":       p.SeoTitle,
        "seo_description": p.SeoDescription}})

    if err != nil {
        if err == mgo.ErrNotFound {
            return ErrNotFound, err
        }

        return ErrDatabase, err
    }

    return 0, nil
}

func PageList() (pages []Page, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTablePage)

    err = c.Find(nil).All(&pages)

    return pages, err
}
