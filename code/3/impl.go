package boring

type Player interface {
	Play()
}

type Person struct{}

func (*Person) Play() {}

// var _ Player = Person{} // :(
var _ = Player(&Person{})     // do you know the difference?
var _ Player = (*Person)(nil) //

// Notes:
// Good when you export an interface, but don't use it by yourself.
// aws-sdk-go is a great example for it.
