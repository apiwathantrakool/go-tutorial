# To declare a variable

Using `:=` for example,

```
text := "Text";
text = "Text edit";
```

Assign with Typescript,

```
const text string = "Hello World";
var text02 string = "Hello World 02";
const numberTest int = 100;
const booleanTest bool = true;
var deckSize int
deckSize = 58
```

# Simple function

```
func main(){
	text := getText();
	fmt.Println(text)
}

func getText() string {
	return "Hello World";
}
```

# Slice and For loop

```
cards := []string{"A", "B"}
cards = append(cards, "C")
for index, card := range cards{
	fmt.Println(index,card)
}
```

# Type and Receiver function

```
<!-- deck.go -->
type deck []string

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}
```

So that, any variables type `deck` can access function `print()`, because we set the receiver parameter `(d deck)`.

```
<!-- main.go -->
cards := deck{"A", "B", "C"}
cards.print()
```
