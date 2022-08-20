package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"notifikasi/config"
	"notifikasi/controllers"
	"notifikasi/models"

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
	DB = config.Connect()
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
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// subcribe to topic
	sub(client)

	// router
	r := chi.NewRouter()

	// check service
	r.Group(func(g chi.Router) {
		g.Get("/", func(w http.ResponseWriter, r *http.Request) {
			controllers.Response(w, http.StatusOK, "Service is running", nil)
		})

		g.Post("/publish", func(w http.ResponseWriter, r *http.Request) {
			decoder := json.NewDecoder(r.Body)
			var datarequest models.RequestNotify
			if err := decoder.Decode(&datarequest); err != nil {
				controllers.Response(w, http.StatusBadRequest, "Invalid request payload", nil)
				return
			}

			// to json string
			newRequst, err := json.Marshal(datarequest)
			if err != nil {
				controllers.Response(w, http.StatusBadRequest, "Invalid request payload", nil)
				return
			}

			// Create
			DB.Create(&models.Notif{IdPemesanan: datarequest.IdPemesanan, Request: string(newRequst)})

			// Publish
			publish(client, string(newRequst))

			// update is_sent
			DB.Model(&models.Notif{}).Where("order_id = ?", datarequest.IdPemesanan).Update("sens", true)

			// Response
			controllers.Response(w, http.StatusOK, "Publish success", datarequest)
		})

		g.Get("/notifies", func(w http.ResponseWriter, r *http.Request) {
			var data []models.Notif
			DB.Find(&data)
			controllers.Response(w, http.StatusOK, "Get notifies success", data)
		})
	})

	log.Println("Service running on " + os.Getenv("HOST") + ":" + os.Getenv("PORT"))

	portServer := os.Getenv("PORT")
	if portServer == "" {
		portServer = "8080"
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
