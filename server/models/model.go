package models

type User struct {
    ID           uint   `gorm:"primaryKey"`
    Username     string `gorm:"size:100;not null"`
    Email        string `gorm:"size:100;unique;not null"`
    Password     string `gorm:"size:255;not null"`
    RollID       uint   `gorm:"not null"`
    Role         Role   `gorm:"foreignKey:RollID"`
    RefreshToken string `json:"refresh_token"`
}

type Role struct {
    ID               uint         `gorm:"primaryKey"`
    EmployeePosition string       `gorm:"size:50;unique;not null"`
    Permissions      []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
    ID          uint   `gorm:"primaryKey"`
    Permissions string `gorm:"size:50;unique;not null"`
}

type RolePermission struct {
    ID           uint `gorm:"primaryKey"`
    RoleID       uint
    PermissionID uint
}