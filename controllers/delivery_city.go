package controllers

import (
    "app/models"
    "encoding/json"
    "github.com/astaxie/beego"
)

// DeliveryCityController definiton.
type DeliveryCityController struct {
    BaseController
}

func (c *DeliveryCityController) Get() {
    code := c.Ctx.Input.Param(":id");
    deliveryCity := models.DeliveryCity{}
    if code, err := deliveryCity.FindByID(code); err != nil {
        beego.Error("FindUserById:", err)
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = deliveryCity
    c.ServeJSON()
}

func (c *DeliveryCityController) List() {
    deliveryCities, _ := models.DeliveryCityList()
    total, _ := models.DeliveryCityTotal()

    response := models.ListResponse{
        Items: deliveryCities,
        Total: total}

    c.Data["json"] = response
    c.ServeJSON()
}

func (c *DeliveryCityController) New() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    var err error = nil
    deliveryCity := models.DeliveryCity{}
    if err = json.Unmarshal(c.Ctx.Input.RequestBody, &deliveryCity); err == nil {
        _, err = deliveryCity.Insert()
    }

    beego.Debug(deliveryCity)

    if err != nil {
        c.Data["json"] = err.Error()
        c.ServeJSON()
        return
    }

    c.Data["json"] = models.NewNormalInfo("Succes")
    c.ServeJSON()
}

func (c *DeliveryCityController) Update() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }
    id := c.Ctx.Input.Param(":id")

    deliveryCity := models.DeliveryCity{}
    if code, err := deliveryCity.FindByID(id); err != nil {
        beego.Error("FindUserById:", err)
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }
        c.ServeJSON()
        return
    }

    var err error = nil
    if err = json.Unmarshal(c.Ctx.Input.RequestBody, &deliveryCity); err == nil {
        _, err = deliveryCity.Update()
    }

    if err != nil {
        c.Data["json"] = err.Error()
        c.ServeJSON()
        return
    }

    c.Data["json"] = models.NewNormalInfo("Succes")
    c.ServeJSON()
}
