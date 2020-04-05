package main

import (
	"NanairoMegane/sqlboiler_test/models"

	"github.com/volatiletech/null"
)

type UserModel struct {
	models.User
}

type UserSliceModel struct {
	models.UserSlice
}

type UserCustomeModel struct {
	UserID int    `boil:"user_id"`
	Name   string `boil:"name"`
}

type UserCustomeSliceModel struct {
	USlice []*UserCustomeModel
}

type UserAndDivision struct {
	UserID       int    `boil:"user_id"`
	UserName     string `boil:"user_name"`
	DivisionID   int    `boil:"division_id"`
	DivisionName string `boil:"division_name"`
}

type UserAndDivisionSliceModel struct {
	UDSlice []*UserAndDivision
}

type UserAndBranch struct {
	UserID     int         `boil:"user_id"`
	UserName   string      `boil:"user_name"`
	BranchID   null.Int    `boil:"branch_id"`
	BranchName null.String `boil:"branch_name"`
}

type UserAndBranchSliceModel struct {
	UBSlice []*UserAndBranch
}
