package main

import "NanairoMegane/sqlboiler_test/models"

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
