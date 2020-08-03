package setup

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	/* mongo info */

	MongoAddress  = os.Getenv("MONGO_ADDRESS")
	MongoUser     = os.Getenv("MONGO_USER")
	MongoPass     = os.Getenv("MONGO_PASS")
	MongoDatabase = os.Getenv("MONGO_DATABASE")
	MongoURI      = fmt.Sprintf("mongodb://%s", MongoAddress)
	Credential    = options.Credential{
		Username: MongoUser,
		Password: MongoPass,
	}

	/* end of mongo info */
)
