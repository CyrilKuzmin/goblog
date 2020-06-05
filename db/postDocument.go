package db

type PostDocument struct {
	ID          string `bson:"_id,omitempty"`
	Title       string
	ContentHTML string
	CreateTime  string
	ModifyTime  string
}
