package dto

type EmpNewAndOldPwDTO struct {
	EmpId uint64	`json:"empId"`
	NewPassword string		`json:"newPassword"` 
	OldPassword string		`json:"oldPassword"`	
}