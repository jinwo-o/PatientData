mysqldump -u root -p patients > patients_backup.sql
mysql -u root -proot1234 < drop_database.sql
mysql -u root -proot1234 < create_database.sql
mysql -u root -proot1234 < create_table.sql
