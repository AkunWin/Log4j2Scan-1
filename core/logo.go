package core

import "fmt"

func PrintLogo(authors []interface{}) {
	template := "    __                __ __  _ _____                \n" +
		"   / /   ____  ____ _/ // / (_) ___/_________ _____ \n  " +
		"/ /   / __ \\/ __ `/ // /_/ /\\__ \\/ ___/ __ `/ __ \\\n " +
		"/ /___/ /_/ / /_/ /__  __/ /___/ / /__/ /_/ / / / /\n" +
		"/_____/\\____/\\__, /  /_/_/ //____/\\___/\\__,_/_/ /_/ \n " +
		"/____/    /___/                         \n" +
		"    coded by %s & %s"
	logo := fmt.Sprintf(template, authors...)
	fmt.Println(logo)
}
