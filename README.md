
# Usage

## Create vereify-token Package
```
$ mkdir ./library/verify-token
$ vim ./library/verify-token/verify.go
```

``` ./library/verify-token/verify.go
package verify

type VerifyClient struct{}

func NewVerifyClient() *VerifyClient {
	return &VerifyClient{}
}

func (t *VerifyClient) VerifyToken(token string) bool {
	// verifying token
	return true
}
```

## Run Cloud Datastore emutrator
[Cloud Datastore エミュレータの実行  |  Cloud Datastore ドキュメント  |  Google Cloud](https://cloud.google.com/datastore/docs/tools/datastore-emulator)

```
$ gcloud components install cloud-datastore-emulator
$ $(gcloud beta emulators datastore env-init)
$ gcloud beta emulators datastore start
```

## Run Application
```
$ go run .
```

# Sample Query

```
query{
  user {
    name
    email
    nick_name
    description
  }
  contents(limit: 1) {
    id
    title
    publish_date
    body
    tags {
      tag_desc
      tag_name
    }
  }
}
```

# Create Content
```
curl -X POST \
  https://localhost:8080/v1/contents \
  -H 'Content-Type: application/json' \
  -d '{
	"title" : "ブログのタイトルをどうしようかなと思っている2",
	"tags": [
		"go"
		],
	"body":"#たいとるが決まりません2"
}'
```
