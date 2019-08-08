package fplib

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBool(t *testing.T) {

	Convey("Subject: Test Bool()\n", t, func() {
		Convey("basic", func() {
			So(Bool(true), ShouldEqual, true)
			So(Bool(false), ShouldEqual, false)
		})
		Convey("ext", func() {
			So(Bool("true"), ShouldEqual, true)
			So(Bool("True"), ShouldEqual, true)
			So(Bool("TRUE"), ShouldEqual, true)
			So(Bool("T"), ShouldEqual, true)
			So(Bool("t"), ShouldEqual, true)
			So(Bool("1"), ShouldEqual, true)
			So(Bool(1), ShouldEqual, true)
			So(Bool(120), ShouldEqual, true)
			So(Bool(-10), ShouldEqual, true)
			So(Bool(1.1), ShouldEqual, true)
			So(Bool([]string{"1"}), ShouldEqual, true)
			So(Bool("false"), ShouldEqual, false)
			So(Bool("False"), ShouldEqual, false)
			So(Bool("FALSE"), ShouldEqual, false)
			So(Bool("F"), ShouldEqual, false)
			So(Bool("f"), ShouldEqual, false)
			So(Bool(""), ShouldEqual, false)
			So(Bool("0"), ShouldEqual, false)
			So(Bool(0), ShouldEqual, false)
			So(Bool(0.0), ShouldEqual, false)
			So(Bool([]string{}), ShouldEqual, false)
		})

	})
}
