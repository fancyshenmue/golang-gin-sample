package pkg

import (
	"context"

	customLogHandle "golang-gin-sample/pkg/loghandle"
)

// InsertSingleDocument | Insert Single Document
func (c *MongoInsertHelper) InsertSingleDocument() {
	customLogHandle.LogInfo(
		"InsertSingleDocument",
		"Insert Single Document",
		"Insert Single Document",
	)

	ctx := context.TODO()

	collection := c.Cl.Database(c.Db).Collection(c.Coll)

	_, err := collection.InsertOne(ctx, c.Data)
	customLogHandle.ErrorHandle(
		"InsertSingleDocument",
		"Insert Single Document",
		"Insert Single Document error",
		err,
	)
}
