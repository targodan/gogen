package conv

import "testing"
import . "github.com/smartystreets/goconvey/convey"

func TestTextToByteSlice(t *testing.T) {
	Convey("Negative values should work.", t, func() {
		i := -1
		res, err := TextToByteSlice("-1")
		So(err, ShouldBeNil)
		So(res, ShouldResemble, []byte{byte(i)})
	})
	Convey("Positive, large values should work.", t, func() {
		res, err := TextToByteSlice("254")
		So(err, ShouldBeNil)
		So(res, ShouldResemble, []byte{254})
	})
	Convey("Multiple entries should work.", t, func() {
		res, err := TextToByteSlice("1, 5, 0, 4, 5")
		So(err, ShouldBeNil)
		So(res, ShouldResemble, []byte{1, 5, 0, 4, 5})
	})
	Convey("Any amount of whitespaces should work.", t, func() {
		res, err := TextToByteSlice("  \t 1  ,  \n 5  ,  \r\n 0,\t 4, 5 \t")
		So(err, ShouldBeNil)
		So(res, ShouldResemble, []byte{1, 5, 0, 4, 5})
	})
	Convey("Wrapping shouldn't change anything.", t, func() {
		res1, err1 := TextToByteSlice("1, 5, 0, 4, 5")
		res2, err2 := TextToByteSlice("[]byte{1, 5, 0, 4, 5}")
		So(err1, ShouldBeNil)
		So(err2, ShouldBeNil)
		So(res1, ShouldResemble, res2)
	})
}
