package main

import (
	"fmt"
	"going/puzzles"
)

func main() {
	//p := Test{Name: "gorgine", Age: 24}
	//fmt.Printf("%v\n", p)
	//fmt.Printf("%+v\n", p)
	//fmt.Printf("%#v\n", p)
	//db, err := g.InitDB()
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//fmt.Println(db)
	//
	//err1 := db.AutoMigrate(&g.User{})
	//if err1 != nil {
	//	log.Fatal(err1.Error())
	//}
	//
	//user := g.User{
	//	Username: "hwwgj",
	//	Age:      24,
	//}
	//db.Create(&user)
	//fmt.Println(user.ID)
	res := puzzles.PrintString('a')
	fmt.Println(res)
}
