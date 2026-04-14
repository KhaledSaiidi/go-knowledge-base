package airportrobot

import "fmt"
type Greeter interface {
    LanguageName() string
    Greet(visitor string) string
}

func SayHello(visitor string, l Greeter) string{
    return fmt.Sprintf("I can speak %s: %s", l.LanguageName(), l.Greet(visitor))
}
type Italian struct{}
type Portuguese struct{}

func(Italian) LanguageName () string{
    return "Italian";
}
func(Portuguese) LanguageName () string{
    return "Portuguese";
}

func (Italian) Greet(visitor string) string{
    return fmt.Sprintf("Ciao %s!", visitor)
}

func (Portuguese) Greet(visitor string) string{
    return fmt.Sprintf("Olá %s!", visitor)
}