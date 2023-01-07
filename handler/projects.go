package handler

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"main/model/db"
	"net/http"
	"strconv"
)

// Getprojects godoc
// @Summary Get one's projects
// @Description Get all the projects from user(login required)
// @Produce json
// @Success 200 {object} db.ProposalInfo
// @Router /api/v1/myproject [get]
func Getprojects(r *gin.Context) {
	id := r.GetInt("userID")
	data, _ := model.GetProposals(id)
	r.JSON(200, data)
}

// Getproject godoc
// @Summary Get a project
// @Description Get a project with its id
// @Param id query string true "the id of the project"
// @Produce json
// @Success 200 {object} db.ProposalInfo
// @Failure 404 {object} handler.Response
// @Router /api/v1/project [get]
func Getproject(r *gin.Context) {
	q := r.Query("InfoID")
	id, err := strconv.Atoi(q)
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusNotFound)
	}
	data := db.ProposalInfo{
		InfoID: int32(id),
	}
	data = model.GetSth(data)
	SendResponse(r, nil, data)
}

// UpdateProject godoc
// @Summary Update one's project
// @Description Update user's project(login required)
// @Param id query string true "the id of the project"
// @Param token header string true "token"
// @Produce json
// @Success 200 {object} json
// @Router /api/v1/project [post]
func UpdateProject(r *gin.Context) {
	sid := r.Query("id")
	if sid == "" {
		SendBadRequest(r, model.ErrBadRequest, nil, model.ErrorSender())
		return
	}
	id, _ := strconv.Atoi(sid)
	uid := r.GetInt("userID")
	data := db.ProposalInfo{
		InfoID: int32(id),
	}
	data = model.GetSth(data)
	if data.UID != int32(uid) {
		SendBadRequest(r, model.ErrNotAuthrized, nil, model.ErrorSender())
		return
	}
	r.ShouldBindJSON(&data)
	model.UpdateSth(data)
	SendResponse(r, nil, model.NoResponse)
}

// GetTemplate godoc
// @Summary Get a templte
// @Description Get a template with its id
// @Param id query string true "the id of the template"
// @Produce json
// @Success 200 {object} db.Template
// @Failure 404 {object} handler.Response
// @Router /api/v1/project/template [get]
func GetTemplate(r *gin.Context) {
	id, err := strconv.Atoi(r.Query("id"))
	if err != nil || id == 0 {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
	}
	data := db.Template{Temid: int32(id)}
	data = model.GetSth(data)
	SendResponse(r, nil, data)
}
