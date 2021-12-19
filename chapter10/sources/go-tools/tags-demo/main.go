package main

import "fmt"

type featureAttr struct {
	name string
	ver  string // community, professional, ultimate
}

var featureList []featureAttr

func init() {
	featureList = append(featureList, featureAttr{
		name: "a",
		ver:  "community",
	})
	featureList = append(featureList, featureAttr{
		name: "b",
		ver:  "community",
	})
}

func dumpFeatures() {
	fmt.Println("Features list:")
	for i, f := range featureList {
		fmt.Printf("\t%d: %s (%s)\n", i+1, f.name, f.ver)
	}
}

func main() {
	dumpFeatures()
}
