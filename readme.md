# MYSQL

```
type User struct {
	Id                 int
	First, Last, Email string
}
	user := []User{}

	db.Select(&user, "select * from users")

	log.Println("users...")
	log.Println(user)

```
