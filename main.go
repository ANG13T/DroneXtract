package main

// OSINT, DAT File parsing, other file format parsing, crystal web server, TUI and GUI

import "fmt"

func main() {
	fmt.Println("hello world")
	f, err := os.Open("test.DAT")
	check(err)
	s := bufio.NewScanner(f)
	count := 0
	for s.Scan() {
		line := s.Text()
		if count < 1 {
			count, err = strconv.Atoi(line)
			check(err)
			continue
		}
		count--
		var tag string
		var n int
		var f float64
		fmt.Sscanf(line, "%s %d %f", &tag, &n, &f)
		// not sure what you really wnant to do with the data!
		fmt.Println(n, f, tag)
	}
}