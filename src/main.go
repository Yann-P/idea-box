package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"./sse"
	"github.com/go-pg/pg"
)

type Entity struct {
	Id      uint64 `json:"id"`
	Name    string `json:"name"`
	Support uint64 `json:"support"`
	Against uint64 `json:"against"`
}

type SSEEvent struct {
	Type string
	Data interface{}
}

func toJSON(v interface{}) []byte {
	json, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return json
}

func sendSSE(broker *sse.Broker, event string, data interface{}) {
	broker.Notifier <- toJSON(&SSEEvent{event, data})
}

func addEntity(name string, db *pg.DB, broker *sse.Broker) {
	entity := &Entity{
		Name: name,
	}

	err := db.Insert(entity)

	if err != nil {
		panic(err)
	}

	go sendSSE(broker, "add", entity)
}

func main() {
	broker := sse.NewServer()
	fs := http.FileServer(http.Dir("/app/public"))
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("POSTGRES_ADDRESS"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	})
	defer db.Close()

	http.Handle("/", fs)
	http.HandleFunc("/api/add", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		if len(name) > 2 {
			addEntity(r.FormValue("name"), db, broker)
		}
	})

	http.HandleFunc("/api/list", func(w http.ResponseWriter, r *http.Request) {
		var entities []Entity
		err := db.Model(&entities).Order("id").Select()
		if err != nil {
			panic(err)
		}
		if entities == nil {
			fmt.Fprintf(w, "[]")
			return
		}
		w.Write(toJSON(&entities))
	})

	http.HandleFunc("/api/vote", func(w http.ResponseWriter, r *http.Request) {
		idRaw := r.FormValue("id")

		id, err := strconv.ParseUint(idRaw, 10, 64)
		if err != nil {
			fmt.Fprintln(w, "Invalid id")
			return
		}

		voteRaw := r.FormValue("vote")
		vote, err := strconv.ParseBool(voteRaw)
		if err != nil {
			fmt.Fprintln(w, "Invalid vote")
			return
		}

		entity := &Entity{Id: id}
		err = db.Select(entity)
		if err != nil {
			fmt.Fprintln(w, "Invalid entity")
			return
		}

		if vote {
			entity.Support++
		} else {
			entity.Against++
		}

		err = db.Update(entity)

		if err != nil {
			panic(err)
		}

		go sendSSE(broker, "vote", map[string]interface{}{
			"id":      id,
			"support": entity.Support,
			"against": entity.Against,
			"vote":    vote,
		})
	})

	http.Handle("/api/sse", broker)
	http.ListenAndServe(":80", nil)
}
