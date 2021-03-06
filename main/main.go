package main

import (
	"fmt"
	. "github.com/Jxck/color"
)

func main() {
	fmt.Println("name    code  (ansi name)")
	fmt.Println("-------------------------")
	fmt.Println(Black("Black   0;30m (black)"))
	fmt.Println(Red("Red     0;31m (red)"))
	fmt.Println(Green("Green   0;32m (green)"))
	fmt.Println(Brown("Brown   0;33m (brown)"))
	fmt.Println(Navy("Navy    0;34m (blue)"))
	fmt.Println(Purple("Purple  0;35m (purple)"))
	fmt.Println(Cyan("Cyan    0;36m (cyan)"))
	fmt.Println(Gray("Gray    0;37m (light Gray)"))
	fmt.Println(Dim("Dim     1;30m (dark Gray)"))
	fmt.Println(Orange("Orange  1;31m (light Red)"))
	fmt.Println(Lime("Lime    1;32m (light Green)"))
	fmt.Println(Yellow("Yellow  1;33m (yellow)"))
	fmt.Println(Blue("Blue    1;34m (light Blue)"))
	fmt.Println(Pink("Pink    1;35m (light Purple)"))
	fmt.Println(Aqua("Aqua    1;36m (light Cyan)"))
	fmt.Println(White("White   1;37m (white)"))
}
