package handlers

//
//type TagHandler struct {
//	signUp *service.TagService
//}
//
//func (h *TagHandler) Tag(c echo.Context) error {
//	email := c.FormValue("email")
//	firstName := c.FormValue("first")
//	lastName := c.FormValue("last")
//	password := c.FormValue("password")
//
//	if tag, err := h.signUp.CreateTag(email, firstName, lastName, password); err != nil {
//		return err
//	}
//
//	return sendMessage(c, http.StatusCreated, http.StatusText(http.StatusCreated))
//}
