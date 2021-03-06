/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  animals_model
* @Version: 1.0.0
* @Date: 2020-05-25
*/
package model

import (
    "github.com/gogf/gf/util/gconv"
    "goadmin/utils/base"
    "github.com/gogf/gf/database/gdb"
    "github.com/gogf/gf/frame/g"
    "github.com/gogf/gf/os/glog"
)

type Animals struct {
    Name         string  `orm:"name"          json:"Name,omitempty" gconv:"Name,omitempty"`                  
    Altername    string  `orm:"altername"     json:"Altername,omitempty" gconv:"Altername,omitempty"`        
    Latinname    string  `orm:"latinname"     json:"Latinname,omitempty" gconv:"Latinname,omitempty"`        
    Englishname  string  `orm:"englishname"   json:"Englishname,omitempty" gconv:"Englishname,omitempty"`    
    Ordermu      string  `orm:"ordermu"       json:"Ordermu,omitempty" gconv:"Ordermu,omitempty"`            
    Family       string  `orm:"family"        json:"Family,omitempty" gconv:"Family,omitempty"`              
    Genus        string  `orm:"genus"         json:"Genus,omitempty" gconv:"Genus,omitempty"`                
    Home         string  `orm:"home"          json:"Home,omitempty" gconv:"Home,omitempty"`                  
    Details      string  `orm:"details"       json:"Details,omitempty" gconv:"Details,omitempty"`            
    Price        float64 `orm:"price"         json:"Price,omitempty" gconv:"Price,omitempty"`                
    Laws         string  `orm:"laws"          json:"Laws,omitempty" gconv:"Laws,omitempty"`                  
    Images       string  `orm:"images"        json:"Images,omitempty" gconv:"Images,omitempty"`              
    ProtectLevel string  `orm:"protect_level" json:"ProtectLevel,omitempty" gconv:"ProtectLevel,omitempty"`  
    Id           int     `orm:"id,primary"    json:"Id,omitempty" gconv:"Id,omitempty"`                      
}

//创建mAnimals
func (model *Animals) Insert() int64 {
    //model.CreateTime = gtime.Now()
    //model.UpdateTime = gtime.Now()
    //model.Status=1
    r, err := model.dbModel().Data(model).Insert()
    if err != nil {
        glog.Error(model.TableName()+" insert error", err)
        return 0
    }

    res, err2 := r.RowsAffected()
    if err2 != nil {
        glog.Error(model.TableName()+" insert res error", err2)
        return 0
    } else if res > 0 {
        lastId, err2 := r.LastInsertId()
        if err2 != nil {
            glog.Error(model.TableName()+" LastInsertId res error", err2)
            return 0
        } else {
           model.Id = gconv.Int(lastId)
        }
    }
    return res
}
//删除Animals
func (model Animals) Delete() int64 {
    if model.Id <= 0 {
        glog.Error(model.TableName() + " delete id error")
        return 0
    }
    r, err := model.dbModel().Where(" id = ?", model.Id).Delete()
    if err != nil {
        glog.Error(model.TableName()+" delete error", err)
        return 0
    }
    res, err2 := r.RowsAffected()
    if err2 != nil {
        glog.Error(model.TableName()+" delete res error", err2)
        return 0
    }
    return res
}

//更新Animals
func (model *Animals) Update() int64 {
    //model.UpdateTime = gtime.Now();
    r, err := model.dbModel().Data(model).Where(" id = ?", model.Id).Update()
    if err != nil {
        glog.Error(model.TableName()+" update error", err)
        return 0
    }
    res, err2 := r.RowsAffected()
    if err2 != nil {
        glog.Error(model.TableName()+" update res error", err2)
        return 0
    }
    return res
}

//根据ID查询Animals
func (model Animals) Get() Animals {
    if model.Id <= 0 {
        glog.Error(model.TableName() + " get id error")
        return Animals{}
    }
    var resData Animals
    err := model.dbModel("t").Where(" id = ?", model.Id).Fields(model.columns()).Struct(&resData)
    if err != nil {
        glog.Error(model.TableName()+" get one error", err)
        return Animals{}
    }
    return resData
}

//分页查询Animals
func (model Animals) Page(form *base.BaseForm) []Animals {
    if form.Page <= 0 || form.Rows <= 0 {
        glog.Error(model.TableName()+" page param error", form.Page, form.Rows)
        return []Animals{}
    }
    where := ""
    var params []interface{}
    if form.Params != nil && form.Params["name"] != "" {
        where += " and name like ? "
        params = append(params, "%"+form.Params["name"]+"%")
    }

    num, err := model.dbModel("t").Where(where, params).Count()
    form.TotalSize = num
    form.TotalPage = num / form.Rows
    
    // 没有数据直接返回
    if num == 0 {
        form.TotalPage = 0
        form.TotalSize = 0
        return []Animals{}
    }
    var resData []Animals
    pageNum, pageSize := (form.Page-1)*form.Rows, form.Rows
    dbModel := model.dbModel("t").Fields(model.columns()).Fields(model.columns())
    err = dbModel.Where(where, params).Limit(pageNum, pageSize).OrderBy(form.OrderBy).Structs(&resData)
    if err != nil {
        glog.Error(model.TableName()+" page list error", err)
        return []Animals{}
    }

    return resData
}

//返回数据库Animals
func (model Animals) dbModel(alias ...string) *gdb.Model {
    var tmpAlias string
    if len(alias) > 0 {
    tmpAlias = " " + alias[0]
    }
    tableModel := g.DB("center").Table(model.TableName() + tmpAlias).Safe()
    return tableModel
}
//返回主键Animals
func (model Animals) PkVal() int {
    return model.Id
}
//表名Animals
func (model Animals) TableName() string {
    return "animals"
}
//列名Animals
func (model Animals) columns() string {
    sqlColumns := "t.home as Home,t.price as Price,t.protect_level as ProtectLevel,t.latinname as Latinname,t.family as Family,t.genus as Genus,t.laws as Laws,t.images as Images,t.name as Name,t.englishname as Englishname,t.ordermu as Ordermu,t.altername as Altername,t.details as Details,t.id as Id"
    return sqlColumns
}