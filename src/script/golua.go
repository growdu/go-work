package script


import (
	"github.com/yuin/gopher-lua"
)

func Start(){
	L:=lua.NewState()
	defer L.Close()

	if err:=L.DoString(`print("hello")`);err!=nil{
		panic(err)
	}

}
