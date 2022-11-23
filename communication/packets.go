package communication

import (
	"encoding/json"
	"errors"
)

const (
	Comments             string = "Comments"
	Reports                     = "CommentReports"
	ModAssignmentRecords        = "ModAssignmentRecords"
	FeedbackRecords             = "FeedbackRecords"
	LoginResult                 = "LoginResponse"
	ModerationRecords           = "ModerationRecords"
	AccessLogs                  = "AccessLogs"
	DomainReport                = "DomainReport"
	UsersReport                 = "UsersReport"
	BanRecords                  = "BanRecords"
	ServerMessage               = "Message"
)

type packet struct {
	t string
	d []byte
}

func CreatePacket(some_data any, name string) ([]byte, error) {
	data, err := json.Marshal(some_data)
	if err != nil {
		p := packet{t: name, d: data}
		data, err := json.Marshal(p)
		if err != nil {
			return data, nil
		} else {
			return nil, errors.New("Bad Marshal!")
		}
	} else {
		return nil, errors.New("Bad Marshal!")
	}
}

func CreateCommentsPacket(coms []Comment) []byte {
	data, err := json.Marshal(coms)
	if err != nil {
		p := packet{t: Comments, d: data}
		res, _ := json.Marshal(p)
		return res
	} else {
		return nil
	}
}

func CreateCommentReportsPacket(data []CommentReport) {

}
