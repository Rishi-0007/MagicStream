package utils
import("net/http";"github.com/gin-gonic/gin")
 type APIError struct{Message string `json:"message"`}
func JSON(c *gin.Context,s int,v any){c.JSON(s,v)}
func Error(c *gin.Context,s int,m string){c.AbortWithStatusJSON(s,APIError{Message:m})}
func OK(c *gin.Context,v any){JSON(c,http.StatusOK,v)}
