package controllers

import (
	"errors"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/dgrijalva/jwt-go"
)

// Predefined const error strings.
const (
	ErrInputData    = "Error: input value is not correct"
	ErrDatabase     = "Error: operation database"
	ErrDupUser      = "Error: duplicate data in database"
	ErrNoUser       = "Error: no data in database"
	ErrNoInDB       = "Error: no data in database"
	ErrPass         = "Error: password is not correct"
	ErrNoUserPass   = "Error no user or password is not correct"
	ErrNoUserChange = "Error no change"
	ErrInvalidUser  = "Error incorrect user"
	ErrOpenFile     = "Error open file"
	ErrWriteFile    = "Error write file"
	ErrSystem       = "Error system"
)

// ControllerError is controller error info structer.
type ControllerError struct {
	Status   int    `json:"status"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
	DevInfo  string `json:"dev_info"`
	MoreInfo string `json:"more_info"`
}

// Predefined controller error values.
var (
	err404          = &ControllerError{404, 404, "page not found", "page not found", ""}
	errInputData    = &ControllerError{400, 10001, "input err", "input err", ""}
	errDatabase     = &ControllerError{500, 10002, "db err", "db err", ""}
	errDupUser      = &ControllerError{400, 10003, "user duplicate", "user duplicate", ""}
	errNoUser       = &ControllerError{400, 10004, "no user", "no user", ""}
	errPass         = &ControllerError{400, 10005, "password is not correct", "password is not correct", ""}
	errNoUserPass   = &ControllerError{400, 10006, "db no user or use has been changed", "db no user or use has been changed", ""}
	errNoUserChange = &ControllerError{400, 10007, "db err", "db err", ""}
	errInvalidUser  = &ControllerError{400, 10008, "Session err", "Session err", ""}
	errOpenFile     = &ControllerError{500, 10009, "server error", "server error", ""}
	errWriteFile    = &ControllerError{500, 10010, "write file error", "write file error", ""}
	errSystem       = &ControllerError{500, 10011, "os error", "os error", ""}
	errExpired      = &ControllerError{400, 10012, "jwt expired", "jwt expired", ""}
	errPermission   = &ControllerError{400, 10013, "no permission", "no permission", ""}
)

// BaseController definiton.
type BaseController struct {
	beego.Controller
}

// RetError return error information in JSON.
func (base *BaseController) RetError(e *ControllerError) {
	if mode := beego.AppConfig.String("runmode"); mode == "prod" {
		e.DevInfo = ""
	}

	base.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	base.Ctx.ResponseWriter.WriteHeader(e.Status)
	base.Data["json"] = e
	base.ServeJSON()

	base.StopRun()
}

var sqlOp = map[string]string{
	"eq": "=",
	"ne": "<>",
	"gt": ">",
	"ge": ">=",
	"lt": "<",
	"le": "<=",
}

// ParseQueryParm parse query parameters.
//   query=col1:op1:val1,col2:op2:val2,...
//   op: one of eq, ne, gt, ge, lt, le
func (base *BaseController) ParseQueryParm() (v map[string]string, o map[string]string, err error) {
	var nameRule = regexp.MustCompile("^[a-zA-Z0-9_]+$")
	queryVal := make(map[string]string)
	queryOp := make(map[string]string)

	query := base.GetString("query")
	if query == "" {
		return queryVal, queryOp, nil
	}

	for _, cond := range strings.Split(query, ",") {
		kov := strings.Split(cond, ":")
		if len(kov) != 3 {
			return queryVal, queryOp, errors.New("Query format != k:o:v")
		}

		var key string
		var value string
		var operator string
		if !nameRule.MatchString(kov[0]) {
			return queryVal, queryOp, errors.New("Query key format is wrong")
		}
		key = kov[0]
		if op, ok := sqlOp[kov[1]]; ok {
			operator = op
		} else {
			return queryVal, queryOp, errors.New("Query operator is wrong")
		}
		value = strings.Replace(kov[2], "'", "\\'", -1)

		queryVal[key] = value
		queryOp[key] = operator
	}

	return queryVal, queryOp, nil
}

// ParseOrderParm parse orders parameters.
//   orders=col1:asc|desc,col2:asc|esc,...
func (base *BaseController) ParseOrderParm() (o map[string]string, err error) {
	var nameRule = regexp.MustCompile("^[a-zA-Z0-9_]+$")
	order := make(map[string]string)

	v := base.GetString("orders")
	if v == "" {
		return order, nil
	}

	for _, cond := range strings.Split(v, ",") {
		kv := strings.Split(cond, ":")
		if len(kv) != 2 {
			return order, errors.New("Order format != k:v")
		}
		if !nameRule.MatchString(kv[0]) {
			return order, errors.New("Order key format is wrong")
		}
		if kv[1] != "asc" && kv[1] != "desc" {
			return order, errors.New("Order val isn't asc/desc")
		}

		order[kv[0]] = kv[1]
	}

	return order, nil
}

// ParseLimitParm parse limit parameter.
//   limit=n
func (base *BaseController) ParseLimitParm() (l int64, err error) {
	if v, err := base.GetInt64("limit"); err != nil {
		return 10, err
	} else if v > 0 {
		return v, nil
	} else {
		return 10, nil
	}
}

// ParseOffsetParm parse offset parameter.
//   offset=n
func (base *BaseController) ParseOffsetParm() (o int64, err error) {
	if v, err := base.GetInt64("offset"); err != nil {
		return 0, err
	} else if v > 0 {
		return v, nil
	} else {
		return 0, nil
	}
}

// VerifyForm use validation to verify input parameters.
func (base *BaseController) VerifyForm(obj interface{}) (err error) {
	valid := validation.Validation{}
	ok, err := valid.Valid(obj)
	if err != nil {
		return err
	}
	if !ok {
		str := ""
		for _, err := range valid.Errors {
			str += err.Key + ":" + err.Message + ";"
		}
		return errors.New(str)
	}

	return nil
}

// ParseToken parse JWT token in http header.
func (base *BaseController) ParseToken() (t *jwt.Token, e *ControllerError) {
	authString := base.Ctx.Input.Header("My-Authorization")
	beego.Debug("AuthString:", authString)

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		beego.Error("AuthString invalid:", authString)
		return nil, errInputData
	}
	tokenString := kv[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		beego.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return nil, errInputData
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, errExpired
			} else {
				// Couldn't handle this token
				return nil, errInputData
			}
		} else {
			// Couldn't handle this token
			return nil, errInputData
		}
	}
	if !token.Valid {
		beego.Error("Token invalid:", tokenString)
		return nil, errInputData
	}
	beego.Debug("Token:", token)

	return token, nil
}
