package list

import "fmt"

// реализация односвязного списка
type listElement struct {
	data int
	next *listElement
}

type linkedList struct {
	len  int
	head *listElement
}

// create new list
func InitList() *linkedList {
	return &linkedList{}
}

func (l *linkedList) PushFront(data int) {
	listElement := &listElement{
		data: data,
	}
	if l.head == nil {
		l.head = listElement
	} else {
		listElement.next = l.head
		l.head = listElement
	}
	l.len++
	return
}

func (l *linkedList) PushBack(data int) {
	listElement := &listElement{
		data: data,
	}
	if l.head == nil {
		l.head = listElement
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = listElement
	}
	l.len++
	return
}

func (l *linkedList) Remove(data int) error {
	if l.head == nil {
		return fmt.Errorf("List is empty")
	} else {
		var prev, next *listElement
		current := l.head
		for current.data != data {
			prev = current
			if current.next != nil {
				current = current.next
			} else {
				return fmt.Errorf("Item not found")
			}
			next = current.next
		}
		if prev != nil {
			prev.next = next
		} else {
			l.head = current.next
		}
	}
	l.len--
	return nil
}

// данные для проверки
// func main() {
// 	newList := InitList()
// 	newList.PushBack(1)
// 	newList.PushFront(22)
// 	newList.PushBack(25)

// 	err := newList.Remove(25)
// 	fmt.Println(err)
// 	err = newList.Remove(22)
// 	fmt.Println(err)
// 	err = newList.Remove(1)
// 	fmt.Println(err)
// 	err = newList.Remove(0)
// 	fmt.Println(err)
// 	err = newList.Remove(253)
// 	fmt.Println(err)

// 	fmt.Println(newList.len, newList.head)

// 	current := newList.head
// 	fmt.Println(current)
// 	// for current.next != nil {
// 	// 	current = current.next
// 	// 	fmt.Println(current.data, current.next)
// 	// }

// }
