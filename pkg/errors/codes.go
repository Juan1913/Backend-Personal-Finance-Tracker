package errors

const (
	// Genéricos
	CodeInternalServerError   = "GEN-500"
	StatusInternalServerError = "INTERNAL_SERVER_ERROR"
	MsgInternalServerError    = "Error interno del servidor"

	// Usuarios
	CodeUserBadRequest = "USR-400"
	StatusBadRequest   = "BAD_REQUEST"
	MsgUserBadRequest  = "Error al parsear el usuario"
	UserNotFound       = "USER_NOT_FOUND"

	CodeUserNotFound = "USR-404"
	StatusNotFound   = "NOT_FOUND"
	MsgUserNotFound  = "Usuario no encontrado"

	CodeUserCreateError = "USR-500"
	MsgUserCreateError  = "Error al crear usuario"
	MsgUserGetError     = "Error al obtener usuarios"
	MsgUserUpdateError  = "Error al actualizar usuario"
	MsgUserDeleteError  = "Error al eliminar usuario"

	CodeUserEmailExists = "USR-409"
	StatusConflict      = "CONFLICT"
	MsgUserEmailExists  = "El email ya está registrado"

	// Categorías
	CodeCategoryBadRequest   = "CAT-400"
	MsgCategoryBadRequest    = "Error al parsear la categoría"
	CodeCategoryUserNotFound = "CAT-404"
	MsgCategoryUserNotFound  = "Usuario para la categoría no encontrado"
	CodeCategoryCreateError  = "CAT-500"
	MsgCategoryCreateError   = "Error al crear la categoría"
)
