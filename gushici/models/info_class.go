/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-14 15:24:51
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 11:48:52
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
)

type InfoClass struct {
	Id        int
	ClassName string
	Linkurl   string
	Desc      string
	Orderid   int
	Status    int
}

func (a *InfoClass) TableName() string {
	return TableName("info_class")
}

func ClassGetList(page, pageSize int, filters ...interface{}) ([]*InfoClass, int64) {
	offset := (page - 1) * pageSize
	list := make([]*InfoClass, 0)
	query := orm.NewOrm().QueryTable(TableName("info_class"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}

func ClassAdd(role *InfoClass) (int64, error) {
	id, err := orm.NewOrm().Insert(role)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func ClassGetById(id int) (*InfoClass, error) {
	r := new(InfoClass)
	err := orm.NewOrm().QueryTable(TableName("info_class")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *InfoClass) ClassUpdate(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}
