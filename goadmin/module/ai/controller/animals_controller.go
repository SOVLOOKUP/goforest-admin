/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  animals_controller
* @Version: 1.0.0
* @Date: 2020-05-25
*/

package controller


import (
   "goadmin/utils/base"
   "goadmin/module/ai/model"
   "github.com/gogf/gf/frame/g"
   "github.com/gogf/gf/net/ghttp"
   "github.com/gogf/gf/os/glog"
   "github.com/gogf/gf/util/gconv"
)

type AnimalsController struct {
    base.BaseRouter
}

var (
    controllerNameAnimals = "AnimalsController"
)

//Animals页面信息
func (controller *AnimalsController) Index(r *ghttp.Request) {
    base.WriteTpl(r, "animals.html", g.Map{})
}

//获取Animals单条信息
func (controller *AnimalsController) Get(r *ghttp.Request) {
    id := r.GetInt("id")
    model := model.Animals{Id: id}.Get()
    if model.Id <= 0 {
       base.Fail(r, controllerNameAnimals+" get fail")
    }
    base.Succ(r, model)
}

//根据id或者ids删除Animals
func (controller *AnimalsController) Delete(r *ghttp.Request) {
   ids := r.GetInts("ids")
    for _, id := range ids {
        model := model.Animals{Id: id}
        model.Delete()
    }
    base.Succ(r, "")
}


//创建一条Animals
func (controller *AnimalsController) Save(r *ghttp.Request) {
    model := model.Animals{}
    err := gconv.Struct(r.GetPostMap(), &model)
    if err != nil {
      glog.Error( controllerNameAnimals+" save struct error", err)
      base.Error(r, "save error")
    }
    var num int64
    if model.Id <= 0 {
      num = model.Insert()
    } else {
      num = model.Update()
    }

    if num <= 0 {
       base.Fail(r, controllerNameAnimals+" save fail")
    }

    base.Succ(r, "")
}

//更新一条Animals
func (controller *AnimalsController) Update(r *ghttp.Request) {
    model := model.Animals{}
    err := gconv.Struct(r.GetPostMap(), &model)
    if err != nil {
    glog.Error( controllerNameAnimals+" save struct error", err)
    base.Error(r, "save error")
    }
    var num int64
    if model.Id <= 0 {
    num = model.Insert()
    } else {
    num = model.Update()
    }

    if num <= 0 {
    base.Fail(r, controllerNameAnimals+" save fail")
    }

    base.Succ(r, "")
}

//分页列表Animals
func (controller *AnimalsController) Page(r *ghttp.Request) {
   form := base.NewForm(r.GetQueryMap())
   model := model.Animals{}
   page := model.Page(&form)
   base.Succ(r, g.Map{"list": page, "form": form})
}