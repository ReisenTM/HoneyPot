package flags

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"honeypot_node/internal/global"
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
	flag.BoolVar(&Options.Version, "vv", false, "打印当前版本")
	flag.BoolVar(&Options.Help, "h", false, "帮助信息")
	flag.StringVar(&Options.Value, "v", "", "值")
	flag.Parse()
	// 注册命令
	RegisterCommand()
}

func RegisterCommand() {
}

func runBaseCommand() {

	if Options.Version {
		logrus.Infof("当前版本: %s\n", global.Version)
		os.Exit(0)
	}
}
func runHelpCommand() {
	if Options.Menu == "" && Options.Type == "" && Options.Help {
		fmt.Printf("菜单项:\n")
		for key, _ := range HelpCommandMap {
			fmt.Printf("%s 使用 -m %s -h 查看具体子菜单\n", key, key)
		}
		os.Exit(0)
	}
	if Options.Menu != "" && Options.Type == "" && Options.Help {
		subMenuMap, ok := HelpCommandMap[Options.Menu]
		if !ok {
			logrus.Fatalf("不存在的菜单项 %s", Options.Menu)
		}
		for key, help := range subMenuMap {
			fmt.Printf("%s %s\n", key, help)
		}
		os.Exit(0)
	}

}
func runCommand() {
	if Options.Menu == "" || Options.Type == "" {
		return
	}
	key := fmt.Sprintf("%s:%s", Options.Menu, Options.Type)
	command, ok := CommandMap[key]
	if !ok {
		logrus.Fatalf("不存在的菜单项 %s %s", Options.Menu, Options.Type)
	}
	command.Func()
	os.Exit(0)
}

type Command struct {
	Menu string
	Type string
	Help string
	Func func()
}

var CommandMap = map[string]*Command{}
var HelpCommandMap = map[string]map[string]string{}

func registerCommand(menu, subMenu, help string, fun func()) {
	key := fmt.Sprintf("%s:%s", menu, subMenu)
	CommandMap[key] = &Command{
		Menu: menu,
		Type: subMenu,
		Help: help,
		Func: fun,
	}
	subMenuMap, ok := HelpCommandMap[menu]
	if ok {
		subMenuMap[subMenu] = help
	} else {
		HelpCommandMap[menu] = map[string]string{
			subMenu: help,
		}
	}
}

func Run() {
	// 运行基本命令
	runBaseCommand()
	// 运行帮助的命令
	runHelpCommand()
	// 运行注册的命令
	runCommand()
}
