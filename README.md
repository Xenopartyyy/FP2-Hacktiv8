# FINAL PROJECT 2 MyGram

[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)


Berikut ini adalah final project ke-2 dari hacktiv8, aplikasi ini bernama bernama MyGram, yang dimana pada aplikasi ini  dapat menyimpan foto maupun membuat comment untuk foto orang lain. Aplikasi ini akan dilengkapi dengan proses CRUD.

## Endpoint
Berikut ini adalah seluruh endpoint yang dapat diakses melalui client

### Users
 
 Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Users
 
 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-hacktiv8-production.up.railway.app/users/register] |
| POST | [https://fp2-hacktiv8-production.up.railway.app/users/login |
| PUT | [https://fp2-hacktiv8-production.up.railway.app/users] |
| DELETE | [https://fp2-hacktiv8-production.up.railway.app/users] |

###### Daftar request users

POST Register User
 ```sh
{
    "age":integer,
    "email": "string",
    "password": "string",
    "username": "string"
}
```
#
POST Login User
 ```sh
{
    "email": "string",
    "password": "string",
}
```
#

PUT User

-Bearer Token <br />
-Param userId
 ```sh
{
    "email": "string",
    "username": "string"
}
```
#
DELETE User

-Bearer Token

> Note: Untuk method PUT dan DELETE diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login
#


### Photos
  Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Photos

 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-hacktiv8-production.up.railway.app/photos] |
| GET | [https://fp2-hacktiv8-production.up.railway.app/photos] |
| PUT | [https://fp2-hacktiv8-production.up.railway.app/photos/:photoId] |
| DELETE | [https://fp2-hacktiv8-production.up.railway.app/photos/:photoId] |

###### Daftar request photos

POST Photo

-Bearer Token
 ```sh
{
    "title":"string",
    "caption": "string",
    "photo_url": "string"
}
```
#
GET Photo

-Bearer Token

#
PUT Photo

-Bearer Token  <br />
-Param photoId
 ```sh
{
    "title": "string",
    "caption": "string"
    "photo_url": "string"
}
```
#
DELETE 

-Bearer Token  <br />
-Param PhotoId
> Note: Seluruh method diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login.Untuk methode PUT dan DELETE hanya bisa dilakukan oleh user yang membuat photo dan diperlukan parameter Id photo pada URL
#


### Comments
  Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Comments

 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-hacktiv8-production.up.railway.app/comments] |
| GET | [https://fp2-hacktiv8-production.up.railway.app/comments] |
| PUT | [https://fp2-hacktiv8-production.up.railway.app/comments/:commentId] |
| DELETE | [https://fp2-hacktiv8-production.up.railway.app/comments/:commentId] |

###### Daftar request comments

POST Comment

-Bearer Token
 ```sh
{
    "message":"string",
    "photo_id": integer
}
```
#
GET Comment

-Bearer Token

#
PUT Comment

-Bearer Token  <br />
-Param commentId
 ```sh
{
    "message": "string"
}
```
#
DELETE 

-Bearer Token  <br />
-Param commentId

> Note: Seluruh method diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login.Untuk methode PUT dan DELETE hanya bisa dilakukan oleh user yang membuat comment dan diperlukan parameter Id comment pada URL 
#


 ### SocialMedias
  Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel SocialMedias

 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-hacktiv8-production.up.railway.app/socialmedias] |
| GET | [https://fp2-hacktiv8-production.up.railway.app/socialmedias] |
| PUT | [https://fp2-hacktiv8-production.up.railway.app/socialmedias/:socialMediaId] |
| DELETE | [https://fp2-hacktiv8-production.up.railway.app/socialmedias/:socialMediaId] |

###### Daftar request socialmedias

POST SocialMedia

-Bearer Token
 ```sh
{
    "name":"string",
    "social_media_url": "string"
}
```
#
GET SocialMedia

-Bearer Token

#
PUT SocialMedia

-Bearer Token  <br />
-Param socialMediaId
 ```sh
{
    "name": "string",
    "social_media_url": "string"
}
```
#
DELETE 

-Bearer Token  <br />
-Param socialMediaId


> Note: Seluruh method diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login.Untuk methode PUT dan DELETE diperlukan parameter Id social media pada URL, dan hanya bisa dilakukan oleh user yang membuat social media
