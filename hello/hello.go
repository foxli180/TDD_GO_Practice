package hello

const helloPrefix  = "Hello "
const spanish  = "Spanish"
const french  = "French"
const spanishHelloPrefix  = "Hola "
const frenchHelloPrefix  = "Bonjour "
const defaultName  = "World"

func Hello(name string, language string) string {
	if name == "" {
		name = defaultName
	}

	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}

