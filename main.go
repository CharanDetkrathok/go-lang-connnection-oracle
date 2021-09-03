package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)
 
func main(){

    // var (
    //     username_email string
    // )

 
	// db, err := sql.Open("godror", `user="scenter01" password="scenter01new" connectString="10.2.1.98:1571/RUBRAM" libDir="C:/instantclient_19_12"`)
	
    
	// if err != nil {
    //     fmt.Println(err)
    //     return
    // }
    // defer db.Close()

	// rows,err := db.Query("SELECT USERNAME FROM SCENTER01.SV_USERS")
    // if err != nil {
    //     fmt.Println("Error running query")
    //     fmt.Println(err)
    //     return
    // }
    // defer rows.Close()
    // for rows.Next() {
    //     err := rows.Scan(&username_email)
    //     if err != nil {
    //         log.Fatal(err)
    //     }
    //     log.Println(username_email)
    // }
    // err = rows.Err()
    // if err != nil {
    //     log.Fatal(err)
    // }

    const dns = `user="scenter01" 
                 password="scenter01new" 
                 connectString="10.2.1.98:1571/RUBRAM?expire_time=2&connect_time=2" 
                 timezone="Asia/Bangkok"
                 libDir="C:/instantclient_19_12"`


    db, err := sql.Open("godror", dns)


    if err != nil {
        fmt.Println(err)
    }
    defer db.Close()

    err =db.Ping()
    if err != nil {
        fmt.Println("Oooops! Ping failed...")
        panic(err)
    }

    fmt.Println("Pinged DB OK... จรัญ")


}