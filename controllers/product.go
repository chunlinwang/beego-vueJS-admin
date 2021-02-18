package controllers

import (
    "app/models"
    "encoding/json"
    "github.com/astaxie/beego"
)

// ProductController definiton.
type ProductController struct {
    BaseController
}

func (c *ProductController) Get() {
    id := c.Ctx.Input.Param(":id");
    product := models.Product{}
    if code, err := product.FindByID(id); err != nil {
        beego.Error("FindProductById:", err)
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = product
    c.ServeJSON()
}

func (c *ProductController) List() {
    products, _ := models.ProductList()

    total, _ := models.ProductTotal()

    response := models.ProductListResponse{
        Items: products,
        Total: total}

    c.Data["json"] = response
    c.ServeJSON()
}

func (c *ProductController) GetByQuery() {
    category := c.Input().Get("category")
    query := c.Input().Get("query")
    products, _  := models.ProductListByQuery(category, query)

    response := models.ProductListResponse{
        Items: products,
        Total: len(products)}

    c.Data["json"] = response
    c.ServeJSON()
}

func (c *ProductController) New() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    var product models.Product
    var err error
    if err = json.Unmarshal(c.Ctx.Input.RequestBody, &product); err == nil {
        product.New()
    } else {
        c.Data["json"] = err.Error()
        c.ServeJSON()
        return
    }
    c.Data["json"] = models.NewNormalInfo("Succes")
    c.ServeJSON()
}

func (c *ProductController) Update() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    id := c.Ctx.Input.Param(":id")

    product := models.Product{}
    if code, err := product.FindByID(id); err != nil {
        beego.Error("FindUserById:", err)
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }
        c.ServeJSON()
        return
    }

    var err error
    if err = json.Unmarshal(c.Ctx.Input.RequestBody, &product); err == nil {
        product.ProductUpdate()
    } else {
        c.Data["json"] = err.Error()
        c.ServeJSON()
        return
    }

    c.Data["json"] = models.NewNormalInfo("Succes")
    c.ServeJSON()
}
