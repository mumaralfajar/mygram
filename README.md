## About MyGram
MyGram is a web service where user can post photo, add comment on other user photos and add social media. This project has 4 group endpoint which is:
1. User :
    - Register [POST]
    - Login [POST]

2. Photo :
    - GetAll [GET]
    - GetOne [GET]
    - CreatePhoto [POST]
    - UpdatePhoto [PUT]
    - DeletePhoto [DELETE]

3. Comment :
    - GetAll [GET]
    - GetOne [GET]
    - CreateComment [POST]
    - UpdateComment [PUT]
    - DeleteComment [DELETE]

4. Social Media :
    - GetAll [GET]
    - GetOne [GET]
    - CreateSocialMedia [POST]
    - UpdateSocialMedia [PUT]
    - DeleteSocialMedia [DELETE]

## Package or Library Used
- Go Framework: **Gin**
- ORM: **GORM**
- Database: **PostgresSQL**
- Environment Variable: **Viper**
- API Documentation: **Swagger**