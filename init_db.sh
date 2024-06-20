#!/bin/bash
# This script initializes the database using the setup_database.sql file.

mysql -h us-cluster-east-01.k8s.cleardb.net -u ba7762a19416ec -p'b30b2bf0' heroku_22e3eea5109bd2b < database/setup_database.sql
