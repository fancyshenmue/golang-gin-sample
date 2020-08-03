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

/* api */

type MongoGetBodyTop struct {
	Collection string
	FindField  string
	FindData   string
}

type MongoInsertSingleDocumentBodyTop interface{}

type MongoInsertManyDocumentBodyTop []MongoInsertSingleDocumentBodyTop

type MongoUpdateSingleDocumentBodyTop struct {
	Name  string
	Field string
	Info  interface{}
}

type MongoDeleteSingleDocumentBodyTop struct {
	Name string
}

/* end of api */
