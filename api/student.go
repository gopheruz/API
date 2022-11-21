package api

import (
	"net/http"
	"strconv"
	"github.com/nurmuhammaddeveloper/API/models"

	"github.com/gin-gonic/gin"
)

// @Router /students/{id} [get]
// @Summary Get Student by id
// @Description Get student by id
// @Tags Student
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.Student
// @Failure 500 {object} models.ResponseError
func (hand *handler) Get(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"massege": err.Error()})
		return
	}
	student, err := hand.storage.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"massege": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, student)
}


// @Router /students [get]
// @Summary Get Students
// @Description Get Students
// @Tags Student
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param firstname query string false "FirstName"
// @Param lastname query string false "LastName"
// @Param username query string false "UserName"
// @Success 200 {object} models.GetStudentResult
// @Failure 500 {object} models.ResponseError
func (hand *handler) GetAll(ctx *gin.Context) {
	queyrparams, err := validateQueryParams(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"massege": err.Error()})
		return
	}
	response, err := hand.storage.GetAll(queyrparams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"massege": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)

}

func validateQueryParams(ctx *gin.Context) (*models.GetStudentsQueryParam, error) {
	var (
		limit int64 = 10
		page  int64 = 1
		err   error
	)

	if ctx.Query("limit") != "" {
		limit, err = strconv.ParseInt(ctx.Query("limit"), 10, 64)
		if err != nil {
			return nil, err
		}
	}

	if ctx.Query("page") != "" {
		page, err = strconv.ParseInt(ctx.Query("page"), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return &models.GetStudentsQueryParam{
		Limit:     limit,
		Page:      page,
		FirstName: ctx.Query("firstname"),
		LastName:  ctx.Query("lastname"),
		UserName:  ctx.Query("username"),
	}, nil
}



// @Router /students [post]
// @Summary Create a student
// @Description Create a student
// @Tags Student
// @Accept json
// @Produce json
// @Param blog body models.CreateStudentRequest true "Student"
// @Success 200 {object} models.Student
// @Failure 500 {object} models.ResponseError
func (hand *handler) Create(ctx *gin.Context){
	var student models.CreateStudentRequest
    err := ctx.ShouldBind(&student)
    if err!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"massege": err.Error()})
        return
    }
	studentresult, err := hand.storage.Create(&student)
    if err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"massege": err.Error()})
        return
	}
	ctx.JSON(http.StatusOK, studentresult)
}


// @Router /students/{id} [delete]
// @Summary Delete a Student
// @Description Delete a Student
// @Tags Student
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ResponseError
func (hand *handler) Delete(ctx *gin.Context){
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
    if err!= nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"massege": err.Error()})
        return
    }
	err = hand.storage.Delete(id)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"massege": err.Error()})
        return
	}
	ctx.JSON(http.StatusOK, gin.H{"massege": "Successfully deleted"})
}


// @Router /students/{id} [put]
// @Summary Update a blog
// @Description Update a blog
// @Tags Student
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param blog body models.CreateStudentRequest true "Student"
// @Success 200 {object} models.Student
// @Failure 500 {object} models.ResponseError
func (hand *handler) Udate(ctx *gin.Context){
	var student models.UpdateStudentRequest
	err := ctx.ShouldBind(&student)
    if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	student.ID  = id
	updated, err := hand.storage.Update(&student)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, updated)
}