package utils

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	"github.com/bwmarrin/snowflake"
)

// TODO: need to set at env for NewNode value
func GenerateSnowflakeID() uint64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	id := node.Generate().Int64()
	return uint64(id)
}

// Get MySQL error column name
func GetColumnNameFromError(err error) string {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		switch mysqlErr.Number {
		case 1062: // Unique constraint violation
			re := regexp.MustCompile(`for key '(.+?)\.(.+?)'`)
			matches := re.FindStringSubmatch(mysqlErr.Message)
			if len(matches) == 3 {
				return matches[2] // Extract column name
			}
		case 1048: // Column cannot be NULL
			re := regexp.MustCompile(`Column '(.+?)' cannot be null`)
			matches := re.FindStringSubmatch(mysqlErr.Message)
			if len(matches) == 2 {
				return matches[1] // Extract column name
			}
		}
	}
	return ""
}

func CheckRow(err error) {
    if errors.Is(err, sql.ErrNoRows) {
        fmt.Println("No rows found in the result set")
    } else if err != nil {
        fmt.Println("Database error:", err)
    }
}