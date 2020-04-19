package test

import (
	"testing"

	"../router"
)

func TestRouterPath(t *testing.T) {

	t.Log("Test")
	t.Run("/ Set", func(t *testing.T) {

		result, status := router.RouterPath("/")

		if result != "index.html" || status != "200" {

			t.Fatal("/ Error")
		}

	})

	t.Run("index.html", func(t *testing.T) {
		result2, status2 := router.RouterPath("/index.html")
		if result2 != "index.html" || status2 != "200" {

			t.Fatal("index.html Error")
		}

	})

	t.Run("notfound", func(t *testing.T) {
		result3, status3 := router.RouterPath("/test5.html")
		if result3 != "notfound.html" || status3 != "404" {

			t.Fatal("Not Found Error")

		}

	})

	t.Log("Test Down")

}
