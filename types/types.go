package types

import (
)

type Pop struct {
	id int
	jobType string
	workforceSize int
	dependantsSize int
	workplaceId int
}

type Good struct {
	id int
	name string
	amount float32
	price float32
}

type Building struct {
	id int
	levels int
	buildingType string
	stateId int
	pms string
	inputGoods []Good
	outputGoods []Good
	profit float32
}

type Province struct{
	id int
	stateId int
	buildings []Building
}


type State struct{
	id int
	name string
	provinces []Province 
}

type Country struct {
	id int
	marketId int
	provinces []Province
}

type SaveFile struct {
	playerCountryId int
	
}