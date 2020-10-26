# yell

<img src=man-yelling-560393.jpg height=300px, width=450px/>
The Bitchy Waiter

## Pretty useless but this is my simple way of shouting, "We Can't Breath!"

## How to use this package:

```go
import "github.com/mohammadrabetian/yell"

func main() {
	message, err := yell.FreeSoftware("YOUR_NATIONALITY")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(message)
}
```

##### You should pass your nationality as an argument to `FreeSoftware()` func like `yell.FreeSoftware("Iranian")`
##### You could also pass an empty string, for what it will return I hope is a surprise for you.
