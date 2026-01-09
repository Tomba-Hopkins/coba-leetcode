# Convert Binary num to Int with LinkedList

## Desc
so we have a Node in here :

### [1] [0] [1]

and we will go through that like this

**1**01

1**0**1

10**1**

## Result

output must be 5

calculate this binary with normal method like this :

101 -> 1 + 0 + 4 -> 5


# Implement with Go

## we have struct in here 

```go
type ListNode1 struct {
	Val  int
	Next *ListNode1
}
```


## i use this method to print it 

```go 
func buatinNode(arr []int) *ListNode1 {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode1{Val: arr[0]}
	curr := head
	for i := 1; i < len(arr); i++ {
		curr.Next = &ListNode1{Val: arr[i]}
		curr = curr.Next
	}
	return head
}

func printNodeNyah(h *ListNode1) {
	for h != nil {
		fmt.Print(h.Val, " ")
		h = h.Next
	}
	fmt.Println()
}

func main() {
	a := []int{1, 0, 1}
	nod := buatinNode(a)
	printNodeNyah(nod)
	fmt.Println(ConvertBinNumInALinkedListToInt(nod))
}
```

## And we have this function to Implement that convert things
```go
func ConvertBinNumInALinkedListToInt(head *ListNode1) int {
	r := 0
	for head != nil {
		r = r*2 + head.Val
		head = head.Next
	}
	return r
}
```

## main formula :
```pseudo
head = *Node
r = 0
while head != nil 
    r = r * 2 + head.Value
    head = head.Next
```

## so it will be like :
```pseudo
input: 1, 0, 1
- formula: r * 2 + head.Value
-> index1: 0 * 2 + 1
-> index2: 1 * 2 + 0
-> index3: 2 * 2 + 1 -> 5
output: 5
```

## Key insight
### about understanding how binary numbers are constructed.

Linked List only forces us to:
- read from left to right
- without indexing
- without converting to string


### Thats it thanks for read, have fun