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

// ServiceCreateQuanxianPZ 创建记录
func ServiceCreateQuanxianPZ(user *md.DengluYH, requestBody []byte) (id int64, err error) {
	var obj md.QuanxianPZ
	json.Unmarshal([]byte(requestBody), &obj)
	var access utils.AccessResult
	if access, err = ServiceCheckMokuaiPZ(user, obj.TableName()); err == nil {
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
	obj.ChuangjianRID = user.ID
	id, err = md.AddQuanxianPZ(&obj, o)

	return
}

// ServiceUpdateQuanxianPZ 更新记录
func ServiceUpdateQuanxianPZ(user *md.DengluYH, requestBody []byte) (id int64, err error) {
	var obj md.QuanxianPZ
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
	obj.GenggaiRID = user.ID
	id, err = md.UpdateQuanxianPZ(&obj, o)

	return
}

// ServiceDeleteQuanxianPZ 删除记录
func ServiceDeleteQuanxianPZ(user *md.DengluYH, requestBody []byte) (num int64, err error) {
	var obj md.QuanxianPZ
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
	num, err = md.DeleteQuanxianPZByID(obj.ID, o)
	return
}

// ServiceGetQuanxianGroupsByID 获得用户的权限组信息
func ServiceGetQuanxianGroupsByID(userID string) (groups []md.QuanxianPZ, err error) {
	var (
		tGroups []md.QuanxianPZ
	)
	query := make(map[string]interface{})
	exclude := make(map[string]interface{})
	cond := make(map[string]map[string]interface{})
	condAnd := make(map[string]interface{})
	fields := make([]string, 0, 0)
	sortby := make([]string, 0, 1)
	order := make([]string, 0, 1)
	o := orm.NewOrm()
	condAnd["dengluYHID"] = userID
	if len(condAnd) > 0 {
		cond["and"] = condAnd
	}
	if _, tGroups, err = md.GetAllQuanxianPZ(o, query, exclude, cond, fields, sortby, order, 0, 0); err == nil {
		for _, group := range tGroups {
			groups = append(groups, group)
		}
	}
	return
}

//ServiceGetQuanxianPZ get QuanxianPZ info
func ServiceGetQuanxianPZ(user *md.DengluYH, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckMokuaiPZ(user, new(md.QuanxianPZ).TableName()); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.QuanxianPZ
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllQuanxianPZ(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
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
