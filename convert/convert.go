package convert

import "github.com/jinzhu/copier"

func Copy(to, from any) {
	err := copier.Copy(to, from)
	if err != nil {
		panic(err)
	}
}
