# simple-bubble

golang bubbleTea components for fast usage in your terminal applications

## What is this?

This is a go module that provides a set of reusable components for use with the [bubbletea]() library. The components are designed to be simple and easy to use, and to provide a good starting point for building more complex applications.

## Usage

To use this module, You can install it by running:

```bash
go get "github.com/nasoooor29/simple-bubble"
```

You can then import the module in your code like this:

```go
import "github.com/nasoooor29/simple-bubble"
```

## Components

The following components are currently available:

-   [check list](#checkList)
-   [list selection](#listSelection)
-   [text input](#textInput)

### checkList
The checkList component is a simple list of items that can be checked or unchecked. It is useful for presenting a list of options to the user and allowing them to select one or more items.
quick example:

```go
package main

import (
    "fmt"

    sb "github.com/nasoooor29/simple-bubble"
)
func main() {
	checkedItems, err := sb.NewCheckList("Title", sb.NewCheckListItems(
		sb.NewCheckListItem("item 1", false), // item not checked by default
		sb.NewCheckListItem("item 2", true), // item checked by default
        // add as many items as you want
	))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("checked items: %v\n", checkedItems)
}
```

### FancyItemSelector
The FancyItemSelector component is a fancy list that you can only select one item from it only. It is useful for presenting a list of options to the user and allowing them to select one item.

quick example:

```go
package main

import (
    "fmt"

    sb "github.com/nasoooor29/simple-bubble"
)
func main() {
	checkedItem, err := sb.NewFancyItemSelector("Title", sb.NewFancyItems(
		sb.NewFancyItem("item 1", "item 1 Description"),
		sb.NewFancyItem("item 2", "item 2 Description"),
		sb.NewFancyItem("item 3", "item 3 Description"),
		sb.NewFancyItem("item 4", "item 4 Description"),
		sb.NewFancyItem("item 5", "item 5 Description"),
		sb.NewFancyItem("item 6", "item 6 Description"),
	))
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("checkedItem.Title(): %v\n", checkedItem.Title())
	fmt.Printf("checkedItem.Description(): %v\n", checkedItem.Description())
}
```

### textInput
The textInput component is a simple text input field that allows the user to enter text.

quick example:

```go
package main

import (
    "fmt"

    sb "github.com/nasoooor29/simple-bubble"
)
func main() {
	text, err := sb.NewTextInput("Enter your name", "John Doe")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("text: %v\n", text)
}
```

## License

This module is released under the [GLWT License](https://raw.githubusercontent.com/nasoooor29/simple-bubble/main/LICENSE).

