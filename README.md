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
| POST | [https://fp2-hacktiv8-production.up.railway.app/users/login] |
| PUT | [https://fp2-hacktiv8-production.up.railway.app/users] |
| DELETE | [https://fp2-hacktiv8-production.up.railway.app/users] |

> Note: Untuk method PUT dan DELETE diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login

#

### Photos
  Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Photos

 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-hacktiv8-production.up.railway.app/photos] |
| GET | [https://fp2-hacktiv8-production.up.railway.app/photos]|
| PUT | [https://fp2-hacktiv8-production.up.railway.app/photos/:photoId] |
| DELETE | [https://fp2-hacktiv8-production.up.railway.app/photos/:photoId]|
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
> Note: Seluruh method diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login.Untuk methode PUT dan DELETE hanya bisa dilakukan oleh user yang membuat comment dan diperlukan parameter Id comment pada URL 

#

 ### SocialMedias
  Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel SocialMedias

 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-hacktiv8-production.up.railway.app/socialmedias] |
| GET | [https://fp2-hacktiv8-production.up.railway.app/socialmedias]|
| PUT | [https://fp2-hacktiv8-production.up.railway.app/socialmedias/:socialMediaId] |
| DELETE | [https://fp2-hacktiv8-production.up.railway.app/socialmedias/:socialMediaId] |

> Note: Seluruh method diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login.Untuk methode PUT dan DELETE diperlukan parameter Id social media pada URL, dan hanya bisa dilakukan oleh user yang membuat social media
