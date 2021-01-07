package animation

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	p := CreateAnimationParam{
		Frames: []CreateFrameParam{
			{URI: "../../test/data/G.png", Delay: 100},
			{URI: "../../test/data/O.png", Delay: 100},
			{URI: "../../test/data/P.png", Delay: 100},
			{URI: "../../test/data/H.png", Delay: 100},
			{URI: "../../test/data/E.png", Delay: 100},
			{URI: "../../test/data/R.png", Delay: 100},
		},
	}

	data, err := Create(p)
	assert.Nil(t, err)

	err = ioutil.WriteFile("../../test/out/TestCreate.gif", data, 0644)
	if err != nil {
		panic(err)
	}
}
