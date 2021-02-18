package controllers

import (
    "app/models"
    "app/models/mymongo"
    "encoding/json"
    "github.com/astaxie/beego"
)

// MenuController definiton.
type MenuController struct {
    BaseController
}

func (c *MenuController) Get() {
    id := c.Ctx.Input.Param(":id");
    menu := models.Menu{}
    if code, err := menu.FindByID(id); err != nil {
        beego.Error("FindMenuById:", err)
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = menu
    c.ServeJSON()
}

func (c *MenuController) List() {
    menus, _ := models.MenuList()

    total, _ := models.MenuTotal()

    response := models.ListResponse{
        Items: menus,
        Total: total}

    c.Data["json"] = response
    c.ServeJSON()
}

func MenuTotal() (total int, err error) {
    mConn := mymongo.Conn()
    defer mConn.Close()

    c := mConn.DB("").C(models.DBTableMenu)

    total, err = c.Find(nil).Count()

    return
}

func (c *MenuController) GetByQuery() {
    category := c.Input().Get("category")
    query := c.Input().Get("query")
    menus, _  := models.MenuListByQuery(category, query)

    response := models.ListResponse{
        Items: menus,
        Total: len(menus)}

    c.Data["json"] = response
    c.ServeJSON()
}

func (c *MenuController) New() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    beego.Debug(c.Ctx.Input.RequestBody)

    var menu models.Menu
    var err error
    if err = json.Unmarshal(c.Ctx.Input.RequestBody, &menu); err == nil {
        menu.New()
    } else {
        c.Data["json"] = err.Error()
        c.ServeJSON()
        return
    }

    c.Data["json"] = models.NewNormalInfo("Succes")
    c.ServeJSON()
}

func (c *MenuController) Update() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    id := c.Ctx.Input.Param(":id")

    menu := models.Menu{}
    if code, err := menu.FindByID(id); err != nil {
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
    if err = json.Unmarshal(c.Ctx.Input.RequestBody, &menu); err == nil {
        menu.MenuUpdate()
    } else {
        c.Data["json"] = err.Error()
        c.ServeJSON()
        return
    }

    c.Data["json"] = models.NewNormalInfo("Succes")
    c.ServeJSON()
}
