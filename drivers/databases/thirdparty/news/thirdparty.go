package news

import (
	"encoding/json"
	"net/http"
	"rulzmotoshop/business/news"
)

type NewsApi struct {
	httpClient http.Client
}

func NewNewsApi() news.Repository {
	return &NewsApi{
		httpClient: http.Client{},
	}
}

func (new *NewsApi) GetNewsByCategory(category string) (news.Domain, error) {
	url := "https://api-berita-indonesia.vercel.app/okezone/" + category
	// fmt.Println(url)
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := new.httpClient.Do(req)

	if err != nil {
		return news.Domain{}, err
	}

	data := Response{}

	err = json.NewDecoder(resp.Body).Decode(&data)

	// fmt.Println(data)
	return data.ToDomain(), err
}
