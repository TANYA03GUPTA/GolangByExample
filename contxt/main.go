package main

import (
	"context"
	"net/http"
	"github.com/google/uuid"
)

func main(){
	hlwhandler := http.HandlerFunc(HelloWorld)
	http.Handle("/welcome", injectMsgID(hlwhandler))
	http.ListenAndServe(":5621",nil)

}

func HelloWorld(w http.ResponseWriter, r *http.Request){
	msgID := ""
    if m := r.Context().Value("msgId"); m != nil {
        if value, ok := m.(string); ok {
            msgID = value
        }
    }
    w.Header().Add("msgId", msgID)
    w.Write([]byte("Hello, world"))
}

func inejctMsgID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        msgID := uuid.New().String()
        ctx := context.WithValue(r.Context(), "msgId", msgID)
        req := r.WithContext(ctx)
        next.ServeHTTP(w, req)
        
    })
}