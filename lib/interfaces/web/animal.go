package web

import (
	"zampapp/lib/entity/model"

	"github.com/gin-gonic/gin"
)

func (s Service) getAnimal(c *gin.Context) {
	idAnimal := c.Param("id_animal")

	a, err := s.repoService.GetAnimal(idAnimal)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(400, gin.H{"msg": "animal not found"})
			return
		}
		c.JSON(500, gin.H{"msg": "unexpected error"})
		return
	}

	c.JSON(200, gin.H{"msg": "ok", "animal": a})
}

func (s Service) getAnimals(c *gin.Context) {
	aa, err := s.repoService.GetAnimals()
	if err != nil {
		c.JSON(500, gin.H{"msg": "unexpected error"})
		return
	}
	c.JSON(200, gin.H{"msg": "ok", "animals": aa})
}

func (s Service) newAnimal(c *gin.Context) {
	type newAnimalPost struct {
		Name          string         `json:"name" binding:"required"`
		Breed         string         `json:"breed" binding:"required"`
		Size          int            `json:"size" binding:"required"`
		Sex           bool           `json:"sex" binding:"required"`
		OwnerID       string         `json:"owner_id" binding:"required"`
		Picture       string         `json:"picture" binding:"required"`
		Wormed        bool           `json:"wormed" binding:"required"`
		ChildFriendly bool           `json:"child_friendly" binding:"required"`
		Position      model.Location `json:"position" binding:"required"`
		PositionDesc  string         `json:"position_desc" binding:"required"`
		Description   string         `json:"description" binding:"required"`
	}

	var ap newAnimalPost
	if err := c.ShouldBindJSON(&ap); err != nil {
		s.logger.Debug("Error login", err)
		c.JSON(400, gin.H{
			"msg": "not valid",
		})
		return
	}

	a, err := model.NewAnimal(
		ap.Name,
		ap.Breed,
		ap.Size,
		ap.Sex,
		ap.OwnerID,
		ap.Picture,
		ap.Wormed,
		ap.ChildFriendly,
		ap.Position,
		ap.PositionDesc,
		ap.Description,
	)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "not valid",
		})
		return
	}

	c.JSON(200, gin.H{
		"msg":    "ok",
		"animal": a,
	})
}
