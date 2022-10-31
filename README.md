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

Simple function

```
func main(){
	text := getText();
	fmt.Println(text)
}

func getText() string {
	return "Hello World";
}
```

Slice and For loop

```
cards := []string{"A", "B"}
cards = append(cards, "C")
for index, card := range cards{
	fmt.Println(index,card)
}
```
