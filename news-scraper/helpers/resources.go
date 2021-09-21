package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Agancy struct {
	Agency       string   `json:"agency,omitempty" bson:"agency,omiempty"`
	AgencyURL    string   `json:"agency_url,omitempty" bson:"agency_url,omiempty"`
	RSS          []RssOBJ `json:"rss,omitempty" bson:"rss,omitempty"`
	FieldQueries struct {
		Title            string `json:"title,omitempty" bson:"title,omitempty"`
		Summary          string `json:"summary,omitempty" bson:"summary,omitempty"`
		Content          string `json:"content,omitempty" bson:"content,omitempty"`
		Author           string `json:"author,omitempty" bson:"author,omitempty"`
		AuthorReplaceKey string `json:"author_replace_key,omitempty" bson:"author_replace_key,omitempty"`
		City             string `json:"city,omitempty" bson:"city,omitempty"`
		RelatedTopics    string `json:"related_topics,omitempty" bson:"related_topics,omitempty"`
		Date             string `json:"date,omitempty" bson:"date,omitempty"`
		Category         string `json:"category,omitempty" bson:"category,omitempty"`
	} `json:"field_queries,omitempty" bson:"field_queries,omitempty"`
	RemovingFields []string `json:"removing_fields,omitempty" bson:"removing_fields,omitempty"`
}

type RssOBJ struct {
	Category string `json:"category"`
	Source   string `json:"source"`
}

func GetRss(lang string) ([]Agancy, error) {
	jsonFile, err := os.Open("rsslist\rss.json")

	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully Opened rss.json")

	defer jsonFile.Close()

	byteVal, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	agencies := []Agancy{}
	err = json.Unmarshal(byteVal, &agencies)
	if err != nil {
		return nil, err
	}

	return agencies, nil
}
