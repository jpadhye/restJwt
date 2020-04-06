module github.com/jpadhye/restJwt

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gorilla/mux v1.7.4 // indirect
	github.com/jpadhye/restJwt/controllers v0.0.0 // indirect
	github.com/jpadhye/restJwt/driver v0.0.0 // indirect
	github.com/jpadhye/restJwt/models v0.0.0 // indirect
	github.com/jpadhye/restJwt/repository/user v0.0.0 // indirect
	github.com/jpadhye/restJwt/utils v0.0.0 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20200403201458-baeed622b8d8 // indirect
)

replace github.com/jpadhye/restJwt/driver v0.0.0 => ./driver
replace github.com/jpadhye/restJwt/controllers v0.0.0 => ./controllers
replace github.com/jpadhye/restJwt/models v0.0.0 => ./models
replace github.com/jpadhye/restJwt/repository/user v0.0.0 => ./repository/user
replace github.com/jpadhye/restJwt/utils v0.0.0 => ./utils
