# Residential College (RC) Forum
Residential College (RC) web forum is a centralized web-based platform to support communication, coordination and community interaction within RC. It merges announcements, issue reporting and day-to-day communication (Buy / Sell /Give and Open Jio) into a single web forum.

# Feacutres
- Announcement page for RA (Residential Assistant) to inform residents
- Report page for residents to submit issue reports
- Marketplace place for residents to buy / sell / give
- Open Jio page for residents to organize events


# Backend
- go 
- go-chi
- sqlc
- Goose
- PostgreSQL

# Updating my backend database
1. Connect to PostgreSQL
    ```psql -h localhost -U rc_user -d rc_forum```
2. List database
    ```\l``` 
3. Connect to database
    ```\c rc_forum```
4. Check current Goose migration
    ```SELECT * FROM goose_db_version;```
5. Remove applied migration (if needed)
    ```DELETE FROM goose_db_version WHERE version_id = <version_number>```
6. Verify current tables
    ```\dt```

# Running go 
```go run cmd/*.go```

# Adding new table
1. Create a new migration file
    ```goose -dir db/migrations -s create create_products sql```
2. Generate SQLC code
    ```sqlc generate```
3. Run the migrations
    ```goose up```

# Git Hub command
1. ```git add <file> ``` or ```git add .```
2. ```git commit -m "your commit message```
3. ```git push```
