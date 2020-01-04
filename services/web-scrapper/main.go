package main

import(
	"net/http"
	"log"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/PuerkitoBio/goquery"
)


func main() {
	run()
}

type QuoraUser struct { 
	Name string `json:"name"`
	Biography string `json:"biography"`
	Answers string `json:"answers"`
	Questions string `json:"questions"`
	Shares string `json:"shares"`
	Publications string `json:"publications"` 
	Followers string `json:"followers"`
	Following string `json:"following"`
	MonthlyViews string `json:"monthlyViews"`
	TotalViews string `json:"totalViews"`
	Credentials Credentials `json:"credentials"`
}

type Credentials struct {
	Location string `json:"location"`
	Work string `json:"work"`
	Studies string `json:"studies"`
}

func newQuoraUser(
	name string, 
	bio string, 
	answers string, 
	questions string,
	shares string,
	publications string,
	followers string,
	following string,
	monthlyViews string,
	totalViews string,
	credentials Credentials,
) *QuoraUser {
	return &QuoraUser{
		Name: name, 
		Biography: bio, 
		Answers: answers, 
		Questions: questions,
		Shares: shares,
		Publications: publications,
		Followers: followers,
		Following: following,
		MonthlyViews: monthlyViews,
		TotalViews: totalViews,
		Credentials: credentials, 
	}
}

func run() {

	r := mux.NewRouter()
	// http.HandlerFunc()

	r.HandleFunc("/user", func (w http.ResponseWriter, r *http.Request){
		response, err := http.Get("https://es.quora.com/profile/Sergio-Gutiérrez-5")
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()

		document, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal("Error loading HTTP response body.", err)
		}

		// Header
		top_header := document.Find("div.ProfileNameAndSig")
		name := top_header.Eq(0).Text()
		bio := document.Find("div.ProfileDescription").Text()

		// Left column info
		left_column_info := document.Find(".list_count")
		answers := left_column_info.Eq(0).Text()
		questions := left_column_info.Eq(1).Text()
		shares := left_column_info.Eq(2).Text()
		publications := left_column_info.Eq(3).Text()
		followers := left_column_info.Eq(4).Text()
		following := left_column_info.Eq(5).Text()

		// Credentials
		credentials := document.Find("span.UserCredential")
		work := credentials.Eq(2).Text()
		studies := credentials.Eq(1).Text()
		location := credentials.Eq(3).Text()

		// Content Insights
		contentViews := document.Find(".ContentViewsAboutListItem span")
		monthlyViews := contentViews.Eq(3).Text();
		totalViews := contentViews.Eq(4).Text();

		// log.Print(credentials)


		quoraUser := newQuoraUser(
			name, 
			bio,
			answers,
			questions,
			shares,
			publications,
			followers,
			following,
			monthlyViews,
			totalViews,
			Credentials {
				Location: location,
				Work: work,
				Studies: studies,
			},
		)

		jsonResponse, err := json.Marshal(quoraUser)

		// log.Println(jsonResponse)

		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200);
		w.Write(jsonResponse)

		// Read and get in string data
		// buffer, err := ioutil.ReadAll(response.Body)
		// pageContent := string(buffer)

	})

	http.ListenAndServe(":8080", r)

}



