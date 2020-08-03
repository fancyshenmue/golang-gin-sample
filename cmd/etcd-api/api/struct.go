package api

/* auth */

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	UserName  string
	FirstName string
	LastName  string
}

/* end of auth */

/* etcd */

type Environment struct {
	Env string
}

type GetETCdKeyTop struct {
	Key string `json:"key"`
}

type GetETCdKeyWithPrefixTop struct {
	Prefix string `json:"prefix"`
}

type PutETCdKeyTop struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DeleteETCdKeyTop struct {
	Key string `json:"key"`
}

/* end of etcd */
