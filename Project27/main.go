package main

import (
	"fmt"
)

type Phone interface {
	info()
}

func info() {
	fmt.Println("NoPhone")
}

type iPhone struct {
	Version string
	Money   int
	Release int
}

type SamsungPhone struct {
	Version string
	Money   int
	Release int
}

type MiPhone struct {
	Version string
	Money   int
	Release int
}

func (ver iPhone) info() {
	fmt.Printf("This is iPhone, The Version is %s, The Money is %d,The Release is %d", ver.Version, ver.Money, ver.Release)
}
func (ver MiPhone) info() {
	fmt.Printf("This is MiPhone, The Version is %s, The Money is %d,The Release is %d", ver.Version, ver.Money, ver.Release)
}
func (ver SamsungPhone) info() {
	fmt.Printf("This is SamsungPhone, The Version is %s, The Money is %d,The Release is %d", ver.Version, ver.Money, ver.Release)
}

var (
	iphone12     iPhone       = iPhone{Version: "12", Money: 5500, Release: 2020}
	miphone      MiPhone      = MiPhone{Version: "RedMiK30U", Money: 2000, Release: 2020}
	samsungphone SamsungPhone = SamsungPhone{Version: "???", Money: 10000, Release: 2020}
)

func info2(input Phone) {
	input.info()
}

func main() {
	/* 	var input string
	   	fmt.Scanln(&input)
	   	if input == "i" {
	   		info2(iphone12)
	   	} else if input == "m" {
	   		info2(miphone)
	   	} else {
	   		info2(samsungphone)
	   	} */
	var phone1 Phone
	phone1.info()
}
