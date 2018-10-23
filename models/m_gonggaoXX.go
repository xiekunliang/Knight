package models

import (
	"Knight/utils"
	"errors"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// 公告信息
type GonggaoXX struct {
	ID            string      `orm:"column(id);pk;size(255)" json:"id"`                                    //主键
	YongrenSL     int64       `orm:"column(yongrenSL);null" json:"yongrenSL"`                              //用人数量
	YongjinFH     float64     `orm:"digits(12);decimals(2);column(yongjinFH);null" json:"yongjinFH"`       //佣金返还
	ShifouTGZS    bool        `orm:"column(shifouTGZS);default(0)" json:"shifouTGZS"`                      //是否提住房
	ShifouTGCL    bool        `orm:"column(shifouTGCL);default(0)" json:"shifouTGCL"`                      //是否提供车辆
	XinziDY       float64     `orm:"digits(12);decimals(2);column(xinziDY);null" json:"xinziDY"`           //薪资待遇
	ZhuangtaiID   string      `orm:"column(zhuangtaiID);null" json:"zhuangtaiID"`                          //紧急状态ID
	ZhandianID    string      `orm:"column(zhandianID);null" json:"zhandianID"`                            //站点ID
	ChuangjianRID string      `orm:"column(chuangjianRID);null" json:"chuangjianRID"`                      //创建人ID
	GenggaiRID    string      `orm:"column(genggaiRID);null" json:"genggaiRIDd"`                           //更改人ID
	ChuangjianSJ  time.Time   `orm:"column(chuangjianSJ);auto_now_add;type(datetime)" json:"chuangjianSJ"` //创建时间
	ZuihouGGSJ    time.Time   `orm:"column(zuihouGGSJ);auto_now;type(datetime)" json:"zuihouGGSJ"`         //最后更改时间
	DaoqiSJ       time.Time   `orm:"column(daoqiSJ);type(datetime)" json:"daoqiSJ"`                        //到期时间
	ZhandianXX    *ZhandianXX `orm:"rel(fk);null;on_delete(do_nothing)"`                                   //设置一对一关系
	GonggaoZT     *GonggaoZT  `orm:"rel(fk);null;on_delete(do_nothing)"`                                   //设置一对一关系
	DengluYH      *DengluYH   `orm:"rel(fk);null;on_delete(do_nothing)"`                                   //设置一对一关系
}

func init() {
	orm.RegisterModel(new(GonggaoXX))
}

//Table Name
func (u *GonggaoXX) TableName() string {
	return "gonggaoXX"
}

// AddGonggaoXX insert a new GonggaoXX into database and returns last inserted Id on success.
func AddGonggaoXX(m *GonggaoXX, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// DeleteGonggaoXXByID delete by ID
func DeleteGonggaoXXByID(id string, ormObj orm.Ormer) (num int64, err error) {
	obj := &GonggaoXX{ID: id}
	num, err = ormObj.Delete(obj)
	return
}

// UpdateGonggaoXX update GonggaoXX into database and returns id on success
func UpdateGonggaoXX(m *GonggaoXX, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Update(m)
	return
}

// GetGonggaoXXByID retrieves GonggaoXX by ID. Returns error if ID doesn't exist
func GetGonggaoXXByID(id string, ormObj orm.Ormer) (obj *GonggaoXX, err error) {
	obj = &GonggaoXX{ID: id}
	err = ormObj.Read(obj)
	return obj, err
}

// GetAllGonggaoXX retrieves all GonggaoXX matches certain condition. Returns empty list if no records exist
func GetAllGonggaoXX(o orm.Ormer, query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{},
	fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []GonggaoXX, error) {
	var (
		objArrs   []GonggaoXX
		err       error
		paginator utils.Paginator
		num       int64
	)
	qs := o.QueryTable(new(GonggaoXX))
	qs = qs.RelatedSel()

	//cond k=v cond必须放到Filter和Exclude前面
	cond := orm.NewCondition()
	if _, ok := condMap["and"]; ok {
		andMap := condMap["and"]
		for k, v := range andMap {
			k = strings.Replace(k, ".", "__", -1)
			cond = cond.And(k, v)
		}
	}
	if _, ok := condMap["or"]; ok {
		orMap := condMap["or"]
		for k, v := range orMap {
			k = strings.Replace(k, ".", "__", -1)
			cond = cond.Or(k, v)
		}
	}
	qs = qs.SetCond(cond)
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	//exclude k=v
	for k, v := range exclude {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Exclude(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + strings.Replace(v, ".", "__", -1)
				} else if order[i] == "asc" {
					orderby = strings.Replace(v, ".", "__", -1)
				} else {
					return paginator, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + strings.Replace(v, ".", "__", -1)
				} else if order[0] == "asc" {
					orderby = strings.Replace(v, ".", "__", -1)
				} else {
					return paginator, nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return paginator, nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return paginator, nil, errors.New("Error: unused 'order' fields")
		}
	}

	qs = qs.OrderBy(sortFields...)
	if cnt, err := qs.Count(); err == nil {
		if cnt > 0 {
			paginator = utils.GenPaginator(limit, offset, cnt)
			if num, err = qs.Limit(limit, offset).All(&objArrs, fields...); err == nil {
				paginator.CurrentPageSize = num
			}
		}
	}
	return paginator, objArrs, err
}
