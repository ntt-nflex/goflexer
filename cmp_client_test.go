package goflexer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAPIGet(t *testing.T) {

	Convey("When a GET request is made to cmp /users", t, func() {
		conf := NewConfigFromYAML()
		client := NewCmpClient(conf)

		result, err := client.Get("/users", nil)

		So(err, ShouldBeNil)
		So(result.StatusCode(), ShouldEqual, 200)

		// t.Logf("RESULT: %+v", result)
	})
}
