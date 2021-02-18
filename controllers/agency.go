package controllers

import (
    "app/models"
    "github.com/astaxie/beego"
    "strconv"
    "time"
)

// AgencyController definition.
type AgencyController struct {
    BaseController
}

func (c *AgencyController) Update() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    form := models.AgencyForm{}
    if err := c.ParseForm(&form); err != nil {
        beego.Debug("AgencyUpdateForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }
    beego.Debug("AgencyUpdateForm:", &form)

    if err := c.VerifyForm(&form); err != nil {
        beego.Debug("AgencyUpdateForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }

    beego.Debug("Agency update:", &form)

    agency := models.Agency{}
    if code, err := agency.FirstAgency(); err != nil {
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }
        c.ServeJSON()
        return
    }

    var Schedules []models.Schedule
    for i := 0; i < 14; i++ {
        strBeginTime := c.Input().Get("schedules[" + strconv.Itoa(i) + "][begin_hour]")
        strDay := c.Input().Get("schedules[" + strconv.Itoa(i) + "][day]")
        day, _ := strconv.Atoi(strDay)

            beginHour, _ := time.Parse(time.RFC3339, strBeginTime)

            strEndTime := c.Input().Get("schedules[" + strconv.Itoa(i) + "][end_hour]")
            endHour, _ := time.Parse(time.RFC3339, strEndTime)

            schedule := models.Schedule{
                Day:       day,
                BeginHour: beginHour,
                EndHour:   endHour}

            Schedules = append(Schedules, schedule)
        }

    form.Schedules = Schedules

    agency.AgencyUpdate(&form)

    c.Data["json"] = models.NewNormalInfo("Succes")
    c.ServeJSON()
    return
}

func (c *AgencyController) Get() {
    agency := models.Agency{}
    if code, err := agency.FirstAgency(); err != nil {
        beego.Error("FindUserById:", err)
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }

        c.ServeJSON()
        return
    }

    c.Data["json"] = agency
    c.ServeJSON()
}
