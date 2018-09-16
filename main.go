package main

var (
	luser     = "postgres"
	lpassword = "1234"
	ldbname   = "nuveo"
)

func main() {
	a := App{}
	a.InitDatabase(
		luser,
		lpassword,
		ldbname)

	a.Run(":8080")
}
