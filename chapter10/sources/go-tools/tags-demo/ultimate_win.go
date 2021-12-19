// +build ultimate,win

package main

func init() {
	featureList = append(featureList, featureAttr{
		name: "g",
		ver:  "ultimate(windows)",
	})
	featureList = append(featureList, featureAttr{
		name: "h",
		ver:  "ultimate(windows)",
	})
}
