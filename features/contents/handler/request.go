package handler

import (
	"sosmedapps/features/contents"
)

type AddContentRequest struct {
	Content string `json:"content" form:"content"`
}

func RequstToCore(dataContent interface{}) *contents.CoreContent {
	res := contents.CoreContent{}
	switch dataContent.(type) {
	case AddContentRequest:
		cnv := dataContent.(AddContentRequest)
		res.Content = cnv.Content
	default:
		return nil
	}
	return &res

}
