# img-proxy url generator


[host](https://docs.imgproxy.net/installation) an instance of img-proxy

img proxy [docs](https://docs.imgproxy.net/usage/processing)

## build

you can optionally hard code a path prefix when building
```shell
go build --ldflags="-X 'github.com/siteworxpro/img-proxy-url-generator/generator.PathPrefix=s3://mybucket'"
```

## config file

### params 

[img-proxy]
- `key` sha256 hmac key
- `salt` sha256 hmac salt
- `host` img-proxy server hostname
- `encryption-key` aes encryption key
- `plain-url` bypasses all encoding and uses plain filename

example config file
```ini
[img-proxy]
key=2c...47
salt=27...27
host=https://i.fooo.com
encryption-key=1c...0b
plain-url=1
```

## usage examples

generate a plain url with an insecure signature
```ini
[img-proxy]
host=https://i.fooo.com
plain-url=1
```

```bash
./imgproxy --image "local:///my-super-image.png"

https://i.fooo.com/insecure/raw:1/plain/local:///my-super-image.png
```

generate a plain url with a signature
```ini
[img-proxy]
host=https://i.fooo.com
key=23...a4
salt=3f...4a
plain-url=1
```

```bash
./imgproxy --image "local:///my-super-image.png"

https://i.fooo.com/g0ociV0BWaK8JPXtkNsGxdK8IFHbVYczrcmiiv5T9pk/raw:1/plain/local://my-super-image.png
```

generate a hmac url with a signature
```ini
[img-proxy]
host=https://i.fooo.com
key=23...a4
salt=3f...4a
```

```bash
./imgproxy --image "local:///my-super-image.png"

https://i.fooo.com/NIColt6GBjtbquJXAtEMMsptARPw0CdeduEcu-S9Voc/raw:1/bG9jYWw6Ly9teS1zdXBlci1pbWFnZS5wbmc
```

generate an encrypted url with a signature
```ini
[img-proxy]
host=https://i.fooo.com
key=23...a4
salt=3f...4a
encryption-key=1c...0b
```

```bash
./imgproxy --image "local:///my-super-image.png"

https://i.fooo.com/m3YtaMSgL86qCnfKCnS2i9_vLRmJSogdBx1o86cWbuc/raw:1/enc/F6FAWktv2SAFe5UQwMme0pB6JwKQJVtTI_6Xx-PUfKANdQk0pD1I13NPnv0CvkFT
```

generate with params
```bash
./imgproxy --image "local:///my-super-image.png" -p h:200 -p rot:90

https://i.fooo.com/q-CfgLiuHTXDiZg7vBsUbZB3nkhzfsPgNrK0x20b878/h:200/rot:90/sm:1/enc/DrSKPtr8JkWx_Bf-vuxDTXRXfhrkZKTlPoQE61BzMfG2Mj1mD0qnthPq_Sfk8giv
```


generate a url with a specified format
```bash
./imgproxy --image "local:///my-super-image.png" --format bmp -p h:200

https://i.fooo.com/UMkz4OUNw6P9ShLdewuvW3ValMgCt263vZzU5gN57WQ/h:200/sm:1/enc/ECYxMeVBTjRxB7F-jdQ7W_-Fnv4YbmSJIKie-Hdtxd9vsmEKjU1YuWVSzdN97Mod.bmp
```
## decryption

if you need to decrypt a url you have already created just copy the encrypted portion of the url

```shell
./imgproxy decrypt -u ECYxMeVBTjRxB7F-jdQ7W_-Fnv4YbmSJIKie-Hdtxd9vsmEKjU1YuWVSzdN97Mod
```

## web service
you can also serve request via a web request

```shell
./imgproxy server
```

```shell
curl --location 'http://localhost:8080/generate' \
--header 'Content-Type: application/json' \
--data '{
    "image": "s3://my.image.bucket/C81A0923.jpg",
    "params": [
        "q:40",
        "w:400"
    ],
    "format": "bmp"
}'
```
`https://i.fooo.com/UMkz4OUNw6P9ShLdewuvW3ValMgCt263vZzU5gN57WQ/h:200/sm:1/enc/ECYxMeVBTjRxB7F-jdQ7W_-Fnv4YbmSJIKie-Hdtxd9vsmEKjU1YuWVSzdN97Mod.bmp`