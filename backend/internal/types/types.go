package types

type UserRole string

const (
	AdminRole  UserRole = "admin"
	MemberRole UserRole = "member"
)

type ServiceType string

const (
	PsqlServiceType  ServiceType = "psql"
	AppServiceType   ServiceType = "app"
)
