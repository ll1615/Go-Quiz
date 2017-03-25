package listops

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestList(t *testing.T) {
	Convey("test list operation ", t, func() {
		list1 := []interface{}{
			"python", "go", "shell", "java", "nodejs", "c#", "go", "go",
		}
		list2 := []interface{}{
			"go", "lisp", "javascript", "html", "css", "lisp", "lisp",
		}
		list3 := []interface{}{}
		list4 := []interface{}{
			1, 2, 3, 4, 5, 6, 7, 7, 7,
		}
		list5 := []interface{}{
			5, 6, 7, 8, 9, 0,
		}
		list6 := []interface{}{
			5, 6, 7, "lisp", "html",
		}

		So(len(UniqueList(list1)), ShouldEqual, 6)

		So(len(ListMinus(list1, list2)), ShouldEqual, 5)
		So(len(ListMinus(list2, list1)), ShouldEqual, 6)
		So(len(ListAnd(list1, list2)), ShouldEqual, 3)
		So(len(ListAnd(list2, list1)), ShouldEqual, 1)

		So(len(ListMinus(list1, list3)), ShouldEqual, 8)
		So(len(ListMinus(list3, list1)), ShouldEqual, 0)
		So(len(ListAnd(list1, list3)), ShouldEqual, 0)

		So(len(ListMinus(list4, list5)), ShouldEqual, 4)
		So(len(ListMinus(list5, list4)), ShouldEqual, 3)
		So(len(ListAnd(list4, list5)), ShouldEqual, 5)
		So(len(ListAnd(list5, list4)), ShouldEqual, 3)

		So(len(ListMinus(list2, list6)), ShouldEqual, 3)
		So(len(ListMinus(list6, list2)), ShouldEqual, 3)
		So(len(ListAnd(list2, list6)), ShouldEqual, 4)
		So(len(ListAnd(list6, list2)), ShouldEqual, 2)

		So(len(ListMinus(list4, list6)), ShouldEqual, 4)
		So(len(ListMinus(list6, list4)), ShouldEqual, 2)
		So(len(ListAnd(list4, list6)), ShouldEqual, 5)
		So(len(ListAnd(list6, list4)), ShouldEqual, 3)
	})
}
