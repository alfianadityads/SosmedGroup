package handler

import "sosmedapps/features/comment"

type NewComment struct {
	Comment   string `json:"comment" form:"comment"`
	ContentID uint
}

func RequstToCore(dataComment interface{}) *comment.Core {
	res := comment.Core{}
	switch dataComment.(type) {
	case NewComment:
		cnv := dataComment.(NewComment)
		res.Comment = cnv.Comment
		res.ContentID = cnv.ContentID
	default:
		return nil
	}
	return &res

}
