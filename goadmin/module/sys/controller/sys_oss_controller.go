/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  sys_oss_controller
* @Version: 1.0.0
* @Date: 2020-05-25
*/

package controller


import (
   "goadmin/utils/base"
   "goadmin/module/sys/model"
   "github.com/gogf/gf/frame/g"
   "github.com/gogf/gf/net/ghttp"
   "github.com/gogf/gf/os/glog"
   "github.com/gogf/gf/util/gconv"
)

type SysOssController struct {
    base.BaseRouter
}

var (
    controllerNameSysOss = "SysOssController"
)

//SysOss页面信息
func (controller *SysOssController) Index(r *ghttp.Request) {
    base.WriteTpl(r, "sys/oss.html", g.Map{})
}

//获取SysOss单条信息
func (controller *SysOssController) Get(r *ghttp.Request) {
    id := r.GetInt("id")
    model := model.SysOss{Id: id}.Get()
    if model.Id <= 0 {
       base.Fail(r, controllerNameSysOss+" get fail")
    }
    base.Succ(r, model)
}

//根据id或者ids删除SysOss
func (controller *SysOssController) Delete(r *ghttp.Request) {
   ids := r.GetInts("ids")
    for _, id := range ids {
        model := model.SysOss{Id: id}
        model.Delete()
    }
    base.Succ(r, "")
}


//创建一条SysOss
func (controller *SysOssController) Save(r *ghttp.Request) {
    model := model.SysOss{}
    err := gconv.Struct(r.GetPostMap(), &model)
    if err != nil {
      glog.Error( controllerNameSysOss+" save struct error", err)
      base.Error(r, "save error")
    }
    var num int64
    if model.Id <= 0 {
      num = model.Insert()
    } else {
      num = model.Update()
    }

    if num <= 0 {
       base.Fail(r, controllerNameSysOss+" save fail")
    }

    base.Succ(r, "")
}

//更新一条SysOss
func (controller *SysOssController) Update(r *ghttp.Request) {
    model := model.SysOss{}
    err := gconv.Struct(r.GetPostMap(), &model)
    if err != nil {
    glog.Error( controllerNameSysOss+" save struct error", err)
    base.Error(r, "save error")
    }
    var num int64
    if model.Id <= 0 {
    num = model.Insert()
    } else {
    num = model.Update()
    }

    if num <= 0 {
    base.Fail(r, controllerNameSysOss+" save fail")
    }

    base.Succ(r, "")
}

//分页列表SysOss
func (controller *SysOssController) Page(r *ghttp.Request) {
   form := base.NewForm(r.GetQueryMap())
   model := model.SysOss{}
   page := model.Page(&form)
   base.Succ(r, g.Map{"list": page, "form": form})
}