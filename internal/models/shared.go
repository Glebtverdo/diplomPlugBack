package models

type ResponseOnlyWithId struct {
	Id int `json:"id"`
}

type AnyObj interface {
	Object | ObjectBody | Request | ServiceOrganization | User
}
