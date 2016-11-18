package main

import ("fmt"
        "io/ioutil"
        "net/http"
)

type Page struct {
  Title string
  Body []byte
}

func filename (page *Page) string{
  return page.Title + ".txt"
}


func (self *Page) save() error {
  filename := filename(self)
  return ioutil.WriteFile(filename, self.Body, 0600)
}


func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, err
}

func handler(w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/view/"):]
  page, err := loadPage(title)
  if err != nil {
    fmt.Fprintf(w, "<h1>Oops Sorry</h1>")
  } else {
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
  }
}

func main() {
  http.HandleFunc("/view/", handler)
  http.ListenAndServe(":8000", nil)
}





















