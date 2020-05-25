/**
* @Author: Rocks
* @Email: 451360994@qq.com
* @Description:
* @File:  module
* @Version: 1.0.0
* @Date: 2020-05-25
*/
package sys

import "goadmin/module/sys/config"

func InitModule()  {
    config.InitRouter()
}