### Helpful commands


Build image command:
```
docker build -t tzmodifier .
```

List images:
```
docker images
```

Remove image by name (or image_id):
```
docker rmi tzmodifier
```

Run a container (interactive mode):
```
docker run -it --privileged \
  --tmpfs /run \
  -v /etc/localtime:/etc/localtime:rw \
  -v /usr/share/zoneinfo:/usr/share/zoneinfo \
  -v /etc/timezone:/etc/timezone \
  -v /var/run/dbus:/var/run/dbus \
  tzmodifier
```

List available timezones:
```
docker run --rm -it --privileged \
-v /etc/localtime:/etc/localtime \
-v /usr/share/zoneinfo:/usr/share/zoneinfo \
-v /etc/timezone:/etc/timezone \
tzmodifier list
```

Set a new timezone:
```
docker run --rm -it --privileged \
-v /etc/localtime:/etc/localtime \
-v /usr/share/zoneinfo:/usr/share/zoneinfo \
-v /etc/timezone:/etc/timezone \
tzmodifier set America/New_York
```

Check a current timezone:
```
docker run --rm -it --privileged \
-v /etc/localtime:/etc/localtime \
-v /usr/share/zoneinfo:/usr/share/zoneinfo \
-v /etc/timezone:/etc/timezone \
tzmodifier current
```