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
)
