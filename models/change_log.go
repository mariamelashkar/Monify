package models

import (
	"monify/utils"
	"database/sql"
	"time"
)

type ChangeLog struct {
	Id              int       `orm:"auto" json:"id"`
	TablesName      string    `orm:"size(255);column(table_name)" json:"table"`
	RowId           string    `orm:"column(row_id)" json:"row_id"`
	OldData         string    `orm:"type(jsonb);null" json:"old_data"`
	NewData         string    `orm:"type(jsonb);null" json:"new_data"`
	ChangeType      string    `orm:"size(255);column(change_type)" json:"change_type"`
	ChangeTimestamp time.Time `orm:"auto_now_add;type(timestamp);column(change_timestamp)" json:"change_timestamp"`
	ActionUser      string      `orm:"column(action_user)"`
}

func (c *ChangeLog) TableName() string {
	return "change_log"
}

func InitializeChangeLogSystem() error {
	db, err := utils.GetDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := createLogChangesFunction(db); err != nil {
		return err
	}

	if err := createTriggersForAllTables(db); err != nil {
		return err
	}

	return nil
}

func createLogChangesFunction(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE OR REPLACE FUNCTION log_changes() RETURNS TRIGGER AS $$
        BEGIN
            IF (TG_OP = 'INSERT') THEN
                INSERT INTO change_log (table_name, row_id, new_data, change_type, action_user, change_timestamp)
                VALUES (TG_TABLE_NAME, NEW.id, row_to_json(NEW), 'create', NEW.created_by, current_timestamp);
                RETURN NEW;
            ELSIF (TG_OP = 'UPDATE') THEN
                INSERT INTO change_log (table_name, row_id, old_data, new_data, change_type, action_user, change_timestamp)
                VALUES (TG_TABLE_NAME, NEW.id, row_to_json(OLD), row_to_json(NEW), NEW.action, NEW.updated_by, current_timestamp);
                RETURN NEW;
            ELSIF (TG_OP = 'DELETE') THEN
                INSERT INTO change_log (table_name, row_id, old_data, change_type, action_user, change_timestamp)
                VALUES (TG_TABLE_NAME, OLD.id, row_to_json(OLD), 'DELETE', 0, current_timestamp);
                RETURN OLD;
            END IF;
            RETURN NULL;
        END;
        $$ LANGUAGE plpgsql;
    `)
	return err
}

func createTriggersForAllTables(db *sql.DB) error {
	_, err := db.Exec(`
        DO $$
        DECLARE
            table_rec RECORD;
        BEGIN
            FOR table_rec IN
                SELECT table_name
                FROM information_schema.tables
                WHERE table_schema = 'public'
                AND table_type = 'BASE TABLE'
                AND table_name NOT IN ('change_log') 
            LOOP
                EXECUTE format('
                    CREATE OR REPLACE TRIGGER trg_%I_log_changes
                    AFTER INSERT OR UPDATE OR DELETE ON %I
                    FOR EACH ROW EXECUTE FUNCTION log_changes();', table_rec.table_name, table_rec.table_name);
            END LOOP;
        END $$;
    `)
	return err
}
