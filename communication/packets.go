package communication

import "encoding/json"

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
)

type packet struct {
	t string
	d []byte
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
