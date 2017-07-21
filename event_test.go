package goflexer

import (
	"encoding/json"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

type TestData struct {
	Counter int64     `json:"counter"`
	Name    string    `json:"name"`
	Floater float64   `json:"floater"`
	Date    time.Time `json:"date"`
	Nested  FooBar    `json:"foobar"`
}

type FooBar struct {
	Foo string `json:"foo"`
	Bar string `json:"bar"`
}

func TestEventGetters(t *testing.T) {

	Convey("When a new Event is created", t, func() {

		ts := time.Now()
		td := TestData{
			Counter: 1,
			Name:    "test",
			Floater: 99.9,
			Date:    ts,
			Nested: FooBar{
				Foo: "funky",
				Bar: "monkey",
			},
		}

		data, err := json.Marshal(td)
		So(err, ShouldBeNil)

		event := NewEvent(json.RawMessage(data))

		Convey("Get method should return values", func() {
			strTest, ok := event.Get("name")
			So(ok, ShouldBeTrue)
			So(strTest, ShouldHaveSameTypeAs, "test")
			So(strTest, ShouldEqual, "test")

			iTest, ok := event.Get("counter")
			So(ok, ShouldBeTrue)
			So(iTest, ShouldEqual, 1)
			//So(iTest, ShouldHaveSameTypeAs, 1)

			fTest, ok := event.Get("floater")
			So(ok, ShouldBeTrue)
			So(fTest, ShouldEqual, 99.9)
			So(iTest, ShouldHaveSameTypeAs, 99.9)
		})

		Convey("Get method should return nil,false for invalid keys", func() {
			strTest, ok := event.Get("xxx")
			So(ok, ShouldBeFalse)
			So(strTest, ShouldBeNil)
		})

		Convey("GetType methods should return typed values", func() {
			strTest, err := event.GetString("name")
			So(err, ShouldBeNil)
			So(strTest, ShouldHaveSameTypeAs, "test")
			So(strTest, ShouldEqual, "test")

			iTest, err := event.GetInt("counter")
			So(err, ShouldBeNil)
			So(iTest, ShouldHaveSameTypeAs, int64(1))
			So(iTest, ShouldEqual, 1)

			fTest, err := event.GetFloat("floater")
			So(err, ShouldBeNil)
			So(fTest, ShouldHaveSameTypeAs, float64(99.9))
			So(fTest, ShouldEqual, 99.9)
		})

		Convey("Unmarshal method give access to struct data", func() {

			structTest := TestData{}
			err := event.Unmarshal(&structTest)
			So(err, ShouldBeNil)
			So(structTest.Counter, ShouldEqual, 1)
			So(structTest.Name, ShouldEqual, "test")
			So(structTest.Floater, ShouldEqual, 99.9)
			So(structTest.Date, ShouldHappenOnOrBetween, ts, ts) // WTF: not equal due to missing location on unmarshaled time object
			So(structTest.Nested.Foo, ShouldEqual, "funky")
		})

		// t.Logf("RESULT: %+v", result)
	})
}
