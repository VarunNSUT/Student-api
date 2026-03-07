package student

import (
	// "go/types"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/VarunNSUT/Students-api/internal/types"
	response "github.com/VarunNSUT/Students-api/internal/utils"
	"github.com/go-playground/validator/v10"
)

func New() http.HandlerFunc {
	return func ( w http.ResponseWriter , r *http.Request){

		var student types.Student 
		// we use json to decode our data and here r is the response we have received 
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err , io.EOF){
			response.WriteJson(w,http.StatusBadRequest , fmt.Errorf("empty body"))
			return 
		}

		if err != nil{
			response.WriteJson( w , http.StatusBadRequest ,response.GeneralError(err) )
		}

		slog.Info("creating a student ")

		// now to get information in go , we first need to decode it in a struct and then get it 

		// w.Write([]byte("welcome to students API goddamnn"))
		// validate the request 
		// for this we have a very powerful go package , we might just import it and use it simply 
		if err:= validator.New().Struct(student) ; err != nil {
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w , http.StatusBadRequest , response.ValidationError(validateErrs))
		} // here we have an argument of slice inside the validation function so we typecasted the err to slice 

		response.WriteJson(w,http.StatusCreated , map[string]string{"success":"ok"})
	}
}