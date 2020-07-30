package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {

	// Handle cli arguments
	infile := flag.String("in", "", "Input config file")
	outfile := flag.String("out", "", "Output config file")
	flag.Parse()

	if *infile == "" {
		messageReturn("No input config file provided", 3)
	}
	if *outfile == "" {
		messageReturn("No output config file provided", 3)
	}

	var inbuff string

	filecontent := readConfig(*infile)
	ifaces := detectIfaces(filecontent)

	fmt.Println("Detected the following interfaces in input config:")
	for _, element := range ifaces {
		fmt.Println("[" + element + "]")
	}
	fmt.Print("Is this correct? [y/N] ")

	fmt.Scanln(&inbuff)

	if inbuff == "n" || inbuff == "N" {
		messageReturn("I'm not sure why all interfaces could not be detected, sorry. :/", 0)
	} else if inbuff != "y" && inbuff != "Y" {
		messageReturn("Invalid option", 3)
	}

	var newifaces []string

	fmt.Println("Please type the new interface names to replace the existing ones.")

	for _, element := range ifaces {
		var b string
		fmt.Print("Which interface will replace [" + element + "]?: ")
		fmt.Scanln(&b)
		newifaces = append(newifaces, b)
	}

	fmt.Println("Please confirm that this is correct:")

	for i := range ifaces {
		fmt.Println("[" + ifaces[i] + "] -> [" + newifaces[i] + "]")
	}
	fmt.Print("Is this correct? [y/N] ")

	fmt.Scanln(&inbuff)

	if inbuff == "n" || inbuff == "N" {
		messageReturn("Aborting!", 0)
	} else if inbuff != "y" && inbuff != "Y" {
		messageReturn("Invalid option", 3)
	}

	replaceIfaces(*outfile, filecontent, ifaces, newifaces)
}

func readConfig(file string) []byte {
	filecontent, err := ioutil.ReadFile(file)
	if err != nil {
		messageReturn("File does not exist or otherwise could not be read.", 3)
	}

	return filecontent
}

// detectIfaces returns string array with each interface found in the config
func detectIfaces(filecontent []byte) []string {
	var ifaces []string
	var re = regexp.MustCompile(`(<if>)([a-zA-Z0-9.-]+)?`)

	lines := strings.Split(string(filecontent), "\n")

	for _, line := range lines {
		iface := re.FindStringSubmatch(line)

		if len(iface) > 1 && iface[2] != "" {
			ifaces = append(ifaces, iface[2])
		}
	}

	ifaces = dedupe(ifaces)
	sort.Strings(ifaces)

	return ifaces
}

// replaceIfaces finds old ifaces and replaces them
func replaceIfaces(outfile string, filecontent []byte, oldifaces []string, newifaces []string) {
	lines := strings.Split(string(filecontent), "\n")

	for i := range oldifaces {
		var re = regexp.MustCompile(`<if>` + oldifaces[i] + `</if>`)

		for j, line := range lines {
			lines[j] = re.ReplaceAllString(line, `<if>`+newifaces[i]+`</if>`)
		}
	}

	output := strings.Join(lines, "\n")
	err := ioutil.WriteFile(outfile, []byte(output), 0644)
	if err != nil {
		messageReturn("Could not write output file.", 3)
	}
	messageReturn("Altered config written to "+outfile, 0)
}

// messageReturn prints a string and returns a status
func messageReturn(text string, code int) {
	fmt.Println(text)
	os.Exit(code)
}

// dedupe removes duplicate strings in a slice
func dedupe(s []string) []string {
	encountered := map[string]bool{}

	for v := range s {
		encountered[s[v]] = true
	}

	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}
