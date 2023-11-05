package main

import (
	"flag"
	"fmt"
)

/*
PACKAGE FLAG
berisikan fungsionalitas untuk memparsing command line argumend
*/

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	user := flag.String("user", "root", "Put your database user")
	password := flag.String("password", "root", "Put your databse password")

	flag.Parse() // Wajib dipanggil

	fmt.Println("Host: ", *host)
	fmt.Println("User: ", *user)
	fmt.Println("Password: ", *password)
}
