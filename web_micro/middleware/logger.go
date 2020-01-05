package middleware

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
)

func LogHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		x, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		rec := httptest.NewRecorder()
		fn(rec, r)
		err = WriteHTTP("access.log", []string{fmt.Sprintf("%q - response %q", x, rec.Body)})
		if err != nil {
			fmt.Println("Server unable to generate and write to file")
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
	}
}


func WriteHTTP(outfile string , data []string) error {

	f, err := os.OpenFile(outfile, os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
	defer f.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return err
		}
	}
	return err
}

func WriteString(outfile string , data []string) error {

	f, err := os.Create(outfile)
	defer f.Close()

	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			return err
		}
	}
	return err
}