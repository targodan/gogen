package conv

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

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
	Convey("Different formats should work with small positive numbers.", t, func() {
		{
			res, err := TextToByteSlice("2")
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []byte{2})
		}
		{
			res, err := TextToByteSlice("0x2a")
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []byte{0x2a})
		}
		{
			res, err := TextToByteSlice("0b1001")
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []byte{9})
		}
		{
			res, err := TextToByteSlice("010")
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []byte{8})
		}
	})
	Convey("Different formats should work with large positive numbers.", t, func() {
		{
			res, err := TextToByteSlice("255")
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []byte{255})
		}
		{
			res, err := TextToByteSlice("0xfa")
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []byte{0xfa})
		}
		{
			res, err := TextToByteSlice("0b10000001")
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []byte{129})
		}
		{
			res, err := TextToByteSlice("0300")
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []byte{192})
		}
	})

	Convey("Different formats should work with negative numbers.", t, func() {
		{
			res, err := TextToByteSlice("-2")
			i := -2
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []byte{byte(i)})
		}
		{
			res, err := TextToByteSlice("-0x2a")
			i := -0x2a
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []byte{byte(i)})
		}
		{
			res, err := TextToByteSlice("-0b1001")
			i := -9
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []byte{byte(i)})
		}
		{
			res, err := TextToByteSlice("-010")
			i := -8
			So(err, ShouldBeNil)
			So(res, ShouldResemble, []byte{byte(i)})
		}
		{
			_, err := TextToByteSlice("fdsa")
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "Invalid input string.")
		}
	})
	Convey("Wrapping shouldn't change anything.", t, func() {
		res1, err1 := TextToByteSlice("1, 5, 0, 4, 5")
		res2, err2 := TextToByteSlice("[]byte{1, 5, 0, 4, 5}")
		So(err1, ShouldBeNil)
		So(err2, ShouldBeNil)
		So(res1, ShouldResemble, res2)
	})
}

func TestBytesToString(t *testing.T) {
	Convey("Empty slices should work.", t, func() {
		So(BytesToString([]byte{}, 0), ShouldEqual, "[]byte{}")
		So(BytesToString([]byte{}, 1), ShouldEqual, "[]byte{\n}")
	})
	Convey("Non-Empty ones should work too.", t, func() {
		So(BytesToString([]byte{5, 1, 10, 3, 9}, 0), ShouldEqual, "[]byte{0x05, 0x01, 0x0a, 0x03, 0x09}")
	})
	Convey("Custom linebreaks should work.", t, func() {
		So(BytesToString([]byte{5, 1, 10, 3, 9}, 0), ShouldEqual, "[]byte{0x05, 0x01, 0x0a, 0x03, 0x09}")
		So(BytesToString([]byte{5, 1, 10, 3, 9}, 1), ShouldEqual, "[]byte{\n0x05, \n0x01, \n0x0a, \n0x03, \n0x09, \n}")
		So(BytesToString([]byte{5, 1, 10, 3, 9}, 2), ShouldEqual, "[]byte{\n0x05, 0x01, \n0x0a, 0x03, \n0x09, \n}")
		So(BytesToString([]byte{5, 1, 10, 3, 9}, 3), ShouldEqual, "[]byte{\n0x05, 0x01, 0x0a, \n0x03, 0x09, \n}")
		So(BytesToString([]byte{5, 1, 10, 3, 9}, 4), ShouldEqual, "[]byte{\n0x05, 0x01, 0x0a, 0x03, \n0x09, \n}")
	})
}

func TestTextOrStdin(t *testing.T) {
	Convey("Reading text should work.", t, func() {
		text, err := TextOrStdin("asdf")
		So(err, ShouldBeNil)
		So(text, ShouldEqual, "asdf")
	})
	Convey("Reading from stdin should work.", t, func() {
		f, err := os.Open("conv_test.data")
		os.Stdin = f

		text, err := TextOrStdin("-")
		So(err, ShouldBeNil)
		So(text, ShouldEqual, "asdf asdjf3nv")
	})
}

func TestFileOrStdin(t *testing.T) {
	Convey("Reading from file should work.", t, func() {
		text, err := FileOrStdin("conv_test.data")
		So(err, ShouldBeNil)
		So(text, ShouldResemble, []byte("asdf asdjf3nv"))
	})
	Convey("Reading from stdin should work.", t, func() {
		f, err := os.Open("conv_test.data")
		tmp := os.Stdin
		os.Stdin = f
		text, err := FileOrStdin("-")
		os.Stdin = tmp

		So(err, ShouldBeNil)
		So(text, ShouldResemble, []byte("asdf asdjf3nv"))
	})
}

func TestImpossibleErrorCases(t *testing.T) {
	Convey("This should error.", t, func() {
		_, err := litToByteUnsigned("asdf")
		So(err, ShouldNotBeNil)

		_, err = litToByteSigned("asdf")
		So(err, ShouldNotBeNil)
	})
}
