package selfdestruction

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	klingonNumbers = [...]string{
		`cha'wejjewej`,
		`cha'wejjecha'`,
		`cha'wejjewa'`,
		`wejjewej`,
		`wejjecha'`,
		`wejjewa'`,
		`wej`,
		`cha'`,
		`wa'`,
	}
)

func TestHumanInitSelfDestruction(t *testing.T) {

	Convey("GIVEN a human captain of a human starship", t, func() {
		kirk := Human{
			Captian: true,
		}
		var uhura, scotty Human
		enterprise := Starship{
			Humans:   []Human{kirk, uhura, scotty},
			Klingons: []Klingon{Klingon{}, Klingon{}},
		}

		Convey("WHEN  the self-destruction was initialized successfully", func() {
			out, err := kirk.InitSelfDestruction(10, "Destruct sequence 0, code 1-1 A", enterprise)
			So(err, ShouldBeNil)
			Convey("THEN  at beginning an extra warning in Klingon should be given", func() {
				So(out[0], ShouldEqual, klingonWarning)
				Convey("AND the last 9 seconds should also be pronounced in the Klingon language", func() {
					if len(out) > 9 {
						out = out[len(out)-9:]
					}
					for i, kn := range out {
						So(kn, ShouldEqual, klingonNumbers[i])
					}
				})
			})

		})

	})

}
