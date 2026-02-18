package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/note"
)

func main(){
	title,content :=makenote()

	userNote,err :=note.New(title,content)

	if (err!=nil){
		fmt.Println(err)
		return
	}
	userNote.Display()
	err =userNote.Save()

	if err!=nil{
		fmt.Println("Error saving the file")
		return 
	}
	fmt.Println("File saved successfully")
}
func makenote()(string,string){
	title :=getuserinput("Note title: ")
	
	content :=getuserinput("Note content: ")

	return title,content

}
func getuserinput(prompt string) (string) {
	fmt.Print(prompt)
	reader :=bufio.NewReader(os.Stdin)

	text,err :=reader.ReadString('\n')

	if err!=nil{
		return ""
	} 
	text =strings.TrimSuffix(text,"\n")
	text =strings.TrimSuffix(text,"\r")
	return text
}