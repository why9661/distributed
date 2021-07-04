package log

import (
	"io/ioutil"
	stdlog "log"
	"net/http"
	"os"
)

var log *stdlog.Logger

type logfile string

func (f logfile) Write(data []byte) (int, error) {
	file, err := os.OpenFile(string(f), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	return file.Write(data)
}

func NewLog(destination string) {
	log = stdlog.New(logfile(destination), "Go ", stdlog.LstdFlags)
}

func RegisterHandler() {
	http.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(r.Body)
			if err != nil || len(msg) == 0 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			write(msg)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func write(data []byte) {
	log.Printf("%v\n", string(data))
}
