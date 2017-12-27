package main

func main() {
	center := newCenterServer()

	center.proc.StartProcess() //blocked here
}
