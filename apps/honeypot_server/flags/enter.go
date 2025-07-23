package flags

import (
	"flag"
	"os"
)

type FlagOptions struct {
	File    string
	Version bool
	DB      bool
	Menu    string
	Type    string
	Value   string
	Help    bool
}

var Options FlagOptions

func init() {
	flag.StringVar(&Options.File, "f", "settings.yaml", "配置文件路径")
	//flag.BoolVar(&Options.Version, "vv", false, "打印当前版本")
	//flag.BoolVar(&Options.Help, "h", false, "帮助信息")
	flag.BoolVar(&Options.DB, "db", false, "迁移表结构")
	//flag.StringVar(&Options.Menu, "m", "", "菜单 user")
	//flag.StringVar(&Options.Type, "t", "", "类型 create list")
	//flag.StringVar(&Options.Value, "v", "", "值")
	flag.Parse()
	// 注册命令
}

func Run() {
	if Options.DB {
		Migrate()
		os.Exit(0)
	}
}
