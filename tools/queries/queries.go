package queries

const (
	GET_ALL_LEVEL = `SELECT * FROM tb_level WHERE level_status = 1`
	GET_BY_ID_LEVEL = `SELECT * FROM tb_level WHERE level_id = ? AND level_status = 1`
	CREATE_LEVEL = `INSERT INTO tb_level VALUES (NULL, ?, 1)`
	UPDATE_LEVEL = `UPDATE tb_level
					SET level_name = ?
					WHERE level_id = ? AND level_status = 1`
	DELETE_LEVEL = `UPDATE tb_level
					SET level_status = 0
					WHERE level_id = ?`


)
