package Utils

import (
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr,"Fatal error: %s",err.Error())
		//fmt.Println(err)
		os.Exit(1)
	}
}
