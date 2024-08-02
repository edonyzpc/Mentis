package chat

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	mentis "github.com/edonyzpc/Mentis/pkg"
	"github.com/sashabaranov/go-openai"
)

/*
chatbot format:
```

	{chat banner}:
	> {user input}
	< ------------
	< {chat response}
	< ------------
	>

```
example: 一个英文翻译的聊天机器人
```
请输入要翻译的内容:
> 好你，我是一个正在学习LLM编程的新人，很高兴可以使用通义千问大模型。
< ------------
<
< ------------
>
```
*/
func CreateChatBot(chatBanner string, mentis *mentis.Mentis, system openai.ChatCompletionMessage) {
	var sentence string
	fmt.Println(chatBanner)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		messagesTrans := []openai.ChatCompletionMessage{
			system,
		}
		sentence = scanner.Text()
		userMessageTrans := openai.ChatCompletionMessage{
			Role:    "user",
			Content: sentence + "EOF",
		}
		messagesTrans = append(messagesTrans, userMessageTrans)
		stream, err := CreateStream(mentis, "qwen-max", messagesTrans)
		if err != nil {
			log.Fatal("create stream chat error: ", err)
		}
		defer stream.Close()

		fmt.Println("< ------------")
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				//fmt.Println("\nStream finished")
				break
			}

			if err != nil {
				//fmt.Printf("\nStream error: %v\n", err)
				break
			}

			if len(response.Choices) > 0 {
				fmt.Printf(response.Choices[0].Delta.Content)
			}
		}
		fmt.Println("< ------------")
		fmt.Print("> ")
	}
}
