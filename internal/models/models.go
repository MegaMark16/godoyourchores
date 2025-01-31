package models

var chores = map[string][]string{
	"user1": {"Take out trash", "Wash dishes", "Mow lawn"},
}

func GetChores(username string) []string {
	return chores[username]
}
