/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package services

import (
	"github.com/MikaelLazarev/filebox-server/core"
	"github.com/MikaelLazarev/filebox-server/errorhandler"
	"io"
	"log"
)

type boxService struct {
	repository core.BoxRepositoryI
	ipfs       core.IPFSRepositoryI
}

func NewBoxService(repository core.BoxRepositoryI, ir core.IPFSRepositoryI) core.BoxServiceI {
	return &boxService{
		repository: repository,
		ipfs:       ir}
}

// Retrieves Box by ID
func (s *boxService) Retrieve(id string) (*core.Box, error) {
	var box core.Box
	if err := s.repository.FindOne(&box, id); err != nil {
		return nil, errorhandler.DBError(err, "Box not found")
	}
	return &box, nil
}

// Find boxes around
func (s *boxService) FindBoxesAround() ([]core.Box, error) {
	panic("implement me")
}

// Creates a new box and return it
func (s *boxService) Create(r io.Reader, name string) (*core.Box, error) {
	ipfsHash, err := s.ipfs.AddFile(r)

	if err!= nil {
		return nil, err
	}

	log.Println(ipfsHash)

	newBox := core.Box{
		IpfsID: ipfsHash,
		Name:   name,
		Lat:    0,
		Lng:    0,
	}

	if err := s.repository.Create(&newBox); err != nil {
		return nil, err
	}

	return &newBox, nil
}


