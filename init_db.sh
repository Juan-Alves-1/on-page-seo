#!/bin/bash
# This script initializes the database using the setup_database.sql file.

mysql -h $DB_HOST -u $DB_USER -p$DB_PASS $DB_NAME < database/setup_database.sql
