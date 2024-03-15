package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

type Employee struct {
	Id     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

var mg MongoInstance

const dbName = "fiber-hrns"
const basePath = "/api/v1/employees"
const mongoURI = "mongodb://localhost:27017/" + dbName

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get(basePath, func(c *fiber.Ctx) error {
		query := bson.D{}
		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}
		var employees []Employee = make([]Employee, 0)
		if err := cursor.All(c.Context(), &employees); err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())

		}
		return c.JSON(employees)
	})

	app.Post(basePath, func(c *fiber.Ctx) error {
		collection := mg.Db.Collection("employees")
		employee := new(Employee)
		if err := c.BodyParser(employee); err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		employee.Id = ""

		insertionResult, err := collection.InsertOne(c.Context(), employee)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
		createdRecord := collection.FindOne(c.Context(), filter)

		createdEmployee := &Employee{}
		createdRecord.Decode(createdEmployee)
		return c.Status(http.StatusCreated).JSON(createdEmployee)

	})

	app.Put(basePath+"/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		employeeId, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		employee := new(Employee)
		if err := c.BodyParser(employee); err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		query := bson.D{{Key: "_id", Value: employeeId}}
		update := bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{Key: "name", Value: employee.Name},
					{Key: "age", Value: employee.Age},
					{Key: "salary", Value: employee.Salary},
				},
			},
		}

		err = mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()

		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(http.StatusNotFound)
			}
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		employee.Id = id

		return c.JSON(employee)
	})

	app.Delete(basePath+"/:id", func(c *fiber.Ctx) error {
		employeeId, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		query := bson.D{{Key: "_id", Value: employeeId}}
		result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), &query)
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		if result.DeletedCount < 1 {
			return c.SendStatus(http.StatusNotFound)
		}

		return c.Status(http.StatusOK).JSON("record deleted")
	})

	log.Fatal(app.Listen(":8080"))
}

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err != nil {
		return err
	}

	errCtx := client.Connect(ctx)
	db := client.Database(dbName)

	if errCtx != nil {
		return errCtx
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}
