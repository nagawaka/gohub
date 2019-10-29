FORMAT: 1A

# Gohub

Returns starred repositories from github.

# Group Repositories

Resources related to starred repositories.

# Starred [/starred/{username}]

Returns starred repositories by username.

+ token - a token to be returned on headers
+ repositories - the list of repositories

+ Parameters
    + username: username (required, string) - Username from github

# Retrieve Starred Repositories [GET]

+ Response 200 (application/json)

        {
            token: "aaappp",
            repositories: [
                {
                    "id": "1",
                    "name": "Gohub",
                    "description": "Something",
                    "language": {name: "javascript", id: "asssss"},
                    "tags": [
                        {name: "javascript", id: "asssss"},
                    ]
                }
            ]
        }


# Repositories [/repositories/{searchstring}]

Returns list of repositories

+ id
+ name
+ description
+ language - The language object for this repository.
+ tags - An array of tags for this user.

+ Parameters
    + searchstring: java (optional, string) - Searchstring

# Retrieve repositories by SearchString [GET]

+ Request (application/json)

    + Headers
    
        Authorization: Bearer aaappp
    
    + Parameters
    
        id: 1

+ Response 200 (application/json)

    + Body

            [
                    {
                            "id": "1",
                            "name": "Gohub",
                            "description": "Something",
                            "language": {name: "javascript", id: "asssss"},
                            "tags": [
                                {name: "javascript", id: "asssss"},
                                {name: "java", id: "asssss"},
                            ]
                    },
            ]

# Repository [/repository/{id}{?tags}]

+ id
+ name
+ description
+ language - The language object for this repository.
+ tags - An array of tags for this user.

+ Parameters
    + id: 1 (string, required) - Repository id

# Add Or Edit Tags on Repositories  [POST]

+ Request (application/json)

    + Headers
    
        Authorization: Bearer aaappp
        
    + Body
    
        {
            tags: [
                "java",
                "javascript"
            ]
        }
    

+ Response 200 (application/json)

    + Body

            [
                    {
                            "id": "1",
                            "name": "Gohub",
                            "description": "Something",
                            "language": {name: "javascript", id: "asssss"},
                            "tags": [
                                {name: "javascript", id: "1"},
                                {name: "java", id: "2"},
                            ]
                    },
            ]
