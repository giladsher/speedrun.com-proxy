package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	API_BASE_URL = "https://www.speedrun.com/api/v1"
	AUTH_HEADER  = "X-API-Key"
)

type PbsResponse struct {
	Data []struct {
		Place int `json:"place"`
		Run   struct {
			ID       string      `json:"id"`
			Weblink  string      `json:"weblink"`
			Game     string      `json:"game"`
			Level    interface{} `json:"level"`
			Category string      `json:"category"`
			Videos   struct {
				Links []struct {
					URI string `json:"uri"`
				} `json:"links"`
			} `json:"videos"`
			Comment string `json:"comment"`
			Status  struct {
				Status     string    `json:"status"`
				Examiner   string    `json:"examiner"`
				VerifyDate time.Time `json:"verify-date"`
			} `json:"status"`
			Players []struct {
				Rel string `json:"rel"`
				ID  string `json:"id"`
				URI string `json:"uri"`
			} `json:"players"`
			Date      string    `json:"date"`
			Submitted time.Time `json:"submitted"`
			Times     struct {
				Primary          string      `json:"primary"`
				PrimaryT         float32     `json:"primary_t"`
				Realtime         string      `json:"realtime"`
				RealtimeT        float32     `json:"realtime_t"`
				RealtimeNoloads  string      `json:"realtime_noloads"`
				RealtimeNoloadsT float32     `json:"realtime_noloads_t"`
				Ingame           interface{} `json:"ingame"`
				IngameT          float32     `json:"ingame_t"`
			} `json:"times"`
			System struct {
				Platform string `json:"platform"`
				Emulated bool   `json:"emulated"`
				Region   string `json:"region"`
			} `json:"system"`
			Splits interface{} `json:"splits"`
			Values struct {
			} `json:"values"`
			Links []struct {
				Rel string `json:"rel"`
				URI string `json:"uri"`
			} `json:"links"`
		} `json:"run"`
	} `json:"data"`
}

type GameResponse struct {
	Data struct {
		ID    string `json:"id"`
		Names struct {
			International string      `json:"international"`
			Japanese      interface{} `json:"japanese"`
			Twitch        string      `json:"twitch"`
		} `json:"names"`
		BoostReceived       int    `json:"boostReceived"`
		BoostDistinctDonors int    `json:"boostDistinctDonors"`
		Abbreviation        string `json:"abbreviation"`
		Weblink             string `json:"weblink"`
		Discord             string `json:"discord"`
		Released            int    `json:"released"`
		ReleaseDate         string `json:"release-date"`
		Ruleset             struct {
			ShowMilliseconds    bool     `json:"show-milliseconds"`
			RequireVerification bool     `json:"require-verification"`
			RequireVideo        bool     `json:"require-video"`
			RunTimes            []string `json:"run-times"`
			DefaultTime         string   `json:"default-time"`
			EmulatorsAllowed    bool     `json:"emulators-allowed"`
		} `json:"ruleset"`
		Romhack    bool          `json:"romhack"`
		Gametypes  []interface{} `json:"gametypes"`
		Platforms  []string      `json:"platforms"`
		Regions    []string      `json:"regions"`
		Genres     []string      `json:"genres"`
		Engines    []string      `json:"engines"`
		Developers []string      `json:"developers"`
		Publishers []string      `json:"publishers"`
		Moderators struct {
			E8En5P80     string `json:"e8en5p80"`
			Zxz7O9Jq     string `json:"zxz7o9jq"`
			O8663W8Z     string `json:"o8663w8z"`
			ZeroJmrl5Z8  string `json:"0jmrl5z8"`
			SevenJ4Ozn58 string `json:"7j4ozn58"`
			Four8Gl4K2X  string `json:"48gl4k2x"`
			E8E0M56J     string `json:"e8e0m56j"`
			Xz9Vng48     string `json:"xz9vng48"`
			J2Yzy968     string `json:"j2yzy968"`
			J4R5Nol8     string `json:"j4r5nol8"`
			EightQ315Lwj string `json:"8q315lwj"`
		} `json:"moderators"`
		Created interface{} `json:"created"`
		Assets  struct {
			Logo struct {
				URI string `json:"uri"`
			} `json:"logo"`
			CoverTiny struct {
				URI string `json:"uri"`
			} `json:"cover-tiny"`
			CoverSmall struct {
				URI string `json:"uri"`
			} `json:"cover-small"`
			CoverMedium struct {
				URI string `json:"uri"`
			} `json:"cover-medium"`
			CoverLarge struct {
				URI string `json:"uri"`
			} `json:"cover-large"`
			Icon struct {
				URI string `json:"uri"`
			} `json:"icon"`
			Trophy1St struct {
				URI string `json:"uri"`
			} `json:"trophy-1st"`
			Trophy2Nd struct {
				URI string `json:"uri"`
			} `json:"trophy-2nd"`
			Trophy3Rd struct {
				URI string `json:"uri"`
			} `json:"trophy-3rd"`
			Trophy4Th struct {
				URI interface{} `json:"uri"`
			} `json:"trophy-4th"`
			Background struct {
				URI interface{} `json:"uri"`
			} `json:"background"`
			Foreground struct {
				URI string `json:"uri"`
			} `json:"foreground"`
		} `json:"assets"`
		Links []struct {
			Rel string `json:"rel"`
			URI string `json:"uri"`
		} `json:"links"`
	} `json:"data"`
}

type CategoryResponse struct {
	Data struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Weblink string `json:"weblink"`
		Type    string `json:"type"`
		Rules   string `json:"rules"`
		Players struct {
			Type  string `json:"type"`
			Value int    `json:"value"`
		} `json:"players"`
		Miscellaneous bool `json:"miscellaneous"`
		Links         []struct {
			Rel string `json:"rel"`
			URI string `json:"uri"`
		} `json:"links"`
	} `json:"data"`
}

func getPersonalBests(user string) (*PbsResponse, error) {
	log.Default().Println("Getting pbs for user", user)
	res, err := http.Get(fmt.Sprintf("%v/users/%v/personal-bests", API_BASE_URL, user))
	if err != nil {
		log.Default().Println("Failed getting results for user", user, err)
		return nil, err
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Default().Println("Error reading the response body", user, err)
		return nil, err
	}
	var responseObject PbsResponse
	json.Unmarshal(responseData, &responseObject)
	return &responseObject, nil
}

func getGameName(gameUrl string) (string, error) {
	res, err := http.Get(gameUrl)
	if err != nil {
		log.Default().Println("Failed to get game name", err)
		return "", err
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Default().Println("Error reading the response body", err)
		return "", err
	}
	var responseObject GameResponse
	json.Unmarshal(responseData, &responseObject)
	return responseObject.Data.Names.International, nil
}

func getCategoryName(categoryUrl string) (string, error) {
	res, err := http.Get(categoryUrl)
	if err != nil {
		log.Default().Println("Failed to get game name", err)
		return "", err
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Default().Println("Error reading the response body", err)
		return "", err
	}
	var responseObject CategoryResponse
	json.Unmarshal(responseData, &responseObject)
	return responseObject.Data.Name, nil
}

func PersonalBests(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("user") || r.URL.Query().Get("user") == "" {
		log.Default().Println("No user in query params", r.URL.RequestURI())
		fmt.Fprintln(w, "Please provide a non-empty user")
		return
	}
	user := r.URL.Query().Get("user")
	responseObject, err := getPersonalBests(user)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	var responseString string
	for i, pb := range responseObject.Data {
		var gameName string
		var categoryName string
		var err error
		runDuration := strings.ToLower(strings.Replace(pb.Run.Times.Primary, "PT", "", 1))
		for _, link := range pb.Run.Links {
			if link.Rel == "game" {
				gameName, err = getGameName(link.URI)
				if err != nil {
					fmt.Fprintln(w, err)
					return
				}
			} else if link.Rel == "category" {
				categoryName, err = getCategoryName(link.URI)
				if err != nil {
					fmt.Fprintln(w, err)
					return
				}
			}
		}
		if i == len(responseObject.Data)-1 {
			responseString += fmt.Sprintf("%v - %v: %v.", gameName, categoryName, runDuration)
		} else {
			responseString += fmt.Sprintf("%v - %v: %v,", gameName, categoryName, runDuration)
		}
	}
	fmt.Fprintln(w, responseString)
}
