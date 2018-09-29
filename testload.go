package main 

import (
"fmt"
"encoding/json"
"io/ioutil"

)
/*
type Config struct {
	Mode string `json:"Mode"`
}
*/

type Config struct {
	Name string 
	Order string
}
var (
	configs []Config
	//dataMarshal []Config
)
func loadConfig() bool {
//=============load json
	buff , err := ioutil.ReadFile("jsn.json")
	if err != nil {
		fmt.Println("load json error : " , err)
		return false
	}
	/*
	var buff = [ ] byte ( ` [ 
        	{ "Name" : "Platypus" , "Order" : "Monotremata" } , 
	        { "Name" : "Quoll" ,     "Order" : "Dasyuromorphia" } 
    	] ` ) 
	*/
	fmt.Println(" the json is : " , string(buff))
//================Unmarshal json
	err = json.Unmarshal(buff , & configs) //  read buff >>> configs
	if err != nil {
		fmt.Println("parse json error : " , err)
		return false
	}
	fmt.Printf("configs struct : %+v \r\n" , configs)
	fmt.Printf("configs[0].Name : %+v \r\n" , configs[0].Name)
//=============== Marshal json
	var p []Config = make([]Config, 0) 
	ooxx := Config{
			Name : "oo" , 
			Order : "xx" ,
		}
	p = append(p,ooxx)
	ooxx.Name = "nooo"
	p = append(p , ooxx)
	str , _ := json.Marshal(p)
	fmt.Printf("json is :%s \r\n" , string(str))
	return true
}
func main(){
	fmt.Println("main")
	loadConfig()
	//fmt.Println("load configs , mode: " , configs.Mode)
}
func init(){
	fmt.Println("init")
}
