package db

import (
	"backend/utils"
	"database/sql"
	"model"
	"time"

	_ "github.com/lib/pq"
)

func OpenDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=123 dbname=postgres sslmode=disable")
	utils.ThrowPanic(err)

	err = db.Ping()

	return db, err
}

func insertUserConfig(userId string) {
	conn, err := OpenDBConnection()
	utils.ThrowPanic(err)
	var id string

	sql := `INSERT INTO user_configuration(Theme,Font_size,User_id) VALUES($1,$2,$3) RETURNING User_id`
	err = conn.QueryRow(sql, "light", 14, userId).Scan(&id)
	utils.ThrowPanic(err)

	defer conn.Close()
}

func InsertUser(user model.UserInsertPayload) {
	conn, err := OpenDBConnection()
	utils.ThrowPanic(err)
	var id string
	timeNow := time.Now().Format(time.RFC3339)

	sql := `INSERT INTO users("user",created_at,updated_at) VALUES($1,$2,$3) RETURNING id`
	err = conn.QueryRow(sql, user.User, timeNow, timeNow).Scan(&id)
	utils.ThrowPanic(err)

	insertUserConfig(id)

	defer conn.Close()
}

func GetUserById(id string) model.UserResponse {
	conn, err := OpenDBConnection()
	utils.ThrowPanic(err)

	userRow := conn.QueryRow(`SELECT * FROM users WHERE id=$1`, id)

	var userId, username, created_at, updated_at, followers, following string
	err = userRow.Scan(&userId, &username, &created_at, &updated_at)
	utils.ThrowPanic(err)

	followersQuery := conn.QueryRow(`SELECT COUNT(*) FROM followers WHERE user_refer=$1`, id)
	followingQuery := conn.QueryRow(`SELECT COUNT(*) FROM followers WHERE user_id=$1`, id)
	utils.ThrowPanic(followersQuery.Scan(&followers))
	utils.ThrowPanic(followingQuery.Scan(&following))

	var user = model.UserResponse{
		Id:              userId,
		User:            username,
		Created_at:      created_at,
		Updated_at:      updated_at,
		FollowersLength: followers,
		FollowingLength: following,
	}

	defer conn.Close()
	return user
}

func GetUserConfigurationById(id string) model.UserConfiguration {
	conn, err := OpenDBConnection()
	utils.ThrowPanic(err)

	configurationRow := conn.QueryRow(`SELECT * FROM user_configuration WHERE user_id=$1`, id)

	var theme, user_id string
	var font_size int
	err = configurationRow.Scan(&user_id, &font_size, &theme)
	utils.ThrowPanic(err)

	var configuration = model.UserConfiguration{
		User_id:   user_id,
		Font_size: font_size,
		Theme:     theme,
	}

	defer conn.Close()

	return configuration
}

func UpdateUser(user model.UserUpdatePayload, userId string) {
	conn, err := OpenDBConnection()
	utils.ThrowPanic(err)
	var id string

	sql := `UPDATE users SET "user"=$1 WHERE id=$2 RETURNING id`
	err = conn.QueryRow(sql, user.User, userId).Scan(&id)
	utils.ThrowPanic(err)

	defer conn.Close()
}

func UpdateUserConfiguration(configuration model.UserConfigurationPayload, userId string) {
	conn, err := OpenDBConnection()
	utils.ThrowPanic(err)
	var id string

	sql := `UPDATE user_configuration SET theme=$1, font_size=$2 WHERE user_id=$3 RETURNING user_id`
	err = conn.QueryRow(sql, configuration.Theme, configuration.Font_size, userId).Scan(&id)
	utils.ThrowPanic(err)

	defer conn.Close()
}
