package api

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/database"
	"github.com/gofiber/fiber/v2"
)

// Return all Events as JSON
func GetAllEvents(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Events []models.Event
		if response := db.Find(&Events); response.Error != nil {
			panic("Error occurred while retrieving Events from the database: " + response.Error.Error())
		}
		// Match user to events
		for index, Event := range Events {
			if Event.UserID != 0 {
				User := new(models.User)
				if response := db.Find(&User, Event.UserID); response.Error != nil {
					panic("An error occurred when retrieving the role: " + response.Error.Error())
				}
				if User.ID != 0 {
					Events[index].User = *User
				}
			}
		}
		err := ctx.JSON(Events)
		if err != nil {
			panic("Error occurred when returning JSON of Events: " + err.Error())
		}
		return err
	}
}

// Return a single Event as JSON
func GetEvent(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		Event := new(models.Event)
		id := ctx.Params("id")
		if response := db.Find(&Event, id); response.Error != nil {
			panic("An error occurred when retrieving the Event: " + response.Error.Error())
		}
		if Event.ID == 0 {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				panic("Cannot return status not found: " + err.Error())
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				panic("Error occurred when returning JSON of a role: " + err.Error())
			}
			return err
		}
		// Match role to Event
		if Event.UserID != 0 {
			User := new(models.User)
			if response := db.Find(&User, Event.UserID); response.Error != nil {
				panic("An error occurred when retrieving the role: " + response.Error.Error())
			}
			if User.ID != 0 {
				Event.User = *User
			}
		}
		err := ctx.JSON(Event)
		if err != nil {
			panic("Error occurred when returning JSON of a Event: " + err.Error())
		}
		return err
	}
}

// Add a single Event to the database
func AddEvent(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		Event := new(models.Event)

		if err := ctx.BodyParser(Event); err != nil {
			panic("An error occurred when parsing the new Event: " + err.Error())
		}
		if response := db.Create(&Event); response.Error != nil {
			panic("An error occurred when storing the new Event: " + response.Error.Error())
		}
		// Match role to Event
		if Event.UserID != 0 {
			User := new(models.User)
			if response := db.Find(&User, Event.UserID); response.Error != nil {
				panic("An error occurred when retrieving the role" + response.Error.Error())
			}
			if User.ID != 0 {
				Event.User = *User
			}
		}
		err := ctx.JSON(Event)
		if err != nil {
			panic("Error occurred when returning JSON of a Event: " + err.Error())
		}
		return err
	}
}

// Edit a single Event
func EditEvent(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		EditEvent := new(models.Event)
		Event := new(models.Event)
		if err := ctx.BodyParser(EditEvent); err != nil {
			panic("An error occurred when parsing the edited Event: " + err.Error())
		}
		if response := db.Find(&Event, id); response.Error != nil {
			panic("An error occurred when retrieving the existing Event: " + response.Error.Error())
		}
		// Event does not exist
		if Event.ID == 0 {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				panic("Cannot return status not found: " + err.Error())
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				panic("Error occurred when returning JSON of a Event: " + err.Error())
			}
			return err
		}

		/*
			Event.Name = EditEvent.Name

			Event.Email = EditEvent.Email
			Event.RoleID = EditEvent.RoleID
		*/
		// Match role to Event
		if Event.UserID != 0 {
			User := new(models.User)
			if response := db.Find(&User, Event.UserID); response.Error != nil {
				panic("An error occurred when retrieving the role" + response.Error.Error())
			}
			if User.ID != 0 {
				Event.User = *User
			}
		}
		// Save Event
		db.Save(&Event)

		err := ctx.JSON(Event)
		if err != nil {
			panic("Error occurred when returning JSON of a Event: " + err.Error())
		}
		return err
	}
}

// Delete a single Event
func DeleteEvent(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var Event models.Event
		db.Find(&Event, id)
		if response := db.Find(&Event); response.Error != nil {
			panic("An error occurred when finding the Event to be deleted" + response.Error.Error())
		}
		db.Delete(&Event)

		err := ctx.JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})
		if err != nil {
			panic("Error occurred when returning JSON of a Event: " + err.Error())
		}
		return err
	}
}
