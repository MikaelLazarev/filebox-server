/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package core

type (
	Location [2]float64

	Box struct {
		BaseModel  `bson:",inline"`
		IPFSHash   string   `json:"ipfsHash"`
		Name       string   `json:"name"`
		Location   Location `json:"location"`
		Altitude   float64  `json:"altitude"`
		Content    string   `json:"content"`
		Opened     int      `json:"opened"`
		Downloaded int      `json:"downloaded"`
	}

	BoxRepositoryI interface {
		BaseRepositoryI
		FindOneAndIncrement(result *Box, id string) error
		FindNearBoxes(result *[]Box, lat, lng float64) error
		FindTopBoxes(*[]Box) error
	}

	BoxServiceI interface {
		Create(req BoxCreateRequest, tmpFilename, filename string) (*Box, error)
		FindNearAndTopBoxes(req BoxListRequest) (*BoxListResponse, error)
		Retrieve(id string) (*Box, error)
		Download(id string) (string, string, error)
	}
)
