package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTestAggregate(t *testing.T) {

	Convey("Given the aggregate test with certain state", t, func() {

		Convey("When I invoke an aggregate's method", func() {

			Convey("Then something will happen", func() {

				So(true, ShouldEqual, true)

			})

		})

	})

}

func TestTestAggregate2(t *testing.T) {

	Convey("Given the aggregate test with certain state", t, func() {

		Convey("When I invoke an aggregate's method", func() {

			Convey("Then something will happen", func() {

				So(true, ShouldEqual, false)

			})

		})

	})

}

func main() {
	var t = testing.T{}
	TestTestAggregate(&t)
	TestTestAggregate2(&t)
}
