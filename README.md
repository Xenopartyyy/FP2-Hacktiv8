# FINAL PROJECT 2 MyGram

[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)


Berikut ini adalah final project ke-2 dari hacktiv8, aplikasi ini bernama bernama MyGram, yang dimana pada aplikasi ini  dapat menyimpan foto maupun membuat comment untuk foto orang lain. Aplikasi ini akan dilengkapi dengan proses CRUD.

## Endpoint
Berikut ini adalah seluruh endpoint yang dapat diakses melalui client

### Users
 
 Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Users
 
 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-hacktiv8-production.up.railway.app/users/register][PlDb] |
| POST | [https://fp2-hacktiv8-production.up.railway.app/users/login][PlGh] |
| PUT | [https://fp2-hacktiv8-production.up.railway.app/users][PlGd] |
| DELETE | [https://fp2-hacktiv8-production.up.railway.app/users][PlOd] |

> Note: Untuk method PUT dan DELETE diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login



### Photos
  Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Photos

 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-hacktiv8-production.up.railway.app/photos][PlDb] |
| GET | [https://fp2-hacktiv8-production.up.railway.app/photos][PlGh] |
| PUT | [https://fp2-hacktiv8-production.up.railway.app/photos/:photoId][PlGd] |
| DELETE | [https://fp2-hacktiv8-production.up.railway.app/photos/:photoId][PlOd] |
> Note: Seluruh method diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login.Untuk methode PUT dan DELETE hanya bisa dilakukan oleh user yang membuat photo dan diperlukan parameter Id photo pada URL
#
#

### Comments
  Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Comments

 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-hacktiv8-production.up.railway.app/comments][PlDb] |
| GET | [https://fp2-hacktiv8-production.up.railway.app/comments][PlGh] |
| PUT | [https://fp2-hacktiv8-production.up.railway.app/comments/:commentId][PlGd] |
| DELETE | [https://fp2-hacktiv8-production.up.railway.app/comments/:commentId][PlOd] |
> Note: Seluruh method diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login.Untuk methode PUT dan DELETE hanya bisa dilakukan oleh user yang membuat comment dan diperlukan parameter Id comment pada URL 
#
#

 ### SocialMedias
  Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel SocialMedias

 | Method | URL |
| ------ | ------ |
| POST | [https://fp2-hacktiv8-production.up.railway.app/socialmedias][PlDb] |
| GET | [https://fp2-hacktiv8-production.up.railway.app/socialmedias][PlGh] |
| PUT | [https://fp2-hacktiv8-production.up.railway.app/socialmedias/:socialMediaId][PlGd] |
| DELETE | [https://fp2-hacktiv8-production.up.railway.app/socialmedias/:socialMediaId][PlOd] |

> Note: Seluruh method diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login.Untuk methode PUT dan DELETE diperlukan parameter Id social media pada URL, dan hanya bisa dilakukan oleh user yang membuat social media
