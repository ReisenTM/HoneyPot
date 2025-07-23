package main

import (
	"Honeypot/apps/honeypot_server/core"
	"Honeypot/apps/honeypot_server/flags"
	"Honeypot/apps/honeypot_server/global"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	global.Config = core.Init_Config()
	global.DB = core.InitDB()
	flags.Run()
}
