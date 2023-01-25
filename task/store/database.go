package store

// // get the service name
// var tasksDatabase = sqlx.NewDb(sqldb.Named("task").Stdlib(), "postgres")

// // GetByField - get task by field
// func GetByField(ctx context.Context, field, ops string, value interface{}) (Task, error) {
// 	// set the data fields for the query
// 	data := map[string]interface{}{
// 		field: value,
// 	}

// 	// query statement to be executed
// 	q := "SELECT * FROM tasks WHERE %v %v :%v LIMIT 1"
// 	// format query parameters
// 	q = fmt.Sprintf(q, field, ops, field)

// 	// declare user
// 	var user Task
// 	// execute query
// 	if err := database.NamedStructQuery(ctx, tasksDatabase, q, data, &user); err != nil {
// 		if err == database.ErrNotFound {
// 			return Task{}, ErrNotFound
// 		}
// 		return Task{}, fmt.Errorf("selecting users by ID[%v]: %w", value, err)
// 	}

// 	return user, nil
// }
