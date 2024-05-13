// func calPoints(operations []string) int {
// 	record := []int{}

// 	for _, o := range operations {
// 		if o == "C" {
// 			record = record[:len(record)-1]
// 		} else if o == "D" {
// 			record = append(record, 2*record[len(record)-1])
// 		} else if o == "+" {
// 			record = append(record, record[len(record)-2]+record[len(record)-1])
// 		} else {
// 			numb, _ := strconv.Atoi(o)
// 			record = append(record, numb)
// 		}
// 	}

// 	res := 0
// 	for _, numb := range record {
// 		res += numb
// 	}
// 	return res
// }

func calPoints(operations []string) int {
	stack := NewStack()
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "+":
			num1, num2 := stack.pop(), stack.pop()
			stack.push(num2)
			stack.push(num1)
			stack.push(num1 + num2)
		case "D":
			num := stack.peek()
			stack.push(num * 2)
		case "C":
			stack.pop()
		default:
			num, err := strconv.Atoi(operations[i])
			if err != nil {
				panic(err)
			}
			stack.push(num)
		}
	}

	sum := 0
	for stack.head != nil {
		sum += stack.pop()
	}
	return sum
}

type Stack struct {
	head *ListNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (stack *Stack) push(value int) {
	newNode := &ListNode{
		Val:  value,
		Next: stack.head,
	}

	stack.head = newNode
}

func (stack *Stack) pop() int {
	if stack.head == nil {
		panic("empty stack")
	} else {
		res := stack.head.Val
		stack.head = stack.head.Next
		return res
	}
}

func (stack *Stack) peek() int {
	if stack.head != nil {
		return stack.head.Val
	} else {
		panic("empty stack")
	}
}