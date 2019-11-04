package starred

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/machinebox/graphql"
)

type ApiVersion struct {
	Version string
}

type ResponseStruct struct {
	User UserStruct
}

type ErrorStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserStruct struct {
	// name string
	Id                  string
	StarredRepositories RepositoriesStruct
}

type RepositoriesStruct struct {
	TotalCount int
	PageInfo   struct {
		HasNextPage bool
		EndCursor   string
		StartCursor string
	}
	Edges []struct {
		Node struct {
			Id          string
			Name        string
			Description string
			Tags        string
		}
	}
}

func (err *ErrorStruct) Error() string {
	return fmt.Sprintf("[%d] %s", err.Code, err.Message)
}

func FetchStarred(w http.ResponseWriter, r *http.Request) {
	client := graphql.NewClient("https://api.github.com/graphql")
	vars := mux.Vars(r)

	urlParams := r.URL.Query()
	var nextCursor = ""

	if urlParams.Get("next") != "" {
		nextCursor = fmt.Sprintf(", after: \"%v\"", urlParams.Get("next"))
	}

	req := graphql.NewRequest(fmt.Sprintf(`
	query ($login: String!) { 
		user(login: $login) {
			id
			starredRepositories(first: 2%v) {
				totalCount
				pageInfo {
					hasNextPage
					endCursor
				}
				edges {
					node {
						id
						name
						description
						primaryLanguage {
							id
							name
						}
					}
				}
			}
		}
	}
	`, nextCursor))

	req.Var("login", fmt.Sprintf("%v", vars["username"]))

	req.Header.Add("Authorization", fmt.Sprintf("bearer %s", os.Getenv("GITHUB_ACCESS_TOKEN")))
	ctx := context.Background()

	log.Print("aaaa")

	var respData ResponseStruct
	if err := client.Run(ctx, req, &respData); err != nil {

	}

	for _, element := range respData.User.StarredRepositories.Edges {
		log.Print(element)
		// index is the index where we are
		// element is the element from someSlice for where we are
	}

	b, err := json.Marshal(respData)
	log.Print(respData)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
