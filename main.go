package main

import (
	"fmt"
)

//practice cmd

// invoker

type Button struct {
	cmd Command
}

func (b *Button) press() {
	b.cmd.execute()
}

type Command interface {
	execute()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// receiver func
type Device interface {
	on()
	off()
}

type TV struct {
	isRunning bool
}

func (t *TV) on() {
	t.isRunning = true
	fmt.Println("tv is on")

}

func (t *TV) off() {
	t.isRunning = false
	fmt.Println("tv is off")

}

//oncmd

//offcmd

func main() {

	//

	tv := &TV{}

	onCmd := &OnCommand{device: tv}
	onButton := &Button{cmd: onCmd}
	onButton.press()
	onButton.press()

	fmt.Println("hello coochiku")
}
