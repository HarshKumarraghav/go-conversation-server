package chat

import "time"

// The above type represents a conversation with an ID, room ID, message, and creation timestamp.
// @property {string} ID - The ID property is a string that represents the unique identifier of the
// conversation. It is tagged with `json:"Id"` and `bson:"_id"` to specify the field names when
// encoding and decoding JSON and BSON respectively.
// @property {string} RoomId - The `RoomId` property is a string that represents the ID of the room
// associated with the conversation.
// @property {string} Message - The "Message" property is a string that represents the content of a
// conversation message.
// @property CreatedAt - CreatedAt is a property of type time.Time that represents the timestamp when
// the conversation was created.
type Conv struct {
	ID        string    `json:"Id" bson:"_id"`
	RoomId    string    `json:"roomid" bson:"roomid"`
	Message   string    `json:"message" bson:"message"`
	CreatedAt time.Time `json:"createAt" bson:"CreatedAt"`
}

// The `ToConv()` function is a method defined on the `Conv` struct. It takes a pointer to a `Conv`
// object as its receiver (`in *Conv`), and it returns a new `Conv` object.
func (in *Conv) ToConv() Conv {
	return Conv{
		ID:        in.ID,
		RoomId:    in.RoomId,
		Message:   in.Message,
		CreatedAt: time.Now(),
	}
}

// The `func (u *Conv) ToOutCov() Conv` function is a method defined on the `Conv` struct. It takes a
// pointer to a `Conv` object as its receiver (`u *Conv`), and it returns a new `Conv` object.
func (u *Conv) ToOutCov() Conv {
	return Conv{
		ID:        u.ID,
		RoomId:    u.RoomId,
		Message:   u.Message,
		CreatedAt: time.Now(),
	}
}
