# Command

To run package: `go run main.go deck.go`
To run test: `go test`

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
Note, we don't ever reference a receiver value as 'this' or 'self'. For example, `func (this deck) print()`

```
<!-- main.go -->
cards := deck{"A", "B", "C"}
cards.print()
```

# Indexing

If `data = ["A", "B", "C"];`
And, `data[startIndexIncluding: upToNotIncluding]`
Then, `data[0:2]` is `["A","B"]`
Or, `data[:2]` is `["A","B"]`
Or, `data[1:]` is `["B","C"]`

# Testing

Create a new file ending in `_test.go` for example `deck_test.go`

# Interface and Struct

Think about OOP concept.

- Struct is a main structure or blueprint of the object. Attribute
- Interface is additional behavior or accessory of the object. Behavior

# Array vs Slice vs Map

- Array = [N]T
- Slice = []T
- Map = map[K]T

N is number max values in array.
T is datatype.
K is key or id of an object.

# Pointer

- Use * when want to update the value which is located in the memory.
- Use & when want to create an object of a struct.