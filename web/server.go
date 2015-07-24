/*
 web package example
 아래 소스를 참조하여 테스트 프로그램 작성
 source : http://www.alexedwards.net/blog/a-recap-of-request-handling

 ServeMuxes: http request router(or muxtiplexor)
 Handlers: writing response headers and bodies

 ServeHTTP(http.ResponseWriter, *http.Request)

*/
package web

import (
	"net/http"
	"log"
	"time"
)


func timeHandler(format string) http.Handler {
	// closure function
  fn := func(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(format)
    w.Write([]byte("The time is: " + tm))
  }
  return http.HandlerFunc(fn)
}

func main() {
	mux := http.NewServeMux()

	var format string = time.RFC1123
	th := timeHandler(format)

 	mux.Handle("/", th)

 	log.Println("Listening....")

	http.ListenAndServe(":8000", mux)
}