/**********************************************
** @Des: This file ...
** @Author: haodaquan
** @Date:   2017-09-16 15:42:43
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-24 11:48:17
***********************************************/
package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type InfoList struct {
	Id         int
	ClassId    int
	Title      string
	Keywords   string
	Content    string
	Desc       string
	Author     string
	Used     	int
	Picurl     string
	Linkurl     string
	Media      string
	Orderid    int
	Hits       int
	Status     int
	Posttime   int64
	Updatetime int64
}

func (a *InfoList) TableName() string {
	return TableName("info_list")
}

func NewsAdd(a *InfoList) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func NewsGetList(page, pageSize int, filters ...interface{}) ([]*InfoList, int64) {
	offset := (page - 1) * pageSize
	list := make([]*InfoList, 0)
	query := orm.NewOrm().QueryTable(TableName("info_list"))
	if len(filters) > 0 {
		l := len(filters)

		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-orderid", "-id").Limit(pageSize, offset).All(&list)

	return list, total
}

func NewsGetListForWeiXin(page, pageSize int, filters ...interface{}) ([]*InfoList, int64) {
	offset := (page - 1) * pageSize
	list := make([]*InfoList, 0)
	query := orm.NewOrm().QueryTable(TableName("info_list"))
	if len(filters) > 0 {
		l := len(filters)

		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	query = query.Filter("class_id__in", 5,3)
	total, _ := query.Count()
	query.OrderBy("-orderid", "-id").Limit(pageSize, offset).All(&list)

	return list, total
}

func NewsSearchList(page, pageSize int, keywords string) ([]*InfoList, interface{})   {
	list := make([]*InfoList, 0)

	offset := (page - 1) * pageSize

	o := orm.NewOrm()

	sql := fmt.Sprintf("select * from " + TableName("info_list") +
		" where status=1 and class_id in(5,3) and (title like '%s' or author like '%s') order by id desc limit %d , %d", "%"+keywords+"%", "%"+keywords+"%", offset, pageSize)


	sqlTotal := fmt.Sprintf("select count(*) as c_num  from " + TableName("info_list") +  " where status=1 and class_id in(5,3) " +
		" and (title like '%s' or author like '%s')", "%"+keywords+"%", "%"+keywords+"%")

	var c_num string
	o.Raw(sqlTotal).QueryRow(&c_num)

	_, er := o.Raw(sql).QueryRows(&list)
	if er != nil {

	}

	return list, c_num
}

func NewsGetById(id int) (*InfoList, error) {
	r := new(InfoList)
	err := orm.NewOrm().QueryTable(TableName("info_list")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func NewsGetNextById(id int) (*InfoList, error) {

	var maps = new(InfoList)

	o := orm.NewOrm()

	sql := fmt.Sprintf("select id,title from " + TableName("info_list") +
		" where status=1 and id < %d order by id desc limit 1", id)
	//fmt.Println(sql)

	err := o.Raw(sql).QueryRow(&maps)

	if err != nil {
		return nil, err
	}
	return maps, nil
}

func (a *InfoList) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
