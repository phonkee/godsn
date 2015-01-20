package godsn

import (
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDSNParse(t *testing.T) {
	Convey("Test parse DSN.", t, func() {
		p, err := Parse("redis://localhost:6379/?test=123")
		So(err, ShouldBeNil)
		So(p.Host(), ShouldEqual, "localhost:6379")
		So(p.Path(), ShouldEqual, "/")
		So(p.Scheme(), ShouldEqual, "redis")

		_, err = Parse("%://")
		So(err, ShouldNotBeNil)
	})

	Convey("Test parse DSN Values.", t, func() {
		_, errParse := ParseQuery("%")
		So(errParse, ShouldNotBeNil)

		p, err := ParseQuery("test=123&prefix=pref&bool1=true&bool2=1&bool3=0")
		So(err, ShouldBeNil)
		So(p.GetInt("test", 345), ShouldEqual, 123)
		So(p.GetInt("non-existing", 345), ShouldEqual, 345)

		So(p.GetString("prefix", "default"), ShouldEqual, "pref")
		So(p.GetString("non-existing-prefix", "default"), ShouldEqual, "default")

		So(p.GetBool("bool1", false), ShouldEqual, true)
		So(p.GetBool("bool2", false), ShouldEqual, true)
		So(p.GetBool("bool3", true), ShouldEqual, false)
		So(p.GetBool("bool4", true), ShouldEqual, true)

		_, errV := NewValues(url.Values{})
		So(errV, ShouldBeNil)
	})

}
