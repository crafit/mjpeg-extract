A simple golang program to extract a single frame from an mjpeg stream.
In my specific case from [IP Webcam](
https://play.google.com/store/apps/details?id=com.pas.webcam&hl=en)

```
go build

# Raspberry Pi
env GOOS=linux GOARCH=arm GOARM=5 go build
```

This program would have been a lot shorter if go's multipart parsing was
more robust.
