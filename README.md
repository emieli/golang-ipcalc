# Dev

## Wgo
> https://github.com/bokwoon95/wgo
Golang web server hot reload

```
go install github.com/bokwoon95/wgo@latest
wgo run cmd/main.go
```

## Templ
Compile templates into .go files
```
wget https://github.com/a-h/templ/releases/download/v0.3.833/templ_Linux_x86_64.tar.gz
tar -xzf templ_Linux_x86_64.tar.gz 

go get github.com/a-h/templ

./templ generate --watch
```
*Also, get templ-vscode extension*

## Tailwind
```
apt install watchman
wget https://github.com/tailwindlabs/tailwindcss/releases/download/v4.0.14/tailwindcss-linux-x64
mv tailwindcss-linux-x64 tailwindcss
chmod +x tailwindcss
./tailwindcss -i ./static/css/input.css -o ./static/css/style.css --watch
```
