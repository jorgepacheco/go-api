# go-first-steps

- Examples in /Users/jorgepacheco/go/src/initialProject/

- /Users/jorgepacheco/Documents/pet-project


# Deploy in Heroku with a container 


- In terminal:
    * Build docker image: docker build -t go-api .
    * heroku login
    * heroku create go-api
    * heroku git:remote -a go-first-steps (is the same of heroku app)
    * heroku git:remote -a go-first-steps


    

- Doc

https://github.com/orlmonteverde/go-postgres-microblog/blob/v1.3.0/internal/server/v1/api.go

https://golang.org/src/database/sql/example_test.go

https://blog.friendsofgo.tech/posts/mysql-en-go/

https://dev.to/orlmonteverde/api-rest-con-go-golang-y-postgresql-parte-4-4i31