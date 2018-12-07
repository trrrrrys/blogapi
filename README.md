

# Sample Query

```
query{
  user {
    name
    email
    nick_name
    description
  }
  contents(num: 1) {
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