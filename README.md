# facebook-account-kit

<img src="./images/screen.png">

Account Kit for Web (Golang), see the [demo site](https://facebook-account-kit-example.herokuapp.com/).

## Setup facebook account kit information

copy the `.env.example` to `.env`

```
TEST_FACEBOOK_APP_ID=xxxxxxxxx
TEST_FACEBOOK_SECRET=xxxxxxxxx
TEST_FACEBOOK_VERSION=v1.1
```

change the `app_id` and `secret value`.

## Run the app in go v1.11 version

Please make sure the go version in v1.11 using [go module](https://github.com/golang/go/wiki/Modules).

```sh
$ go run main.go
```
