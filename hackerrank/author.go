package main

import (
    "net/http"
    "io/ioutil"
    "bufio"
    "fmt"
    "io"
    "strings"
    "encoding/json"
    // "os"
)

type UserData struct {
    About string
}

type User struct {
    Data []UserData
}

type ArticleData struct {
    Title string
    Story_title string
}

type Article struct {
    Data []ArticleData
    Total int
}


func getAuthorHistory(author string) []string {
    history := []string{}

    client := &http.Client{}
    userurl := string("https://jsonmock.hackerrank.com/api/article_users?username=" + author)
    req,_ := http.NewRequest("GET", userurl, nil)
    w := req.URL.Query()
    req.URL.RawQuery = w.Encode()
    // fmt.Println(req.URL.String())
    response, err := client.Do(req)
    if err != nil {
        fmt.Printf("Failed")
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        jsonuserdata := string(data)
        // fmt.Print(jsonuserdata)
        var m User
        if err := json.Unmarshal([]byte(jsonuserdata), &m); err != nil {
            panic(err)
        }

        if len(m.Data[0].About) > 0 {
          about := TrimSuffix(m.Data[0].About, "")
          history = append(history, about)
        }
    }



authurl := string("https://jsonmock.hackerrank.com/api/articles?author=" + author)
req2 , _ := http.NewRequest("GET", authurl, nil)
w2 := req2.URL.Query()
req2.URL.RawQuery = w2.Encode()
// fmt.Println(req2.URL.String())
response2, err := client.Do(req2)
if err != nil {
    fmt.Printf("Failed")
} else {
    data, _ := ioutil.ReadAll(response2.Body)
    jsonuserdata2 := string(data)
    // fmt.Print(jsonuserdata2)
    var a Article
    if err := json.Unmarshal([]byte(jsonuserdata2), &a); err != nil {
        panic(err)
    }

    var total = a.Total
    // fmt.Println(total)

    for i := 0; i < total; i++ {
      if len(a.Data[i].Title) > 0 {
        title := TrimSuffix(a.Data[i].Title, "")
        history = append(history, title)
      } else  {
        history = append(history, a.Data[i].Story_title)
      }
    }

    fmt.Println(history)
}

    return history
}

func TrimSuffix(s, suffix string) string {
    if strings.HasSuffix(s, suffix) {
        s = s[:len(s)-len(suffix)]
    }
    return s
}

func main() {
  getAuthorHistory("epaga")
  getAuthorHistory("saintamh")
  getAuthorHistory("patricktomas")
    //
    //
    // reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
    //
    // stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    // checkError(err)
    //
    // defer stdout.Close()
    //
    // writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)
    //
    // author := readLine(reader)
    //
    // result := getAuthorHistory(author)
    //
    // for i, resultItem := range result {
    //     fmt.Fprintf(writer, "%s", resultItem)
    //
    //     if i != len(result) - 1 {
    //         fmt.Fprintf(writer, "\n")
    //     }
    // }
    //
    // fmt.Fprintf(writer, "\n")
    //
    // writer.Flush()
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
