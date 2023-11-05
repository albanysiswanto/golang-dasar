package database

var connection string

// Akan di eksekusi pertama kali
func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
