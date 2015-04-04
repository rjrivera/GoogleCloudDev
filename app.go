package basic

import( 
     "net/http"
	 "appengine"
	)
func init() {
      http.HandleFunc("/notice/ablelana-gcp-in", notificationHandler)
}

func notificationHandler(w http.ResponseWriter, r *http.Request)
     c := appengine.NewContext(r)

     body, err := ioutil.ReadAll(r.Body)
     if err != nil {
          http.Error(w, "oops!", http.StatusInternalServerError)
          c.Errorf("read all: %v", err)
          return
     }

     c.Infof("request: %s", body)
}