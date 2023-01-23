# <span style="color:#C0BFEC">**🦔 Short Link Generator**</span>

## <span style="color:#C0BFEC">***Description:*** </span>

### <span>**This service implements an `API` for creating `short links` with the following parameters:**</span>

* One unique URL is linked to one short link
* Default length is 10 characters
* The short link contains Latin characters in the lower and upper
  case, numbers and symbol `_`

### <span>**Requests received by the server:**</span>
* `Post` which will keep the original URL
in the database and return the reduced one.  (`/getshort`)
* `Get` which gets the shortened URL
  and return the original. (`/getoriginal`)

### <span>**Data storage:**</span>

The storage uses a `PostgreSQL` database and `in-memory`, 
the code of which is located in the `cache_manager` folder

## <span style="color:#C0BFEC">***Enter to run:*** </span>

```shell
sudo docker-compose up
```
