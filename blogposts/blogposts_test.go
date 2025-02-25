package blogposts_test

import (
	"github/bek-shoyatbek/blogposts"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
  const (
    firstBody = `Title: Post 1
Description: Description 1`
    secondBody = `Title: Post 2
Description: Description 2`
  )

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))

	}

  assertPost(t, post[0], blogposts.Post{Title: "Post 1",Description: "Description: 1"})

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}


func assertPost(t *testing.T, got,want blogposts.Post,want ){
  t.Helper()
  if !reflect.DeepEqual(got, want) {
    t.Errorf("got %d, want %+v",got, want)
  }
}
