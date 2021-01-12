# go-gif-animation

## Install

    go get github.com/fktym/go-gif-animation@v0.0.1

## Usage

```
// create param
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

// create
data, err := Create(p)
if err != nil { // handle error
	panic(err)
}

// save data
err = ioutil.WriteFile("../../test/out/TestCreate.gif", data, 0644)
if err != nil {
	panic(err)
}
```
