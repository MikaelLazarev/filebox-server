/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package controllers

import (
	"github.com/MikaelLazarev/filebox-server/config"
	"github.com/MikaelLazarev/filebox-server/core"
	"github.com/MikaelLazarev/filebox-server/errorhandler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type BoxController struct {
	service core.BoxServiceI
	tempDir string
}

func RegisterBoxController(config *config.Config, g *gin.Engine, ls core.BoxServiceI) {

	controller := BoxController{
		service: ls,
		tempDir: config.TemporaryDir,
	}

	r := g.Group("/api/boxes/") //, middlewares.JWTAuthMiddleware())
	r.GET("/", controller.ListByCoord)
	r.GET("/:id/", withId(controller.Retrieve))
	r.POST("/", controller.Upload)

}

// GET: /api/boxes/
// Returns array of boxes around user by his/her coordinate
func (bc *BoxController) ListByCoord(c *gin.Context) {
	result, err := bc.service.FindBoxesAround()
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET: /api/boxes/:id/
// Return Box info for particular id
func (bc *BoxController) Retrieve(c *gin.Context, id string) {
	result, err := bc.service.Retrieve(id)
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// POST: /api/boxes/
// Returns 201 if successfully created
func (bc *BoxController) Upload(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(file.Filename)

	// absFilename - absolute filename for temporary file
	absFilename :=  bc.tempDir + file.Filename

	// Defer removing file after putting it to IPFS
	defer os.Remove(absFilename)

	// Saving file on disk
	err = c.SaveUploadedFile(file, absFilename)
	if err != nil {
		log.Fatal(err)
	}
	//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	// Open file from disk and provide io.Reader to BoxService to create a file
	f, err := os.Open(absFilename)
	if err != nil {
		log.Fatal(err)
	}

	// Creating Box with file contents
	result, err := bc.service.Create(f, file.Filename)
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	// Return 201 of succeeded with newBox parameters
	c.JSON(http.StatusCreated, result)
}