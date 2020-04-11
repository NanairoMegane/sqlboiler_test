package main

import (
	"NanairoMegane/sqlboiler_test/models"
	"context"
	"fmt"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// practical_test is a function to test for practical and complex query with sqlboiler.
func practical_test() error {

	if err := prepare_data(); err != nil {
		return err
	}

	if err := built_query(); err != nil {
		return err
	}

	if err := custom_query(); err != nil {
		return err
	}

	if err := raw_query(); err != nil {
		return err
	}

	if err := custom_struct(); err != nil {
		return err
	}

	if err := inner_join(); err != nil {
		return err
	}

	if err := outer_join(); err != nil {
		return err
	}

	if err := delete_data(); err != nil {
		return err
	}

	return nil
}

// built_query is a function to execute query with userQuery receiver.
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

// custom_query is a function to execute query with Query receiver.
func custom_query() error {

	var result = new(UserSliceModel)

	var queries []qm.QueryMod
	{
		queries = append(queries, qm.Select("*"))
		queries = append(queries, qm.From("User"))
		queries = append(queries, qm.Where("age = ?", 27))
	}

	err := models.NewQuery(queries...).Bind(context.Background(), DB, &result.UserSlice)

	if err == nil {
		fmt.Println("--- custom query ---")
		for _, user := range result.UserSlice {
			fmt.Printf("result : \n  user_id: %d\n  name: %s\n  age: %d\n",
				user.UserID, user.Name, user.Age)
		}
	}
	fmt.Println()

	return err
}

// raw_query is a function to execute query with raw query.
func raw_query() error {

	var result = new(UserSliceModel)

	var query = `
	SELECT 
		user_id,
		name,
		age
	FROM 
		user
	WHERE
		age = 25
	`

	err := models.NewQuery(qm.SQL(query)).Bind(context.Background(), DB, &result.UserSlice)
	if err == nil {
		fmt.Println("--- raw query ---")
		for _, user := range result.UserSlice {
			fmt.Printf("result : \n  user_id: %d\n  name: %s\n  age: %d\n",
				user.UserID, user.Name, user.Age)
		}
	}
	fmt.Println()

	return err
}

// custom_struct is a function to execute query with Query receiver and bind result to custom struct.
func custom_struct() error {

	var result = new(UserCustomSliceModel)

	var queries []qm.QueryMod
	{
		queries = append(queries, qm.Select("user_id, name"))
		queries = append(queries, qm.From("User"))
		queries = append(queries, qm.Where("user_id = ?", 3))
	}

	err := models.NewQuery(queries...).Bind(context.Background(), DB, &result.USlice)
	if err == nil {
		fmt.Println("--- custom struct ---")
		for _, user := range result.USlice {
			fmt.Printf("result : \n  user_id: %d\n  name: %s\n",
				user.UserID, user.Name)
		}
	}
	fmt.Println()

	return err
}

// inner_join is a function to execute inner join query with Query receiver.
// When use inner join, you need prepare custom struct.
func inner_join() error {

	var result = new(UserAndDivisionSliceModel)

	var queries []qm.QueryMod
	{
		queries = append(queries, qm.Select("U.user_id AS user_id, U.name AS user_name, D.division_id AS division_id, D.name AS division_name"))
		queries = append(queries, qm.From("User AS U"))
		queries = append(queries, qm.InnerJoin("Division AS D ON U.user_id = D.user_id"))
	}

	err := models.NewQuery(queries...).Bind(context.Background(), DB, &result.UDSlice)
	if err == nil {
		fmt.Println("--- inner join ---")
		for _, ud := range result.UDSlice {
			fmt.Printf("result : \n  user_id: %d\n  user_name: %s\n  division_id: %d\n  division_name: %s\n",
				ud.UserID, ud.UserName, ud.DivisionID, ud.DivisionName)
		}
	}
	fmt.Println()

	return err
}

// outer_join is a function to execute inner join query with raw query.
// When use outer join, you need prepare custom struct which is constituted nullable variable.
func outer_join() error {
	var result = new(UserAndBranchSliceModel)

	var outer_join_query = `
		SELECT 
			U.user_id AS user_id, 
			U.name AS user_name, 
			B.branch_id AS branch_id, 
			B.name AS branch_name
		FROM 
			user AS U
		LEFT OUTER JOIN 
			Branch AS B 
		ON 
			U.user_id = B.user_id`

	err := models.NewQuery(qm.SQL(outer_join_query)).Bind(context.Background(), DB, &result.UBSlice)
	if err == nil {
		fmt.Println("--- outer join ---")
		for _, ub := range result.UBSlice {
			fmt.Printf("result : \n  user_id: %d\n  user_name: %s\n  branch_id: %d\n  branch_name: %s\n",
				ub.UserID, ub.UserName, ub.BranchID.Int, ub.BranchName.String)
		}
	}
	fmt.Println()

	return err
}

// prepare_data is a function to prepare test data.
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
	if err := (&models.Division{DivisionID: 1, Name: "Akita", UserID: 1}).Insert(context.Background(), DB, boil.Infer()); err != nil {
		return err
	}
	if err := (&models.Branch{BranchID: 1, Name: "Oodate", UserID: null.IntFrom(1)}).Insert(context.Background(), DB, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// delete_data is a function to delete test data.
func delete_data() error {

	if _, err := models.Divisions().DeleteAll(context.Background(), DB); err != nil {
		return err
	}
	if _, err := models.Branches().DeleteAll(context.Background(), DB); err != nil {
		return err
	}
	if _, err := models.Users().DeleteAll(context.Background(), DB); err != nil {
		return err
	}

	return nil
}
