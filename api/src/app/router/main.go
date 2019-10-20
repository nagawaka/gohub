package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"context"

	"github.com/gorilla/mux"
	"github.com/machinebox/graphql"
)

type ApiVersion struct {
	Version string
}

type ResponseStruct struct {
	User UserStruct
}

type UserStruct struct {
	// name string
	Id string
	StarredRepositories RepositoriesStruct
}

type RepositoriesStruct struct {
	TotalCount int
	PageInfo struct {
		HasNextPage bool
		EndCursor string
		StartCursor string
	}
	Edges []struct {
		Node struct {
			Id string
			Name string
			Description string
			Tags string
		}
	}
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	m := ApiVersion{"1.0"}
	b, err := json.Marshal(m)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func fetchStarred(w http.ResponseWriter, r *http.Request) {
	client := graphql.NewClient("https://api.github.com/graphql")
	vars := mux.Vars(r)

	urlParams := r.URL.Query();
	var nextCursor = ""

	if urlParams.Get("next") != "" {
		nextCursor = fmt.Sprintf(", after: \"%v\"", urlParams.Get("next"));
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

	// set any variables
	// req.Var("Authorization", fmt.Sprintf("bearer %s", os.Getenv("GITHUB_ACCESS_TOKEN")))
	req.Header.Add("Authorization", fmt.Sprintf("bearer %s", os.Getenv("GITHUB_ACCESS_TOKEN")))
	ctx := context.Background()

	// log.Print(req.Header);

	// run it and capture the response
	var respData ResponseStruct
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}

	// w.Write([]byte(fmt.Sprintf("%v", respData)));
	
	// m := ApiVersion{"1.0"}
	respData.User.StarredRepositories.Edges[0].Node.Tags = "Tags"
	b, err := json.Marshal(respData)
	log.Print(respData);

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func Init() {
	fmt.Printf("api Started\n")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/starred/{username}", fetchStarred).Methods("GET")
	// router.HandleFunc("/tags", fetchTags).Methods("GET")
	// router.HandleFunc("/tags/{tag_id}", fetchTags).Methods("POST")
	// router.HandleFunc("/tags/{tag_id}", fetchTags).Methods("PUT")
	// router.HandleFunc("/tags/{tag_id}", fetchTags).Methods("DELETE")
	// router.HandleFunc("/repositories/tag/{tag_id}", fetchReposByTag).Methods("GET")
	
	
	log.Fatal(http.ListenAndServe(":8080", router))
}