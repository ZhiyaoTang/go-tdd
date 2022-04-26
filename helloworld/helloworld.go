package helloworld

const (
	chinese = "Chinese"
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	chineseHelloPrefix = "你好，"
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case chinese:
		prefix = chineseHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
