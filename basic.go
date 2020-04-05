package main

import (
	"NanairoMegane/sqlboiler_test/models"
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/boil"
)

// basic_test is a function to test for simple query with sqlboiler.
// all method act to user_id = 1. so when after function doing, data is not exist on db.
func basic_test() error {

	// Insert
	if err := basic_insert(); err != nil {
		return err
	}

	// Update
	if err := basic_update(); err != nil {
		return err
	}

	// Select
	if err := basic_select_one(); err != nil {
		return err
	}

	// Delete
	if err := basic_delete(); err != nil {
		return err
	}

	return nil
}

func basic_insert() error {

	userModel := new(UserModel)

	userModel.UserID = 1
	userModel.Name = "INSERT"
	userModel.Age = 27

	return userModel.Insert(context.Background(), DB, boil.Infer())
}

func basic_update() error {

	userModel := new(UserModel)
	userModel.UserID = 1
	userModel.Name = "Updated"
	userModel.Age = 28

	_, err := userModel.Update(context.Background(), DB, boil.Infer())
	return err
}

func basic_select_one() error {

	findedUser, err := models.FindUser(context.Background(), DB, 1)
	if err == nil {
		fmt.Printf("--- selected ---\nid: %d\nname: %s\nage: %d\n",
			findedUser.UserID, findedUser.Name, findedUser.Age)
	}

	return err
}

func basic_delete() error {

	userModel := new(UserModel)
	userModel.UserID = 1

	_, err := userModel.Delete(context.Background(), DB)

	return err
}
