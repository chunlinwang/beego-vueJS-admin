package models

import (
    "app/models/mymongo"
    "github.com/astaxie/beego"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "reflect"
    "strconv"
    "time"
)

type Order struct {
    Id           bson.ObjectId        `bson:"_id"           json:"_id,omitempty"`
    UserId       string               `bson:"user_id"       json:"user_id,omitempty"`
    OrderNumber  string               `bson:"order_number"  json:"order_number,omitempty"`
    shippingMode string               `bson:"shipping_mode" json:"shipping_mode,omitempty"`
    OrderItems   map[string]OrderItem `bson:"order_items"   json:"order_items,omitempty"`
    Total        int                  `bson:"total"         json:"total,omitempty"`
    TotalPromo   int                  `bson:"total_promo"   json:"total_promo,omitempty"`
    PromoCode    string               `bson:"promo_code"    json:"promo_code,omitempty"`
    Status       string               `bson:"status"        json:"status,omitempty"`
    CreateAt     time.Time            `bson:"create_at"     json:"create_at,omitempty"`
}

type OrderItem struct {
    Product Product
    Menu    Menu
    Count   int
}

func init() () {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableOrder)

    index := mgo.Index{
        Key:        []string{"order_number"},
        Unique:     true,
        DropDups:   true,
        Background: true, // See notes.
        Sparse:     true,
    }

    _ = c.EnsureIndex(index)
}

func NewOrder(userId string) (o Order) {
    o = Order{
        Id:       bson.NewObjectId(),
        UserId:   userId,
        CreateAt: time.Now(),
        Status:   "pending"}

    o.Insert()

    return o
}

// Insert insert a document to collection.
func (o *Order) Insert() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableOrder)

    err = c.Insert(o)
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

func (o *Order) FindByID(id string) (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableOrder)
    err = c.FindId(bson.ObjectIdHex(id)).One(o)

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

func OrderList() (orders []Order, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableOrder)

    err = c.Find(nil).All(&orders)

    return orders, err
}

func OrderTotal() (total int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableOrder)

    total, err = c.Find(nil).Count()

    return
}

func (o *Order) Update() (code int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(DBTableOrder)

    code = 0

    beego.Debug(o.OrderItems)
    err = c.Update(bson.M{"_id": o.Id}, bson.M{"$set": bson.M{
        "user_id":       o.UserId,
        "order_number":  o.OrderNumber,
        "shipping_mode": o.shippingMode,
        "order_items":   o.OrderItems,
        "total":         o.Total,
        "total_promo":   o.TotalPromo,
        "promo_code":    o.PromoCode}})

    if err != nil {
        if err == mgo.ErrNotFound {
            return ErrNotFound, err
        }

        return ErrDatabase, err

        code = -1
    }

    return
}

func (o *Order) AddProduct(productId string) (code int, err error) {
    product := Product{}
    if _, err := product.FindByID(productId); err != nil {
        beego.Error("FindOrderById:", err)
    }

    if ! bson.IsObjectIdHex(product.Id.String()) {
        beego.Error("no product", err)

        return -1, err
    }

    o.updateOrderItems(product)

    return 0, nil
}

func (o *Order) updateOrderItems(product Product) {
    if orderItem, ok := o.OrderItems[product.Id.String()]; ok {
        orderItem.Count = orderItem.Count + 1
        beego.Debug(orderItem.Count)
        o.OrderItems[product.Id.String()] = orderItem
    } else {
        orderItems := o.OrderItems
        orderItem = OrderItem{
            Product: product,
            Count:   1}
        orderItems[product.Id.String()] = orderItem
        o.OrderItems = orderItems
    }

    o.computePrice()
    o.Update()
}

func (o *Order) computePrice() {
    var total int = 0
    keys := reflect.ValueOf(o.OrderItems).MapKeys()
    for i := 0; i < len(keys); i++ {
        productKey := keys[i].String()
        product := o.OrderItems[productKey].Product
        count := o.OrderItems[productKey].Count
        total += product.GetPrice() * count
    }

    o.Total = total
}

func (o *Order) ApplyPromoCode(promoCode PromoCode) (code int, msgErr string) {
    var promoPrice int = o.Total
    if promoCode.MinConsomm == 0 || (promoCode.MinConsomm > 0 && o.Total > promoCode.MinConsomm) {
        o.PromoCode = promoCode.Code
        if promoCode.Type == "euro" {
            promoPrice = promoPrice - promoCode.Value
        } else {
            promoPrice = promoPrice * (1 - promoCode.Value)
        }
        o.TotalPromo = promoPrice

        return 0, ""
    }

    return -1, "minConsomm " + strconv.Itoa(promoCode.MinConsomm)
}
