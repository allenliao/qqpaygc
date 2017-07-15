package models

type QQGroupInfo struct {
	IdGroup   string
	QQ        string
	GroupName string
}

type QQExpireInfo struct {
	IdAccountGroup string
	QQ             string
	GroupName      string
	StartDate      string
	ExpireDate     string
}

func init() {
}
