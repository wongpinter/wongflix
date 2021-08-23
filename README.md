## Wongflix

Movie database using omdbapi.com API Clean Architecture

```
$ git clone https://github.com/wongpinter/wongflix.git

$ vi config/config_dev.go // use your configuration, default already using docker 

$ cd wongflix

$ docker compose up -d

$ go run ./cmd
```

## API
User authentication

    /sign-up
    /sign-in

Movie

    v1/movie/search?query=wars&page=1
    v1/movie/id/tt0103359
    v1/movie/title/batman

## TODO
Unit Testing
Mocking
Membership
payment
etc..

## Hope
This simple project become real application..