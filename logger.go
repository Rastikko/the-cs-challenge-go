package main

import (
	"log"
	"net/http"
	"time"
	"fmt"
)

// func Logger(inner http.Handler, name string) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()
//
// 		inner.ServeHTTP(w, r)
//
// 		log.Printf(
// 			"%s\t%s\t%s\t%s",
// 			r.Method,
// 			r.RequestURI,
// 			name,
// 			time.Since(start),
// 		)
// 	})
// }



func Logger(fn http.HandlerFunc, name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
		fmt.Println("HEY!")

		fn(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
  }
}
