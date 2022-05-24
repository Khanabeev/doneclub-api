package role

type Role string

const (
	User      Role = "user"
	Member    Role = "member"
	Moderator Role = "moderator"
	Admin     Role = "admin"
)
