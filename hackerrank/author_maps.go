package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)


/*
 * Complete the 'getAuthorHistory' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING author as parameter.
 *
 * Base urls:
 *   https://jsonmock.hackerrank.com/api/article_users?username=
 *   https://jsonmock.hackerrank.com/api/articles?author=
 *
 */
 
import (
  "log"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func getAuthorHistory(author string) []string {
    
    history := []string{}
        
    response, err := http.Get("https://jsonmock.hackerrank.com/api/article_users?username="+author)
    if err != nil {
        log.Fatalln(err)
    }
    body, err := ioutil.ReadAll(response.Body)


    var response_body map[string]interface{}

    
    if err := json.Unmarshal(body, &response_body); err != nil {
        log.Fatalln(err)
    }
    
    //fmt.Println(response_body)
    
    data := response_body["data"].([]interface{})
    
    about := data[0].(map[string]interface{})["about"]
    
    if about != nil {
        if len(about.(string)) > 0 {
            history = append(history, about.(string)) 
        }
    }
    
    page_count := 1
    
    for page_count != 0 {
        
        articles_response, err := http.Get(fmt.Sprintf("https://jsonmock.hackerrank.com/api/articles?author=%v&page=%d", author,page_count))
        if err != nil {
            log.Fatalln(err)
        }
        a_body, err := ioutil.ReadAll(articles_response.Body)
        
        var a_response_body map[string]interface{}
        
        if err := json.Unmarshal(a_body, &a_response_body); err != nil {
            log.Fatalln(err)
        }
        
        a_data := a_response_body["data"].([]interface{})
        
        for k := range a_data {
            item := a_data[k].(map[string]interface{})
            
            if item["title"] != nil {
                title := item["title"].(string)
                if len(title) > 0 {
                    history = append(history, title)
                }
            }
            if item["story_title"] != nil {
                story_title := item["story_title"].(string)
                if len(story_title) > 0 {
                    history = append(history, story_title)
                }
            }
        }
        
        total_pagescount := int(a_response_body["total_pages"].(float64))
        fmt.Println(total_pagescount)
        
        if total_pagescount < page_count {
            
            page_count += 1
            
        }
        if total_pagescount == page_count {
            page_count = 0
        }
        
    }
    
    for i := range history {
        fmt.Println(history[i])
    }
    
    return history

}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    author := readLine(reader)

    result := getAuthorHistory(author)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%s", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
