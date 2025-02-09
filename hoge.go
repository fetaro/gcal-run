package main

import (
	"fmt"
	"github.com/fetaro/gcal_forcerun_go/lib/common"
	"os/exec"
)

func main() {
	fmt.Println(common.GetAppDir())
	//err := exec.Command("cmd", "/c", "start", common.GetAppDir()).Run()
	//if err != nil {
	//	panic(err)
	//}
	out, err := exec.Command("cmd", "/c", "powershell.exe", "resource\\register_startup.ps1").Output()
	fmt.Println(string(out))
	if err != nil {
		panic(err)
	}
	out, err = exec.Command("cmd", "/c", "powershell.exe", "resource\\make_desktop_shortcut.ps1").Output()
	fmt.Println(string(out))
	if err != nil {
		panic(err)
	}
}
