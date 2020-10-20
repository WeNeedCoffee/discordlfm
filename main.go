package main

import (
	"os"
	"unicode"
	"strings"
	"fmt"
	"bufio"
)

func redactVowels(s string) string{
	vowels := []byte{'a','e','i','o','u'}
	var res strings.Builder

	first := true
	for i := range s{
		found := false
		if first{
			for _,v := range vowels{
				if s[i]==v {
					res.WriteString("*")
					found=true
					first=false
					break
				}
			}
		}
		if !found{
			res.WriteByte(s[i])
		}
	}
	return res.String()
}

func censure(s string) string{
    for i := 0; i < len(s); i++ {
        if s[i] > unicode.MaxASCII {
            return s
        }
    }
	fi, err := os.Open("swears.txt")
	if err != nil{
		panic(err)
	}

	scanner:=bufio.NewScanner(fi)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		swear := scanner.Text()
		fmt.Println(swear)
		strings.ReplaceAll(swear,"\n","")
		for strings.Contains(strings.ToLower(s),swear) {
			idx := strings.Index(strings.ToLower(s),swear)
			s = s[:idx] + redactVowels(s[idx:idx+len(swear)]) + s[idx+len(swear):]
		}
	}
	if err:= scanner.Err(); err!= nil{
		panic(err)
	}
	fi.Close()
    return s
}


func main(){
	s := censure("dick")
	fmt.Println(s)

}
