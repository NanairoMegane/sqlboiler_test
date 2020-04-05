package main

import (
	"NanairoMegane/sqlboiler_test/models"
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// practical_test is a function to test for practical and complex query with sqlboiler.
func practical_test() error {

	if err := prepare_data(); err != nil {
		return err
	}

	if err := custome_query(); err != nil {
		return err
	}

	if err := built_query(); err != nil {
		return err
	}

	if err := custome_struct(); err != nil {
		return err
	}

	if err := delete_data(); err != nil {
		return err
	}

	return nil
}

func built_query() error {

	result, err := models.Users(qm.Where("age = ?", 27)).All(context.Background(), DB)

	if err != nil {
		return err
	}

	fmt.Println("--- built query ---")
	for _, user := range result {
		fmt.Printf("result : \n  user_id: %d\n  name: %s\n  age: %d\n",
			user.UserID, user.Name, user.Age)
	}
	fmt.Println()

	return nil
}

func custome_query() error {

	var result = new(UserSliceModel)

	var queries []qm.QueryMod
	{
		queries = append(queries, qm.Select("*"))
		queries = append(queries, qm.From("User"))
		queries = append(queries, qm.Where("age = ?", 27))
	}

	err := models.NewQuery(queries...).Bind(context.Background(), DB, &result.UserSlice)

	if err == nil {
		fmt.Println("--- custome query ---")
		for _, user := range result.UserSlice {
			fmt.Printf("result : \n  user_id: %d\n  name: %s\n  age: %d\n",
				user.UserID, user.Name, user.Age)
		}
	}
	fmt.Println()

	return err
}

func custome_struct() error {

	var result = new(UserCustomeSliceModel)

	var queries []qm.QueryMod
	{
		queries = append(queries, qm.Select("user_id, name"))
		queries = append(queries, qm.From("User"))
		queries = append(queries, qm.Where("user_id = ?", 3))
	}

	err := models.NewQuery(queries...).Bind(context.Background(), DB, &result.USlice)
	if err == nil {
		fmt.Println("--- custome struct ---")
		for _, user := range result.USlice {
			fmt.Printf("result : \n  user_id: %d\n  name: %s\n",
				user.UserID, user.Name)
		}
	}
	fmt.Println()

	return err
}

func inner_join() {

}

func outer_join() {

}

func prepare_data() error {

	if err := (&models.User{UserID: 1, Name: "Ren", Age: 27}).Insert(context.Background(), DB, boil.Infer()); err != nil {
		return err
	}
	if err := (&models.User{UserID: 2, Name: "Kim", Age: 25}).Insert(context.Background(), DB, boil.Infer()); err != nil {
		return err
	}
	if err := (&models.User{UserID: 3, Name: "Sasa", Age: 27}).Insert(context.Background(), DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func delete_data() error {
	if _, err := models.Users().DeleteAll(context.Background(), DB); err != nil {
		return err
	}
	if _, err := models.Divisions().DeleteAll(context.Background(), DB); err != nil {
		return err
	}
	if _, err := models.Branches().DeleteAll(context.Background(), DB); err != nil {
		return err
	}

	return nil
}
