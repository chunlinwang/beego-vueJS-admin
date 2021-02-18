package controllers

import (
    "io"
    "os"
    "fmt"
    "time"
    "crypto/md5"
    "github.com/dgrijalva/jwt-go"
    "net/http"
    "path/filepath"
    "app/models"
    "github.com/astaxie/beego"
)

// UserController definiton.
type UserController struct {
    BaseController
}

// Register method.
func (c *UserController) Register() {

    beego.Debug(c.Input().Get("admin"))
    beego.Debug(c.Input().Get("roles"))
    form := models.RegisterForm{}
    if err := c.ParseForm(&form); err != nil {
        beego.Debug("ParseRegsiterForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }
    beego.Debug("ParseRegsiterForm:", &form)

    if err := c.VerifyForm(&form); err != nil {
        beego.Debug("ValidRegsiterForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }

    regDate := time.Now()
    user, err := models.NewUser(&form, regDate)
    if err != nil {
        beego.Error("NewUser:", err)
        c.Data["json"] = models.NewErrorInfo(ErrSystem)
        c.ServeJSON()
        return
    }
    beego.Debug("NewUser:", user)

    if code, err := user.Insert(); err != nil {
        beego.Error("InsertUser:", err)
        if code == models.ErrDupRows {
            c.Data["json"] = models.NewErrorInfo(ErrDupUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }
        c.ServeJSON()
        return
    }

    go models.IncTotalUserCount(regDate)

    claims := make(jwt.MapClaims)
    claims["id"] = user.ID
    claims["name"] = form.Email
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString([]byte("secret"))
    if err != nil {
        beego.Error("jwt.SignedString:", err)
        c.RetError(errSystem)
        return
    }

    c.SetSession("user_id", form.Email)
    user.UserUpdateField("last_login", time.Now())

    c.Data["json"] = &models.RoleAuthInfo{Token: "Bearer " + tokenString}
    c.ServeJSON()
}

// Login method.
func (c *UserController) Login() {
    form := models.LoginForm{}

    if err := c.ParseForm(&form); err != nil {
        beego.Debug("ParseLoginForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }
    beego.Debug("ParseLoginForm:", &form)

    if err := c.VerifyForm(&form); err != nil {
        beego.Debug("ValidLoginForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }

    user := models.User{}
    if code, err := user.FindByEmail(form.Email); err != nil {
        beego.Error("FindByEmail:", err)
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }
        c.ServeJSON()
        return
    }
    beego.Debug("user found.")
    if ok, err := user.CheckPass(form.Password); err != nil {
        beego.Error("CheckUserPass:", err)
        c.Data["json"] = models.NewErrorInfo(ErrSystem)
        c.ServeJSON()
        return
    } else if !ok {
        c.Data["json"] = models.NewErrorInfo(ErrPass)
        c.ServeJSON()
        return
    }
    user.ClearPass()

    // Create the token with some claims
    claims := make(jwt.MapClaims)
    claims["id"] = user.ID
    claims["name"] = form.Email
    claims["admin"] = user.Admin
    claims["exp"] = time.Now().Add(time.Hour * 24000).Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString([]byte("secret"))
    if err != nil {
        beego.Error("jwt.SignedString:", err)
        c.RetError(errSystem)
        return
    }

    c.SetSession("user_id", form.Email)
    user.UserUpdateField("last_login", time.Now())

    if user.Admin {
        c.Data["json"] = &models.AdminRoleAuthInfo{Token: "Bearer " + tokenString, IsAdmin: true, Roles: user.Roles}
    } else {
        c.Data["json"] = &models.RoleAuthInfo{Token: "Bearer " + tokenString}
    }
    c.Ctx.SetCookie("My-Authorization", "Bearer " + tokenString)
    c.ServeJSON()
}

// Login method.
func (c *UserController) GetInfo() {
    token, err := c.ParseToken()
    if err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    claims := token.Claims.(jwt.MapClaims)

    user := models.User{}

    user.FindByID(claims["id"].(string))

    beego.Debug(user)

    c.Data["json"] = user
    c.ServeJSON()
}

// Logout method.
func (c *UserController) Logout() {
    form := models.LogoutForm{}
    if err := c.ParseForm(&form); err != nil {
        beego.Debug("ParseLogoutForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }
    beego.Debug("ParseLogoutForm:", &form)

    if err := c.VerifyForm(&form); err != nil {
        beego.Debug("ValidLogoutForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }

    if c.GetSession("user_id") != form.Email {
        c.Data["json"] = models.NewErrorInfo(ErrInvalidUser)
        c.ServeJSON()
        return
    }

    c.DelSession("user_id")

    c.Data["json"] = models.NewNormalInfo("Succes")
    c.Ctx.SetCookie("My-Authorization", "");
    c.ServeJSON()
}

// Passwd method.
func (c *UserController) Passwd() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    form := models.PasswdForm{}
    if err := c.ParseForm(&form); err != nil {
        beego.Debug("ParsePasswdForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }
    beego.Debug("ParsePasswdForm:", &form)

    if err := c.VerifyForm(&form); err != nil {
        beego.Debug("ValidPasswdForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }

    if c.GetSession("user_id") != form.Email {
        c.Data["json"] = models.NewErrorInfo(ErrInvalidUser)
        c.ServeJSON()
        return
    }

    code, err := models.ChangePass(form.Email, form.OldPass, form.NewPass)
    if err != nil {
        beego.Error("ChangeUserPass:", err)
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUserPass)
        } else if code == models.ErrDatabase {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrSystem)
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = models.NewNormalInfo("Succes")
    c.ServeJSON()
}

// Uploads method.
func (c *UserController) Uploads() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    form := models.UploadsForm{}
    if err := c.ParseForm(&form); err != nil {
        beego.Debug("ParseUploadsForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }
    beego.Debug("ParseUploadsForm:", &form)

    if err := c.VerifyForm(&form); err != nil {
        beego.Debug("ValidUploadsForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }

    if c.GetSession("user_id") != form.Email {
        c.Data["json"] = models.NewErrorInfo(ErrInvalidUser)
        c.ServeJSON()
        return
    }

    //files := c.Ctx.Request.MultipartForm.File["photos"]
    files, err := c.GetFiles("photos")
    if err != nil {
        beego.Debug("GetFiles:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }
    for i := range files {
        src, err := files[i].Open()
        if err != nil {
            beego.Error("Open MultipartForm File:", err)
            c.Data["json"] = models.NewErrorInfo(ErrOpenFile)
            c.ServeJSON()
            return
        }
        defer src.Close()

        hash := md5.New()
        if _, err := io.Copy(hash, src); err != nil {
            beego.Error("Copy File to Hash:", err)
            c.Data["json"] = models.NewErrorInfo(ErrWriteFile)
            c.ServeJSON()
            return
        }
        hex := fmt.Sprintf("%x", hash.Sum(nil))

        dst, err := os.Create(beego.AppConfig.String("apppath") + "static/" + hex + filepath.Ext(files[i].Filename))
        if err != nil {
            beego.Error("Create File:", err)
            c.Data["json"] = models.NewErrorInfo(ErrWriteFile)
            c.ServeJSON()
        }
        defer dst.Close()

        src.Seek(0, 0)
        if _, err := io.Copy(dst, src); err != nil {
            beego.Error("Copy File:", err)
            c.Data["json"] = models.NewErrorInfo(ErrWriteFile)
            c.ServeJSON()
            return
        }
    }

    c.Data["json"] = models.NewNormalInfo("Succes")
    c.ServeJSON()
}

// Downloads method.
func (c *UserController) Downloads() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    if c.GetSession("user_id") == nil {
        c.Data["json"] = models.NewErrorInfo(ErrInvalidUser)
        c.ServeJSON()
        return
    }

    file := beego.AppConfig.String("apppath") + "logs/test.log"
    http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, file)
}

func (c *UserController) Update() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    idStr := c.Ctx.Input.Param(":id")
    form := models.UpdateUserForm{}
    if err := c.ParseForm(&form); err != nil {
        beego.Debug("UpdateUserForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }
    beego.Debug("UpdateUserForm:", &form)

    if err := c.VerifyForm(&form); err != nil {
        beego.Debug("ValidLoginForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }

    beego.Debug("User update:", &form)

    user := models.User{}
    if code, err := user.FindByID(idStr); err != nil {
        beego.Error("FindUserById:", err)
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }
        c.ServeJSON()
        return
    }

    user.UserUpdate(&form)

    beego.Debug("User update:", &user)

    c.Data["json"] = models.NewNormalInfo("Succes")
    c.ServeJSON()
    return
}

func (c *UserController) RequestPasswd() {
    if _, err := c.ParseToken(); err != nil {
        c.Data["json"] = models.NewErrorInfo("nologin")
        c.ServeJSON()
        return
    }

    form := models.RequestPassWdForm{}
    if err := c.ParseForm(&form); err != nil {
        beego.Debug("RequestPassWdForm:", err)
        c.Data["json"] = models.NewErrorInfo(ErrInputData)
        c.ServeJSON()
        return
    }

    beego.Debug(form.Email)

    user := models.User{}
    if code, err := user.FindByID(form.Email); err != nil {
        beego.Error("FindUserById:", err)
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }
        c.ServeJSON()
        return
    }

    if code, err := user.GeneratePwToken(); err != nil {
        beego.Error("FindUserById:", err)
        if code == models.ErrNotFound {
            c.Data["json"] = models.NewErrorInfo(ErrNoUser)
        } else {
            c.Data["json"] = models.NewErrorInfo(ErrDatabase)
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = models.NewNormalInfo("Succes")
    c.ServeJSON()
}
