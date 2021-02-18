package controllers

import (
	"app/models"
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// OrderController definiton.
type OrderController struct {
	BaseController
}

//func (c *OrderController) Get() {
//	code := c.Ctx.Input.Param(":id");
//	page := models.Page{}
//	if code, err := page.FindByID(code); err != nil {
//		beego.Error("FindOrderById:", err)
//		if code == models.ErrNotFound {
//			c.Data["json"] = models.NewErrorInfo(ErrNoInDB)
//		} else {
//			c.Data["json"] = models.NewErrorInfo(ErrDatabase)
//		}
//		c.ServeJSON()
//		return
//	}
//
//	c.Data["page_title"] = page.Title
//	c.Data["page_content"] = page.Content
//	c.Data["page_seo_description"] = page.SeoDescription
//	c.Data["page_seo_title"] = page.SeoTitle
//	c.TplName = "page.tpl"
//	c.Render()
//}

func (c *OrderController) List() {
	token, err := c.ParseToken()
	if err != nil {
		c.Data["json"] = models.NewErrorInfo("nologin")
		c.ServeJSON()
		return
	}

	claims := token.Claims.(jwt.MapClaims)

	if !claims["admin"].(bool) {
		c.Data["json"] = "not admin"
		c.ServeJSON()
		return
	}

	orders, _ := models.OrderList()

	total, _ := models.OrderTotal()

	response := models.ListResponse{
		Items: orders,
		Total: total}

	c.Data["json"] = response
	c.ServeJSON()
}

func (c *OrderController) CreateOrUpdate() {
	var userId string
	productId := c.Ctx.Input.Param(":product_id")

	if userId := c.GetSession("user_id"); userId == nil {
		c.Data["json"] = "user_not_logged"
		c.ServeJSON()
		return
	}

	beego.Debug(userId)
	order := models.Order{UserId:userId}

	if orderId := c.GetSession("orders"); orderId != nil {
		beego.Debug(orderId)
		if _, err := order.FindByID(orderId.(string)); err != nil {
			c.DelSession("orders")
			order = models.NewOrder(userId)
		}
	} else {
		order = models.NewOrder(userId)
	}

	if _, err := order.AddProduct(productId); err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	c.SetSession("orders", order.Id.String())

	c.Data["json"] = models.NewNormalInfo("Succes")
	c.ServeJSON()
}

func (c *OrderController) ApplyPromoCode() {
	promoCodeId := c.Ctx.Input.GetData(":promo_code")

	promoCode := models.PromoCode{}

	promoCode.FindByID(promoCodeId.(string))

	now := time.Now()
	if ! (now.After(promoCode.BeginDate) && now.Before(promoCode.EndDate) && promoCode.Active) {
		c.Data["json"] = "promo_code_invalid"
		c.ServeJSON()
		return
	}

	if userId := c.GetSession("user_id"); userId == nil {
		c.Data["json"] = "user_not_logged"
		c.ServeJSON()
		return
	}

	orderId := c.GetSession("orders");
	if orderId == nil {
		c.Data["json"] = "no_order"
		c.ServeJSON()
		return
	}

	order := models.Order{}
	order.FindByID(orderId.(string))

	order.ApplyPromoCode(promoCode)
	c.Data["json"] = order
	c.ServeJSON()
}
