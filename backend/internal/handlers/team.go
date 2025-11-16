package handlers

import (
	"avito/consts"
	"avito/internal/dto"
	apperrors "avito/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	@Summary      Create team
//  @Description  create team method dont need verify
//	@Tags         Team
//	@Accept       json
//	@Produce      json
//	@Param        request    body     dto.TeamReq  true  "request body for create team"
//	@Success      201  {array}   dto.TeamResp
//	@Failure      400  {object}  dto.ErrorResponse
//	@Failure      404  {object}  dto.ErrorResponse
//	@Failure      500  {object}  dto.ErrorResponse
//	@Router       /team/add [post]
func (h *Handler) CreateTeam(c *gin.Context) {
	var req dto.TeamReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code: consts.ErrBadRequest,	
			Message: err.Error(),
		})
		return
	}

	err = h.srv.CreateTeam(&req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	users, err := h.srv.GetUsersByTeam(req.TeamName)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.ErrorResponse{
			Code:    consts.ErrNotFound,
			Message: apperrors.ErrTeamNotFound.Error(),
		})
	}

	c.JSON(http.StatusCreated, dto.TeamResp{
		TeamName: req.TeamName,
		Members:  users,
	})
}

//	@Summary      Get team
//	@Description  Get team by name
//  @Description  For use this method you need to be user/admin
//	@Tags         Team
//	@Accept       json
//	@Produce      json
//	@Param        team_name    query     string  true  "query param for get team"
//	@Success      200  {array}   dto.TeamResp
//	@Failure      400  {object}  dto.ErrorResponse
//	@Failure      404  {object}  dto.ErrorResponse
//	@Failure      500  {object}  dto.ErrorResponse
//	@Router       /team/get [get]
//	@Security     BearerAuth
func (h *Handler) GetTeam(c *gin.Context) {
	teamName := c.Query("team_name")
	if teamName == "" {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    consts.ErrBadRequest,
			Message: apperrors.ErrNoTeamQuery.Error(),
		})
		return
	}

	users, err := h.srv.GetTeamByName(teamName)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, dto.TeamResp{
		TeamName: teamName,
		Members:  users,
	})
}
