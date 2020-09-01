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


#### Commands

* heroku apps --> Listado de Apps
* heroku create --> Crea app en Heroku
* heroku stack:set container -a warm-savannah-31868 --> Convierte la app en modo container
* heroku plugins:install @heroku-cli/plugin-manifest
* heroku git:remote -a warm-savannah-31868 --> Attach to heroku git repo
* heroku addons:create heroku-postgresql:hobby-dev --> Create addons 
* git push heroku master
* heroku logs --tail : Logs de Heroku