# QRCode

This project generates 232x232 PNG QRCodes for any given data in the URL.

Usage:

```bash
PORT=1718 go run main.go
curl http://localhost:1718/my-data-here
```

## Deployment

If you use [Heroku](https://heroku.com), you can push to there and it will work out of the box.

## TODO

* net/http is enough or do we have better webservers?
* add specs
