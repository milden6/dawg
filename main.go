package main

import (
	"fmt"

	"github.com/milden6/dawg/dawg-dict"
)

func main() {
	d := dawg.NewDawg()
	d.Add("tuc;tuc;NOUN;0")
	d.Add("tucatuc;tuc;NOUN;0")
	d.Add("tucatuca;tuc;NOUN;0")
	d.Add("tucatucami;tuc;NOUN;0")
	d.Add("tucatuci;tuc;NOUN;0")
	d.Add("tucatuh;tuc;NOUN;0")

	finder := d.Finish()

	for _, word := range finder.FindAllByPrefix("tuc") {
		fmt.Println(word)
	}
}
