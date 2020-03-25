package main


func main() {
	config := get_config("/Users/swayam.raina/opensrc/fcd/test.yaml")
	do_setup(&config)
	daemon(&config)
}