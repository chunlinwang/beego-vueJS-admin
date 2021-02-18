package models

import (
    "app/models/mymongo"
    "crypto/rand"
    "fmt"
    "io"
    "reflect"
    "time"

    "golang.org/x/crypto/scrypt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

// User model definiton.
type User struct {
    // install go get github.com/liip/sheriff
    // `json:"name" groups:"trending,detail"` add json Group
    ID          bson.ObjectId `bson:"_id"          json:"_id,omitempty"`
    LastName    string        `bson:"last_name"    json:"last_name,omitempty"`
    FirstName   string        `bson:"first_name"   json:"first_name,omitempty"`
    Phone       string        `bson:"phone"        json:"phone,omitempty"`
    Email       string        `bson:"email"        json:"email,omitempty"`
    Civility    string        `bson:"civility"     json:"civility,omitempty"`
    Password    string        `bson:"password"     json:"password,omitempty"`
    Address     string        `bson:"address"      json:"address,omitempty"`
    SupAddress  string        `bson:"sup_address"  json:"sup_address,omitempty"`
    ZipCode     string        `bson:"zip_code"     json:"zip_code,omitempty"`
    City        string        `bson:"city"         json:"city,omitempty"`
    Admin       bool          `bson:"admin"        json:"admin,false"`
    Salt        string        `bson:"salt"         json:"salt,omitempty"`
    PasswdToken string        `bson:"passwd_token" json:"passwd_token,omitempty"`
    RegDate     time.Time     `bson:"reg_date"     json:"reg_date,omitempty"`
    LastLogin   time.Time     `bson:"last_login"   json:"last_login,omitempty"`
    Roles       []string      `bson:"roles"        json:"roles,[]"`
}

const (
    man              = "m"
    woman            = "f"
    pwHashBytes      = 64
    pwTokenHashBytes = 32
)

func init() () {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableUser)

    index := mgo.Index{
        Key:        []string{"email"},
        Unique:     true,
        DropDups:   true,
        Background: true, // See notes.
        Sparse:     true,
    }

    _ = c.EnsureIndex(index)
}

func generateSalt() (salt string, err error) {
    buf := make([]byte, pwHashBytes)
    if _, err := io.ReadFull(rand.Reader, buf); err != nil {
        return "", err
    }

    return fmt.Sprintf("%x", buf), nil
}

func (user *User) GeneratePwToken() (code int, err error) {
    buf := make([]byte, pwTokenHashBytes)
    if _, err := io.ReadFull(rand.Reader, buf); err != nil {
        return -1, err
    }

    return user.UserUpdateField("passwd_token", fmt.Sprintf("%x", buf));
}

func generatePassHash(password string, salt string) (hash string, err error) {
    h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("%x", h), nil
}

// NewUser alloc and initialize a user.
func NewUser(r *RegisterForm, t time.Time) (u *User, err error) {
    salt, err := generateSalt()
    if err != nil {
        return nil, err
    }
    hash, err := generatePassHash(r.Password, salt)
    if err != nil {
        return nil, err
    }

    user := User{
        ID:         bson.NewObjectId(),
        FirstName:  r.FirstName,
        Phone:      r.Phone,
        LastName:   r.LastName,
        Civility:   r.Civility,
        Address:    r.Address,
        SupAddress: r.SupAddress,
        Email:      r.Email,
        ZipCode:    r.ZipCode,
        City:       r.City,
        Password:   hash,
        Salt:       salt,
        RegDate:    t,
        Admin:      r.Admin,
        Roles:      r.Roles,
    }

    return &user, nil
}

// Insert insert a document to collection.
func (u *User) Insert() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableUser)

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

// FindByID query a document according to input id.
func (u *User) FindByID(id string) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableUser)

    err = c.FindId(bson.ObjectIdHex(id)).One(u)

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

// FindByEmail query a document according to input email.
func (u *User) FindByEmail(email string) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableUser)

    err = c.Find(bson.M{"email": email}).One(u)
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

// CheckPass compare input password.
func (u *User) CheckPass(pass string) (ok bool, err error) {
    hash, err := generatePassHash(pass, u.Salt)
    if err != nil {
        return false, err
    }

    return u.Password == hash, nil
}

// ClearPass clear password information.
func (u *User) ClearPass() {
    u.Password = ""
    u.Salt = ""
}

// ChangePass update password and salt information according to input id.
func ChangePass(id, oldPass, newPass string) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableUser)
    u := User{}
    err = c.FindId(id).One(&u)
    if err != nil {
        if err == mgo.ErrNotFound {
            return ErrNotFound, err
        }

        return ErrDatabase, err
    }

    oldHash, err := generatePassHash(oldPass, u.Salt)
    if err != nil {
        return ErrSystem, err
    }
    newSalt, err := generateSalt()
    if err != nil {
        return ErrSystem, err
    }
    newHash, err := generatePassHash(newPass, newSalt)
    if err != nil {
        return ErrSystem, err
    }

    err = c.Update(bson.M{"_id": id, "password": oldHash}, bson.M{"$set": bson.M{"password": newHash, "salt": newSalt}})
    if err != nil {
        if err == mgo.ErrNotFound {
            return ErrNotFound, err
        }

        return ErrDatabase, err
    }

    return 0, nil
}

func (user *User) UserUpdateField(field string, val interface{}) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableUser)

    var bsonValue interface{}
    bsonValue = bson.M{reflect.ValueOf(field).String(): val}

    err = c.Update(bson.M{"_id": user.ID}, bson.M{"$set": bsonValue})
    if err != nil {
        if err == mgo.ErrNotFound {
            return ErrNotFound, err
        }

        return ErrDatabase, err
    }

    return 0, nil
}

// Update a user.
func (user *User) UserUpdate(u *UpdateUserForm) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableUser)

    err = c.Update(bson.M{"_id": user.ID}, bson.M{"$set": bson.M{
        "first_name":  u.FirstName,
        "phone":       u.Phone,
        "last_name":   u.LastName,
        "civility":    u.Civility,
        "address":     u.Address,
        "sup_address": u.SupAddress,
        "zip_code":    u.ZipCode,
        "city":        u.City,
    }})

    if err != nil {
        if err == mgo.ErrNotFound {
            return ErrNotFound, err
        }

        return ErrDatabase, err
    }

    return 0, nil
}
