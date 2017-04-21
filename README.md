
# stan-go-test

A Go app, which can easily be deployed to Heroku.

This application is written for [Stan Coding Challenge](https://challengeaccepted.streamco.com.au/).

## Running Locally

Make sure you have [Go](http://golang.org/doc/install) and the [Heroku Toolbelt](https://toolbelt.heroku.com/) installed.

```sh
$ go get -u github.com/armansa/stan-go-test
$ cd $GOPATH/src/github.com/armansa/stan-go-test
$ heroku local
```

You should also install [Godep](https://github.com/tools/godep)

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku master
$ heroku open
```
