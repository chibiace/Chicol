package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	// Command line args
	cBasicFlag := flag.String("c", "", "valid colours: black, red, green, yellow, blue, magenta, cyan, white")
	cHexFlag := flag.String("x", "", "valid colours: any 6 digit hex code, eg. '#FFAA00'")
	cRandomFlag := flag.Bool("r", false, "enable random colours (default basic mode)")
	cRainbowFlag := flag.Bool("rainbow", false, "r a i n b o w s :  everybody loves rainbows")
	cRainbowStepFlag := flag.Int("s", 1, "rainbow step size")
	cRandomTypeFlag := flag.String("t", "basic", "random mode: basic or range")
	cRandomMinFlag := flag.Int("min", 0, "range minimum value (0-255), also applies to rainbow (default 0)")
	cRandomMaxFlag := flag.Int("max", 255, "range maximum value (0-255), also applies to rainbow")
	flag.Parse()

	// Pipework here
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("No input provided. Exiting.")
		os.Exit(0)
	}
	piped, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	pipeStr := string(piped)
	pipeStr = strings.TrimSuffix(pipeStr, "\n")

	// Default basic colours
	basicColours := map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}

	// The mess
	if *cBasicFlag != "" {
		fmt.Println(basicColour(pipeStr, *cBasicFlag, basicColours))
		return
	} else if *cHexFlag != "" {
		fmt.Println(hexColour(pipeStr, *cHexFlag))
		return
	} else if *cRainbowFlag {
		if *cRandomMinFlag < 0 || *cRandomMaxFlag > 255 {
			fmt.Println("bad range")
			os.Exit(1)
		} else {
			fmt.Println(rainbow(pipeStr, *cRainbowStepFlag, *cRandomMinFlag, *cRandomMaxFlag))
			return
		}
	} else if *cRandomFlag {
		if *cRandomTypeFlag == "basic" {
			fmt.Println(randomBasic(pipeStr, basicColours))
			return
		} else if *cRandomTypeFlag == "range" {
			if *cRandomMinFlag < 0 || *cRandomMaxFlag > 255 {
				fmt.Println("bad range")
				os.Exit(1)
			} else {
				fmt.Println(randomRange(pipeStr, *cRandomMinFlag, *cRandomMaxFlag))
				return
			}
		} else {
			fmt.Println("bad random type")
		}
	} else {
		fmt.Println("choose a mode and or colour")
		os.Exit(1)
	}
}

// Functions

func basicColour(s string, colour string, basicColours map[string]string) string {
	colour = strings.ToLower(colour)
	return fmt.Sprintf("%s%s\033[0m", basicColours[colour], s)
}

func hexColour(s string, hex string) string {
	hex = strings.TrimPrefix(hex, "#")
	if len(s) > 6 {
		hex = hex[:6]
	}
	if len(hex) != 6 {
		panic(errors.New("hex length wrong"))
	}
	r, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		panic(err)
	}
	g, err := strconv.ParseUint(hex[2:4], 16, 8)
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseUint(hex[4:6], 16, 8)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", uint8(r), uint8(g), uint8(b), s)
}

func randomBasic(s string, basicColours map[string]string) string {
	var colours []string
	for _, v := range basicColours {
		colours = append(colours, v)
	}
	c := CUgetRandomElement(colours)
	return fmt.Sprintf("%v%s\033[0m", c, s)
}

func rainbow(s string, step int, minimum int, maximum int) string {
	if step < 1 {
		step = 1
	}
	r := maximum
	g := minimum
	b := minimum
	newStr := ""
	for _, v := range s {

		if r == maximum && g < maximum {
			g += step
			if g > maximum {
				g = maximum
			}
		} else if r > minimum && g == maximum {
			r -= step
			if r < minimum {
				r = minimum
			}
		} else if g == maximum && b < maximum {
			b += step
			if b > maximum {
				b = maximum
			}
		} else if g > minimum && b == maximum {
			g -= step
			if g < minimum {
				g = minimum
			}
		} else if r < maximum && b == maximum {
			r += step
			if r > maximum {
				r = maximum
			}
		} else if r == maximum && b > minimum {
			b -= step
			if b < minimum {
				b = minimum
			}
		}

		newStr += fmt.Sprintf("\033[38;2;%d;%d;%dm%v\033[0m", r, g, b, string(v))
	}
	return newStr
}

func randomRange(s string, minimum int, maximum int) string {
	r := CURandomRange(minimum, maximum)
	g := CURandomRange(minimum, maximum)
	b := CURandomRange(minimum, maximum)
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%v\033[0m", r, g, b, s)
}

func CURandomRange(a int, b int) int {
	return rand.Intn(b-a+1) + a
}

func CUgetRandomElement(slice []string) string {
	index := rand.Intn(len(slice))
	return slice[index]
}
