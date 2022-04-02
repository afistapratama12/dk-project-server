db_migrate:
	go run migration/migrate.go migrate_db
db_drop:
	go run migration/migrate.go drop_db
create_user_dl:
	go run migration/migrate.go create_user_dl
create_seed_data:
	go run migration/migrate.go create_seed_data
create_admin_data:
	go run migration/migrate.go create_admin_data
scenarion_ro_wd:
	go run migration/migrate.go scenarion_ro_wd