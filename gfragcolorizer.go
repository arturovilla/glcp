package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	wordPtr := flag.String("values", "[[0.500 0.500 0.500] [0.500 0.500 0.500] [1.000 1.000 1.000] [0.000 0.333 0.667]]", "a string")
	flag.Parse()
	processedValues := *wordPtr
	if len(processedValues) > 2 {
		valuesProcessed := strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(processedValues[1:len(processedValues)-1], " ", ","), "[", ""), "]", "")

		var colorList = strings.Split(valuesProcessed, ",")
		fmt.Println("\n\t**********   COLOR PALLETE FUNCTION   **********\n")
		fmt.Println("vec3 pallete(float t){")
		fmt.Printf("\tvec3 a = vec3(%s, %s, %s);\n", colorList[0], colorList[1], colorList[2])
		fmt.Printf("\tvec3 b = vec3(%s, %s, %s);\n", colorList[3], colorList[4], colorList[5])
		fmt.Printf("\tvec3 c = vec3(%s, %s, %s);\n", colorList[6], colorList[7], colorList[8])
		fmt.Printf("\tvec3 d = vec3(%s, %s, %s);\n", colorList[9], colorList[10], colorList[11])
		fmt.Println("\treturn a + b*cos(6.28318*(c+t+d));")
		fmt.Println("}\n")

	} else {
		fmt.Println("Invalid input format")
	}
}
