package model

// why we need to import primitives?
// because we are using primitive.ObjectID as type for _id field in Course struct which is given by mongo
import(
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// why `json:"_id, omitempty" bson:"_id,omitempty"`
// because we want to use this field as _id in mongo and we want to omit it if it is empty, why bson?
// because we are using bson to marshal and unmarshal data in mongo
type Netflix struct{
	Id primitive.ObjectID `json:"_id, omitempty" bson:"_id,omitempty"`
	Movie string `json:"movie"`
	Watched bool `json:"watched"`
}