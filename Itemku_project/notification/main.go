package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notification/connection"
	"notification/controller"
	"notification/models"

	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"github.com/go-chi/chi"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("%s", msg.Payload())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

var DB *gorm.DB

func init() {
	DB = connection.ConnectToDb()
}

func main() {
	var broker = os.Getenv("BROKER")
	var port = cast.ToInt(os.Getenv("MQTT_PORT"))
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername("emqx")
	opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	sub(client)

	r := chi.NewRouter()

	// check service
	r.Group(func(g chi.Router) {
		g.Get("/", func(w http.ResponseWriter, r *http.Request) {
			controller.ResponseApi(w, http.StatusOK, nil, "Service is running")
		})

		g.Post("/publish", func(w http.ResponseWriter, r *http.Request) {
			decoder := json.NewDecoder(r.Body)
			var datarequest models.RequestNotify
			if err := decoder.Decode(&datarequest); err != nil {
				controller.ResponseApi(w, http.StatusBadRequest, nil, "Invalid request payload")
				return
			}

			// to json string
			newRequst, err := json.Marshal(datarequest)
			if err != nil {
				controller.ResponseApi(w, http.StatusBadRequest, nil, "Invalid request payload")
				return
			}

			// Create
			DB.Create(&models.NotifItemku{IdPemesanan: datarequest.IdPemesanan, Request: string(newRequst)})

			// Publish
			publish(client, string(newRequst))

			// update is_sent
			DB.Model(&models.NotifItemku{}).Where("order_id = ?", datarequest.IdPemesanan).Update("sens", true)

			// Response
			controller.ResponseApi(w, http.StatusOK, datarequest, "Publish success")
		})

		g.Get("/notifies", func(w http.ResponseWriter, r *http.Request) {
			var data []models.NotifItemku
			DB.Find(&data)
			controller.ResponseApi(w, http.StatusOK, data, "Get notifies success")
		})
	})

	log.Println("Service running on " + os.Getenv("HOST") + ":" + os.Getenv("PORT"))

	portServer := os.Getenv("PORT")
	if portServer == "" {
		portServer = "8081"
	}

	if err := http.ListenAndServe(":"+portServer, r); err != nil {
		log.Println("Error Starting Service")
	}
}

func publish(client mqtt.Client, msg string) {
	token := client.Publish(os.Getenv("TOPIC"), 0, false, msg)
	token.Wait()
}

func sub(client mqtt.Client) {
	topic := os.Getenv("TOPIC")
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
}
