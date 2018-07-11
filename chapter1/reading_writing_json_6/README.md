```go
http.Handle("/images/", newFooHandler())
http.Handle("/images/", newBarHandler())
http.Handle("/images", newBuzzHandler())
/images                  => Buzz
/images/                 => Foo
/images/cat              => Foo
/images/cat.jpg          => Foo
/images/persian/cat.jpg  => Bar
```
Longer paths will always take precedence over shorter ones so it is possible to have a explicit route that points to a different handler to a catch all route