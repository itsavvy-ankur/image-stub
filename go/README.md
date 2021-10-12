# A Golang based image stub to service images over a rest endpoint

* Uses a yaml configuration to specify loading images locally from disk or over http

```yaml
- requestId : "1234"
  imageUrl: "C:/images/tree.jpg"
  itemBarcode: "9876"

- requestId : "123456"
  imageUrl: "https://images.unsplash.com/photo-1453728013993-6d66e9c9123a?ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8dmlld3xlbnwwfHwwfHw%3D&ixlib=rb-1.2.1&w=1000&q=80"
  itemBarcode: "9876"
```

## dgdf
