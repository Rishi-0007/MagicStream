package handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/rishi-0007/magicstream-backend/internal/services"
	"github.com/rishi-0007/magicstream-backend/internal/utils"
)

type MovieHandler struct { svc *services.MovieService }
func NewMovieHandler(s *services.MovieService) *MovieHandler { return &MovieHandler{svc: s} }

func (h *MovieHandler) List(c *gin.Context){
	limit := int64(50)
	if ls := c.Query("limit"); ls!="" { if v,err := strconv.ParseInt(ls,10,64); err==nil && v>0 && v<=200 { limit = v } }
	movies, err := h.svc.List(c, limit)
	if err != nil { utils.Error(c,http.StatusInternalServerError,"could not list movies"); return }
	utils.OK(c, gin.H{"data": movies})
}

func (h *MovieHandler) Get(c *gin.Context){
	id := c.Param("imdb_id")
	m, err := h.svc.ByIMDBID(c, id)
	if err != nil { utils.Error(c,http.StatusNotFound,"movie not found"); return }
	utils.OK(c, m)
}

func (h *MovieHandler) Search(c *gin.Context){
	q := c.Query("q"); if q=="" { utils.Error(c,http.StatusBadRequest,"missing query parameter q"); return }
	limit := int64(50)
	if ls := c.Query("limit"); ls!="" { if v,err := strconv.ParseInt(ls,10,64); err==nil && v>0 && v<=200 { limit = v } }
	movies, err := h.svc.Search(c, q, limit)
	if err != nil { utils.Error(c,http.StatusInternalServerError,"search failed"); return }
	utils.OK(c, gin.H{"data": movies})
}
