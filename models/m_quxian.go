package models

import (
	"Knight/utils"
	"errors"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// 区县
type Quxian struct {
	ID            string         `orm:"column(id);pk;size(255)" json:"id"`                                    //主键
	ChengshiID    string         `orm:"column(chengshiID);null" json:"chengshiID"`                            //城市ID
	Quxian        string         `orm:"column(quxian);size(255)" json:"quxian"`                               //区县
	ChuangjianRID string         `orm:"column(chuangjianRID);null" json:"chuangjianRID"`                      //创建人ID
	GenggaiRID    string         `orm:"column(genggaiRID);null" json:"genggaiRIDd"`                           //更改人ID
	ChuangjianSJ  time.Time      `orm:"column(chuangjianSJ);auto_now_add;type(datetime)" json:"chuangjianSJ"` //创建时间
	ZuihouGGSJ    time.Time      `orm:"column(zuihouGGSJ);auto_now;type(datetime)" json:"zuihouGGSJ"`         //最后更改时间
	Chengshi      *Chengshi      `orm:"rel(fk);null;on_delete(do_nothing)"`                                   //设置一对一关系
	DengluYH      *DengluYH      `orm:"rel(fk);null;on_delete(do_nothing)"`                                   //设置一对一关系
	Jiedaos       []*Jiedao      `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	QishouJJLXRs  []*QishouJJLXR `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	QishouXXs     []*QishouXX    `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	ZhandianXXs   []*ZhandianXX  `orm:"reverse(many)"`                                                        //设置一对多的反向关系
}

func init() {
	orm.RegisterModel(new(Quxian))
}

//Table Name
func (u *Quxian) TableName() string {
	return "quxian"
}

// AddQuxian insert a new Quxian into database and returns last inserted Id on success.
func AddQuxian(m *Quxian, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// DeleteQuxianByID delete by ID
func DeleteQuxianByID(id string, ormObj orm.Ormer) (num int64, err error) {
	obj := &Quxian{ID: id}
	num, err = ormObj.Delete(obj)
	return
}

// UpdateQuxian update Quxian into database and returns id on success
func UpdateQuxian(m *Quxian, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Update(m)
	return
}

// GetQuxianByID retrieves Quxian by ID. Returns error if ID doesn't exist
func GetQuxianByID(id string, ormObj orm.Ormer) (obj *Quxian, err error) {
	obj = &Quxian{ID: id}
	err = ormObj.Read(obj)
	return obj, err
}

// GetAllQuxian retrieves all Quxian matches certain condition. Returns empty list if no records exist
func GetAllQuxian(o orm.Ormer, query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{},
	fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []Quxian, error) {
	var (
		objArrs   []Quxian
		err       error
		paginator utils.Paginator
		num       int64
	)
	qs := o.QueryTable(new(Quxian))
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
