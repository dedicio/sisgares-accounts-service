package dto

type UserDto struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	LevelId   string `json:"level_id"`
	CompanyId string `json:"company_id"`
}

type UserResponseDto struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	LevelId string `json:"level_id"`
}

type LevelDto struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	CompanyId   string   `json:"company_id"`
}

type LevelResponseDto struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

type PermissionDto struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"company_id"`
}

type PermissionResponseDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AccountInputDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Company  string `json:"company"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDto struct {
	ID        string `json:"id"`
	CompanyID string `json:"company_id"`
	Token     string `json:"token"`
}
