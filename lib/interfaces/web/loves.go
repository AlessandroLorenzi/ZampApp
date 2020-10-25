package web

import (
	"net/http"
	"zampapp/lib/entity/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (s *Service) newLove(c *gin.Context) {
	userID, found := c.Get(identityKey)
	if !found {
		c.JSON(http.StatusForbidden, gin.H{"msg": "login required"})
		return
	}
	animalID := c.Param("animal_id")

	l := model.Love{
		PersonID: userID.(string),
		AnimalID: animalID,
	}

	err := s.repoService.CreateLove(l)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"err":    err.Error(),
			"userId": userID,
		}).Error("error AnimalIsLovedBy")
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "error fetching data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "ok"})

}

func (s *Service) listAnimalsLoved(c *gin.Context) {
	userID, found := c.Get(identityKey)
	if !found {
		c.JSON(http.StatusForbidden, gin.H{"msg": "login required"})
		return
	}
	aa, err := s.repoService.AnimalsLovedBy(userID.(string))
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"err":    err.Error(),
			"userId": userID,
		}).Error("error AnimalIsLovedBy")
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "error fetching data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "ok", "content": gin.H{
		"animals": aa,
	}})

}
