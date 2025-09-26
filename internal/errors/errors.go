package errors

type AppError struct {
    Code    int
    Message string
}

var (
    ErrInvalidCredentials = AppError{Code: 401, Message: "Credenciais inválidas"}
    ErrUserExists        = AppError{Code: 409, Message: "Usuário já existe"}
    ErrInvalidTOTP      = AppError{Code: 401, Message: "Código TOTP inválido"}
    ErrDatabaseError    = AppError{Code: 500, Message: "Erro no banco de dados"}
) 