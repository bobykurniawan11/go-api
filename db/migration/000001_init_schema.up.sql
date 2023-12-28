-- type User struct {
-- 	//ID UUID
-- 	ID uuid.UUID `gorm:"size:36;"`
-- 	//Name string
-- 	Name string `gorm:"size:255"`
-- 	//Email string
-- 	Email string `gorm:"size:255,unique"`
-- 	//Password string
-- 	Password string `gorm:"size:255"`

-- 	//Avatar string
-- 	Avatar string `gorm:"size:255"`

-- 	//CreatedAt time.Time
-- 	CreatedAt int64 `gorm:"autoCreateTime"`
-- 	//UpdatedAt time.Time
-- 	UpdatedAt int64 `gorm:"autoUpdateTime"`
-- }

create table users (
    id varchar(36) not null,
    name varchar(255) not null,
    email varchar(255) not null unique primary key,
    password varchar(255) not null,
    avatar varchar(255) not null,
    created_at timestamp not null,
    updated_at timestamp not null
);