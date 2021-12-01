package news

import "rulzmotoshop/business/news"

type NewsResponse struct {
	Article interface{} `json:"news"`
}

func FromDomain(domain news.Domain) NewsResponse {
	return NewsResponse{
		Article: domain.Article,
	}
}

func FromListDomain(domain []news.Domain) []NewsResponse {
	var response []NewsResponse
	for _, value := range domain {
		response = append(response, FromDomain(value))
	}
	return response
}
