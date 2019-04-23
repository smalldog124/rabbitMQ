package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ", err)
	}
	defer connection.Close()

	chanel, err := connection.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel", err)
	}
	defer chanel.Close()

	q, err := chanel.QueueDeclare(
		"hello", // name
		false,   // durable เป็น queue ถาวร หรือไม่
		false,   // delete when unused ถ้าของไม่มีคนถึงข้อมูลจาก qeue จะให้ลบทิ้งเลยหรือไม่
		false,   // exclusive ถ้าคนที่ส่ง message หยุดแล้วจะลบ qeue เลยหรือไม่
		false,   // no-wait ไม่สนว่ามี qeue ใน server อยู่หรื่อเปล่าก็ใช้งานเลย
		nil,     // arguments เป็น optional เช่น จำกัดเวลาของ message, ความจุของ queue
	)
	if err != nil {
		log.Fatal("Failed to declare a queue", err)
	}

	body := "Hello Smalldoc"
	err = chanel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Fatal("Failed to publish a message", err)
	}
}
