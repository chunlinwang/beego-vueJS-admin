package controllers

import (
	"app/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// PageController definiton.
type PageController struct {
	BaseController
}

func (c *PageController) Get() {
	code := c.Ctx.Input.Param(":id");
	page := models.Page{}
	if code, err := page.FindByID(code); err != nil {
		beego.Error("FindUserById:", err)
		if code == models.ErrNotFound {
			c.Data["json"] = models.NewErrorInfo(ErrNoUser)
		} else {
			c.Data["json"] = models.NewErrorInfo(ErrDatabase)
		}
		c.ServeJSON()
		return
	}

	c.Data["page_title"] = page.Title
	c.Data["page_content"] = page.Content
	c.Data["page_seo_description"] = page.SeoDescription
	c.Data["page_seo_title"] = page.SeoTitle
	c.TplName = "page.tpl"
	c.Render()
}

func (c *PageController) GetPage() {
	code := c.Ctx.Input.Param(":id");
	page := models.Page{}
	if code, err := page.FindByID(code); err != nil {
		beego.Error("FindUserById:", err)
		if code == models.ErrNotFound {
			c.Data["json"] = models.NewErrorInfo(ErrNoUser)
		} else {
			c.Data["json"] = models.NewErrorInfo(ErrDatabase)
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = page
	c.ServeJSON()
}

func (c *PageController) List() {
	pages, _ := models.PageList()

	total, _ := models.PageTotal()

	response := models.PageListResponse {
		Items: pages,
		Total: total}

	c.Data["json"] = response
	c.ServeJSON()
}

func (c *PageController) New() {
	//form := models.PageForm{}
	//if err := c.ParseForm(&form); err != nil {
	//	beego.Debug("PageForm:", err)
	//	c.Data["json"] = models.NewErrorInfo(ErrInputData)
	//	c.ServeJSON()
	//	return
	//}
	//beego.Debug("PageForm:", &form)
	//
	//if err := c.VerifyForm(&form); err != nil {
	//	beego.Debug("PageForm:", err)
	//	c.Data["json"] = models.NewErrorInfo(ErrInputData)
	//	c.ServeJSON()
	//	return
	//}

	//page, err := models.NewPage(&form)
	//if err != nil {
	//	beego.Error("NewPage:", err)
	//	c.Data["json"] = models.NewErrorInfo(ErrSystem)
	//	c.ServeJSON()
	//	return
	//}
	//beego.Debug("NewPage:", page)

	var err error = nil
	page := models.Page{}
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &page); err == nil {
		_, err = page.New()
	}

	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	c.Data["json"] = models.NewNormalInfo("Succes")
	c.ServeJSON()
}

func (c *PageController) Update() {
	if _, err := c.ParseToken(); err != nil {
		c.Data["json"] = models.NewErrorInfo("nologin")
		c.ServeJSON()
		return
	}

	id := c.Ctx.Input.Param(":id")

	page := models.Page{}
	if code, err := page.FindByID(id); err != nil {
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
	if err = json.Unmarshal(c.Ctx.Input.RequestBody, &page); err == nil {
		page.PageUpdate()
	} else {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	c.Data["json"] = models.NewNormalInfo("Succes")
	c.ServeJSON()
}
