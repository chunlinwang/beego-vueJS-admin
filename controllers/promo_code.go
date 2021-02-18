package controllers

import (
    "app/models"

    "github.com/astaxie/beego"
)

// PromoCodeController definiton.
type PromoCodeController struct {
    BaseController
}

func (c *PromoCodeController) Get() {
    code := c.Ctx.Input.Param(":id");
    promoCode := models.PromoCode{}
    if code, err := promoCode.FindByID(code); err != nil {
        beego.Error("FindUserById:", err)
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = promoCode
    c.ServeJSON()
}

func (c *PromoCodeController) List() {
    promoCodes, _ := models.PromoCodeList()
    total, _ := models.PromoCodeTotal()

    response := models.ListResponse{
        Items: promoCodes,
        Total: total}

    c.Data["json"] = response
    c.ServeJSON()
}

func (c *PromoCodeController) New() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    form := models.PromoCodeForm{}
    if err := c.ParseForm(&form); err != nil {
        beego.Debug("PromoCodeForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }
    beego.Debug("PromoCodeForm:", &form)

    if err := c.VerifyForm(&form); err != nil {
        beego.Debug("PromoCodeForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }

    promoCode, err := models.NewPromoCode(&form)
    if err != nil {
        beego.Error("NewPromoCode:", err)
        c.Data["json"] = models.NewErrorInfo(ErrSystem)
        c.ServeJSON()
        return
    }
    beego.Debug("NewPromoCode:", promoCode)

    c.Data["json"] = models.NewNormalInfo("Succes")
    c.ServeJSON()
}

func (c *PromoCodeController) Update() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    code := c.Ctx.Input.Param(":id")
    form := models.PromoCodeUpdateForm{}
    if err := c.ParseForm(&form); err != nil {
        beego.Debug("PromoCodeUpdateForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }
    beego.Debug("PromoCodeUpdateForm:", &form)

    if err := c.VerifyForm(&form); err != nil {
        beego.Debug("PromoCodeUpdateForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }

    beego.Debug("PromoCode update:", &form)

    promoCode := models.PromoCode{}
    if code, err := promoCode.FindByID(code); err != nil {
        beego.Error("FindUserById:", err)
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }
        c.ServeJSON()
        return
    }

    promoCode.PromoCodeUpdate(&form)

    beego.Debug("PromoCode update:", &promoCode)

    c.Data["json"] = models.NewNormalInfo("Succes")
    c.ServeJSON()
    return
}
