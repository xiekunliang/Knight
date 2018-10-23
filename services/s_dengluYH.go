package services

import (
	md "Knight/models"
	"Knight/utils"
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

// ServiceCreateDengluYH 创建记录
func ServiceCreateDengluYH(requestBody []byte, ownerID string) (id int64, err error) {
	obj := &md.DengluYH{}
	bodyJson := make(map[string]map[string]string)
	json.Unmarshal([]byte(requestBody), &bodyJson)
	obj.ID = bodyJson["data"]["id"]
	obj.Xingming = bodyJson["data"]["xingMing"]
	obj.Mima = bodyJson["data"]["miMa"]
	obj.LianxiFS = bodyJson["data"]["shouJi"]
	var access utils.AccessResult
	if access, err = ServiceCheckMokuaiPZ(obj, obj.TableName()); err == nil {
		if !access.Create {
			err = errors.New("has no create permission ")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	err = o.Begin()
	defer func() {
		if err == nil {
			if o.Commit() != nil {
				if errRollback := o.Rollback(); errRollback != nil {
					err = errRollback
				}
			}
		}
	}()
	if err != nil {
		return
	}
	obj.ChuangjianRID = ownerID
	id, err = md.AddDengluYH(obj, o)

	return
}

// ServiceUpdateDengluYH 更新记录
func ServiceUpdateDengluYH(requestBody []byte, ownerID string) (id int64, err error) {
	obj := &md.DengluYH{}
	json.Unmarshal([]byte(requestBody), &obj)
	var access utils.AccessResult
	if access, err = ServiceCheckMokuaiPZ(obj, obj.TableName()); err == nil {
		if !access.Update {
			err = errors.New("has no create permission ")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	err = o.Begin()
	defer func() {
		if err == nil {
			if o.Commit() != nil {
				if errRollback := o.Rollback(); errRollback != nil {
					err = errRollback
				}
			}
		}
	}()
	if err != nil {
		return
	}
	obj.GenggaiRID = ownerID
	obj.Mima = utils.PasswordMD5(obj.Mima, obj.YonghuM)
	id, err = md.UpdateDengluYH(obj, o)

	return
}

// ServiceDeleteDengluYH 删除记录
func ServiceDeleteDengluYH(user *md.DengluYH, requestBody []byte) (num int64, err error) {
	var obj md.DengluYH
	json.Unmarshal([]byte(requestBody), &obj)
	var access utils.AccessResult
	if access, err = ServiceCheckMokuaiPZ(user, obj.TableName()); err == nil {
		if !access.Update {
			err = errors.New("has no create permission ")
			return
		}
	} else {
		return
	}
	o := orm.NewOrm()
	err = o.Begin()
	defer func() {
		if err == nil {
			if o.Commit() != nil {
				if errRollback := o.Rollback(); errRollback != nil {
					err = errRollback
				}
			}
		}
	}()
	if err != nil {
		return
	}
	if strings.TrimSpace(obj.ID) == "" {
		err = errors.New("ID is null ")
		return
	}
	num, err = md.DeleteDengluYHByID(obj.ID, o)
	return
}

//用户登录
func ServiceUserLogin(name string, password string) (user md.DengluYH, err error) {
	o := orm.NewOrm()
	o.Using("default")
	cond := orm.NewCondition()
	cond = cond.And("active", true).And("yonghuM", name)
	qs := o.QueryTable(&user)
	qs = qs.SetCond(cond)
	if err = qs.One(&user); err == nil {
		if user.Mima == utils.PasswordMD5(password, user.YonghuM) {
			err = nil
		} else {
			err = errors.New("name or password error")
		}
	}
	return
}

//用户登出
func ServiceUserLogout(id string) (ok bool, err error) {
	return
}

//ServiceGetDengluYH get DengluYH info
func ServiceGetDengluYH(user *md.DengluYH, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckMokuaiPZ(user, new(md.DengluYH).TableName()); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.DengluYH
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllDengluYH(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
		lenArrs := len(arrs)
		lenFields := len(fields)
		for i := 0; i < lenArrs; i++ {
			obj := arrs[i]
			objInfo := make(map[string]interface{})
			for j := 0; j < lenFields; j++ {
				objInfo[fields[i]] = reflect.ValueOf(obj).FieldByName(fields[i])
			}
			results = append(results, objInfo)
		}
	}
	return
}
