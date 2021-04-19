# Crypto server

listening localhost:8020 

## API

### POST /api/script

Request 
```json
{ 
  "address" : "tb1q0g5y6yhk7z25n7zxl4vtcanxyp44ng74x8esta",
  "coin_type" : 9
}
``` 

```json
{ 
  "address" : "2MvgWG4eorkNVfVDH5aoqq3Ub4FiobsczfA",
  "coin_type" : 9
}
``` 

Response 
```json
{
      "result": {
          "script": "00147a284d12f6f09549f846fd58bc7666206b59a3d5"
      }
}
```

