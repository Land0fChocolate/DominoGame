package main

import(
	"time"
)

func main(){
	var text = `I need a front door for my hall,
The replacement I bought was too tall.
So I hacked it and chopped it,
And carefully lopped it,
And now the dumb thing is too small.`

	var wordList map[word]num
	spltStr := strings.Split(text, " ")
	for _, word := range spltStr{
		if num, ok := wordList[word]; ok{
			num++
		} else{
			worldList[word] = 1
		}
		
	}
}