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

/* consul */

type CommonInterface interface{}
type CommonMap map[string]interface{}
type CommonMapSlice map[string]interface{}

type GetConsulKeyTop struct {
	Key string `json:"key" binding:"required"`
}

type GetConsulKeyWithPrefixTop struct {
	Prefix string `json:"prefix" binding:"required"`
}

type PutConsulKeyTop struct {
	Key   string      `json:"key" binding:"required"`
	Value interface{} `json:"value" binding:"required"`
}

type DeleteConsulKeyTop struct {
	Key string `json:"key" binding:"required"`
}

/* end of etcd */
