package main

func getValue(n map[string]string, key string) (string, bool) {
	a, o := n[key]
	return a, o
}

func main() {
	a := true
	if a {
		println("hello world")
	}

	b := 1
	if b == 1 {
		println("неявне перетворення ( if b ) не працює")
	}

	mm := map[string]string{"firstName": "Johnny", "lastName": "Cash"}
	middleName, ok := mm["middleName"]
	println(middleName, ok)

	if firstName, ok := mm["firstName"]; ok || firstName == "Johnny" {
		middleName = "Some"
		println("firstName key exist, = ", firstName, "middleName", middleName)
		mm["middleName"] = middleName
	} else {
		println("no firstName , ", firstName)
	}
	xa, c := 0, 3
	println(xa, c)
	c, xa = xa, c
	println(xa, c)

	if firstName, ok := getValue(mm, "firstName"); !ok {
		println("no firstName")
	} else if firstName == "Johnny" {
		println("firstName is Johnny")
	} else {
		println("firstName is not Johnny")
	}

	for {
		println("нескінченний цикл")
		break
	}

	sl := []int{3, 4, 5, 6, 7, 8}
	// value := 0
	// idx := 0
	var idx int
	var value int
	// Операції з slice
	for idx < 4 {
		if idx < 2 {
			idx++
			continue
		}
		value = sl[idx]
		idx++
		println("while-style loop, idx:", idx, "value:", value)
	}

	for i := 0; i < len(sl); i++ {
		println("c-style loop", i, sl[i])
	}

	for idx := range sl {
		println("range slice by index", idx)
	}

	for idx, val := range sl {
		println("range slice by idx-value", idx, val)
	}

	// Операції з map
	for key := range mm {
		println("range map by key", key)
	}

	for key, val := range mm {
		println("range map by key-val", key, val)
	}

	for _, val := range mm {
		println("range map by val", val)
	}

	mm["firstName"] = "Vasily"
	mm["flag"] = "Ok"

	switch mm["firstName"] {
	case "Vasily", "Evgeny":
		println("switch - name is Vasily")
		// На відміну від інших мов - не переходимо в інший варіант за замовчуванням
	case "Petr":
		if mm["flag"] == "Ok" {
			break // виходимо зі switch, щоб не виконувати перехід в інший варіант
		}
		println("switch - name is Pert")
		fallthrough // Переходимо в наступний варіант
	default:
		println("switch - some other name")
	}

	// як заміна множинним if else
	switch {
	case mm["firstName"] == "Vasily":
		println("switch2 - Vasily")
	case mm["lastName"] == "Romanov":
		println("switch2 - Romanov")
	default:
		println("unknown person")
	}

	// вихід з циклу бувши всередині switch
Loop:
	for key, val := range mm {
		println("switch in loop", key, val)
		if val == "" {
			break
		}
		switch {
		case key == "firstName" && val == "Vasily":
			println("switch - break loop here")
			break Loop
		}
	}
}
