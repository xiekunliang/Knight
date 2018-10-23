package models

import (
	"Knight/utils"
	"errors"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

// 登录用户
type DengluYH struct {
	ID            string         `orm:"column(id);pk;size(255)" json:"id"`                                    //主键
	Active        bool           `orm:"column(active);default(1)" json:"active"`                              //是否可用
	IsAdmin       bool           `orm:"column(isadmin);default(0)" json:"isadmin"`                            //是否超级管理员
	YonghuM       string         `orm:"column(yonghuM);null" json:"yonghuM"`                                  //用户名
	Mima          string         `orm:"column(mima);null" json:"mima"`                                        //密码
	Gonghao       string         `orm:"column(gonghao);null" json:"gonghao"`                                  //工号
	LianxiFS      string         `orm:"column(lianxiFS);null" json:"lianxiFS"`                                //联系方式
	Xingbie       string         `orm:"column(xingbie);null" json:"xingbie"`                                  //性别
	Xingming      string         `orm:"column(xingming);null" json:"xingming"`                                //姓名
	ZhandianID    string         `orm:"column(zhandianID);null" json:"zhandianID"`                            //站点ID
	ChuangjianRID string         `orm:"column(chuangjianRID);null" json:"chuangjianRID"`                      //创建人ID
	GenggaiRID    string         `orm:"column(genggaiRID);null" json:"genggaiRIDd"`                           //更改人ID
	ChuangjianSJ  time.Time      `orm:"column(chuangjianSJ);auto_now_add;type(datetime)" json:"chuangjianSJ"` //创建时间
	ZuihouGGSJ    time.Time      `orm:"column(zuihouGGSJ);auto_now;type(datetime)" json:"zuihouGGSJ"`         //最后更改时间
	ZhandianXX    *ZhandianXX    `orm:"rel(fk);null;on_delete(do_nothing)"`                                   //设置一对一关系
	CheliangXXs   []*CheliangXX  `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	CheliangZTs   []*CheliangZT  `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	GonggaoXXs    []*GonggaoXX   `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	GonggaoZTs    []*GonggaoZT   `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	Guojias       []*Guojia      `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	Shengfens     []*Shengfen    `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	Chengshis     []*Chengshi    `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	Quxians       []*Quxian      `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	Jiedaos       []*Jiedao      `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	MokuaiPZs     []*MokuaiPZ    `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	MokuaiZs      []*MokuaiZ     `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	QishouJJLXRs  []*QishouJJLXR `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	QishouXXs     []*QishouXX    `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	QishouZTs     []*QishouZT    `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	QishouDZXXs   []*QishouDZXX  `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	QuanxianPZs   []*QuanxianPZ  `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	QuanxianZs    []*QuanxianZ   `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	RenjiGXs      []*RenjiGX     `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	SusheXXs      []*SusheXX     `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	ZhandianLXs   []*ZhandianLX  `orm:"reverse(many)"`                                                        //设置一对多的反向关系
	ZhandianXXs   []*ZhandianXX  `orm:"reverse(many)"`                                                        //设置一对多的反向关系
}

func init() {
	orm.RegisterModel(new(DengluYH))
}

//Table Name
func (u *DengluYH) TableName() string {
	return "dengluYH"
}

// AddDengluYH insert a new DengluYH into database and returns last inserted Id on success.
func AddDengluYH(m *DengluYH, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Insert(m)
	return
}

// DeleteDengluYHByID delete by ID
func DeleteDengluYHByID(id string, ormObj orm.Ormer) (num int64, err error) {
	obj := &DengluYH{ID: id}
	num, err = ormObj.Delete(obj)
	return
}

// UpdateDengluYH update DengluYH into database and returns id on success
func UpdateDengluYH(m *DengluYH, ormObj orm.Ormer) (id int64, err error) {
	id, err = ormObj.Update(m)
	return
}

// GetDengluYHByID retrieves DengluYH by ID. Returns error if ID doesn't exist
func GetDengluYHByID(id string, ormObj orm.Ormer) (obj *DengluYH, err error) {
	obj = &DengluYH{ID: id}
	err = ormObj.Read(obj)
	return obj, err
}

// GetAllDengluYH retrieves all DengluYH matches certain condition. Returns empty list if no records exist
func GetAllDengluYH(o orm.Ormer, query map[string]interface{}, exclude map[string]interface{}, condMap map[string]map[string]interface{},
	fields []string, sortby []string, order []string, offset int64, limit int64) (utils.Paginator, []DengluYH, error) {
	var (
		objArrs   []DengluYH
		err       error
		paginator utils.Paginator
		num       int64
	)
	qs := o.QueryTable(new(DengluYH))
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
