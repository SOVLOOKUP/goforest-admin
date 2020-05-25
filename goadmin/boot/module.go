package boot

import (
	"goadmin/module/home"
	"goadmin/module/public"
	"goadmin/module/sys"
	"goadmin/module/ai"
)

func InitModules() {
	public.InitModule()
	home.InitModule()
	sys.InitModule()
	ai.InitModule()
}
