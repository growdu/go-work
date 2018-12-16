package main

import (
   // "mymath"
   "fmt"
    "runtime"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"
)


func say(s string){
    for i:=0;i<5;i++ {
        runtime.Gosched()
        fmt.Println(s)
    }
}

func sum(a []int,c chan int){
    total:=0
    for _,v:=range a{
        total+=v
    }
    c <- total
}

func getHttpPage(){
    resp,err :=http.Get("http://www.baidu.com")
    if err !=nil{
        fmt.Println("http get error!")
    }
    defer resp.Body.Close()
    body,err :=ioutil.ReadAll(resp.Body)
    if err !=nil{
        fmt.Println("http read erro")
        return
    }
    src :=string(body)
    re,_ :=regexp.Compile("\\<[\\S\\s]+?\\>")
    src =re.ReplaceAllStringFunc(src,strings.ToLower)
    re,_=regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
    src=re.ReplaceAllString(src,"")
    fmt.Println(strings.TrimSpace(src))
}

func main(){
    //fmt.Printf("Sqrt(20)=%v\n",mymath.Sqrt(20))
    go say("world")
    say("Hello")
    a:=[]int{7,2,8,-9,4,0}
    c:=make(chan int)
    go sum(a[:len(a)/2],c)
    go sum(a[len(a)/2:],c)
    x,y := <-c,<-c
    fmt.Println(x,y,x+y)
    getHttpPage()
}
