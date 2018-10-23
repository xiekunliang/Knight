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

// ServiceCreateMokuaiPZ 创建记录
func ServiceCreateMokuaiPZ(user *md.DengluYH, requestBody []byte) (id int64, err error) {
	var obj md.MokuaiPZ
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
	id, err = md.AddMokuaiPZ(&obj, o)

	return
}

// ServiceUpdateMokuaiPZ 更新记录
func ServiceUpdateMokuaiPZ(user *md.DengluYH, requestBody []byte) (id int64, err error) {
	var obj md.MokuaiPZ
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
	id, err = md.UpdateMokuaiPZ(&obj, o)

	return
}

// ServiceDeleteMokuaiPZ 删除记录
func ServiceDeleteMokuaiPZ(user *md.DengluYH, requestBody []byte) (num int64, err error) {
	var obj md.MokuaiPZ
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
	num, err = md.DeleteMokuaiPZByID(obj.ID, o)
	return
}

// ServiceCheckMokuaiPZ 登录用户权限检查
func ServiceCheckMokuaiPZ(user *md.DengluYH, moduleName string) (access utils.AccessResult, err error) {
	var (
		groups  []md.QuanxianPZ
		modules []md.MokuaiPZ
	)
	// 若为系统管理员拥有所有的权限
	// if user.IsAdmin {
	access.Create = true
	access.Update = true
	access.Read = true
	access.Delete = true
	return
	// }
	// 获得用户所有的权限组
	if groups, err = ServiceGetQuanxianGroupsByID(user.ID); err == nil {
		// 获得权限组下所有的模块访问权限
		leng := len(groups)
		if leng > 0 {
			groupIDs := make([]string, leng, leng)
			for index, group := range groups {
				groupIDs[index] = group.ID
			}
			query := make(map[string]interface{})
			exclude := make(map[string]interface{})
			cond := make(map[string]map[string]interface{})
			condAnd := make(map[string]interface{})
			fields := make([]string, 0, 0)
			sortby := make([]string, 0, 1)
			order := make([]string, 0, 1)
			o := orm.NewOrm()
			condAnd["quanxianZ.id__in"] = groupIDs
			condAnd["mokuaiZ.mokuai"] = moduleName
			if len(condAnd) > 0 {
				cond["and"] = condAnd
			}
			if _, modules, err = md.GetAllMokuaiPZ(o, query, exclude, cond, fields, sortby, order, 0, 0); err == nil {
				for _, module := range modules {
					access.Create = module.Chuangjian || access.Create
					access.Update = module.Gengxin || access.Update
					access.Read = module.Duqu || access.Read
					access.Delete = module.Shanchu || access.Delete
				}

			}
		} else {
			err = errors.New("user has no  any permissions")
		}
	}
	return
}

//ServiceGetMokuaiPZ get MokuaiPZ info
func ServiceGetMokuaiPZ(user *md.DengluYH, query map[string]interface{}, exclude map[string]interface{},
	condMap map[string]map[string]interface{}, fields []string, sortby []string, order []string,
	offset int64, limit int64) (access utils.AccessResult, paginator utils.Paginator, results []map[string]interface{}, err error) {
	if access, err = ServiceCheckMokuaiPZ(user, new(md.MokuaiPZ).TableName()); err == nil {
		if !access.Read {
			err = errors.New("has no read permission")
			return
		}
	} else {
		return
	}
	var arrs []md.MokuaiPZ
	o := orm.NewOrm()
	if paginator, arrs, err = md.GetAllMokuaiPZ(o, query, exclude, condMap, fields, sortby, order, offset, limit); err == nil {
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
