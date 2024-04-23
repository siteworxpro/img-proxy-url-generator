# img-proxy url generator

img proxy [docs](https://docs.imgproxy.net/usage/processing)

## build

you can optionally hard code a path prefix on build
```shell
go build --ldflags="-X 'github.com/siteworxpro/img-proxy-url-generator/generator.PathPrefix=s3://mybucket'"
```

## config file

###### params 

[img-proxy]
- `key` sha256 hmac key
- `salt` sha256 hmac salt
- `host` img-proxy server hostname
- `encryption-key` aes encryption key
- `plain-url` send plain filename. bypasses all encoding

```ini
[img-proxy]
key=2c...47
salt=27...27
host=https://i.fooo.com
encryption-key=1c...0b
plain-url=1
```

## usage

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

generate an encrypted url
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

generate a url with params
```bash
./imgproxy --image "local:///my-super-image.png" -p h:200 -p rot:90

https://i.fooo.com/q-CfgLiuHTXDiZg7vBsUbZB3nkhzfsPgNrK0x20b878/h:200/rot:90/sm:1/enc/DrSKPtr8JkWx_Bf-vuxDTXRXfhrkZKTlPoQE61BzMfG2Mj1mD0qnthPq_Sfk8giv
```


generate a url a format
```bash
./imgproxy --image "local:///my-super-image.png" --format bmp -p h:200

https://i.fooo.com/UMkz4OUNw6P9ShLdewuvW3ValMgCt263vZzU5gN57WQ/h:200/sm:1/enc/ECYxMeVBTjRxB7F-jdQ7W_-Fnv4YbmSJIKie-Hdtxd9vsmEKjU1YuWVSzdN97Mod.bmp
```