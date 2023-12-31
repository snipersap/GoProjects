# Bed and Breakfast training

Update 30.09.2023: Completed the training and the basics of bed and breakfast project. Now will clean up and move it to [Bookings & Reservations](https://github.com/snipersap/bookings) project, which will be created as a separate repo.  

This repo is the project *bed-and-breakfast* as described in the course  [Building Modern Web Applications with Go (Golang)](https://udemy.com/course/building-modern-web-applications-with-go/).

This is the `bed and breakfast` project built with go as backend.

## How to test the repo?

1. Run the files in `cmd\web` such as `main.go` and any other file using the command `go run .`.
2. Go to your web browser and put in the URL `http:\\localhost:8080`.

Run the tests with
`go test`

### Test Result

In case of confusion, check the commits.

## Expected output

### Web browser

#### Home page

About is a link to the about page.

```html
This is the Home page.
You can go to the About page.
```

#### About page reporting IP without registration

Home is a link to the Home page

```html
This is the About page.
You can go to the Home page.

Another paragraph

About me: Software Engineering Manager from Germany

Your address was not yet registered. Please visit the Home page first and then come back to me.
```

#### About page reporting IP after registration

Home is a link to the Home page

```html
This is the About page.
You can go to the Home page.

Another paragraph

About me: Software Engineering Manager from Germany

Your remote address is:[::1]:56140
```

### Console

```bash
$ go run cmd/web/*.go
2023/09/30 12:28:13 Starting web server on port: :8080
Page was hit!: /about
2023/09/30 12:31:54 Rendered About
Page was hit!: /
2023/09/30 12:32:06 Rendered Home
Page was hit!: /about
2023/09/30 12:32:13 Rendered About
```

#### Simulating error in reading template files

```bash
2023/09/26 21:53:30 Starting web server on port: :8080
2023/09/26 21:53:34 [templates\about.page.tmpl templates\home.page.tmpl]
2023/09/26 21:53:34 ERROR could not read layouts from disk "search query"=./templates/*.lyout.tmpl
2023/09/26 21:53:34 ERROR could not read layouts from disk
2023/09/26 21:53:34 ERROR could not load cache from disk!
```

#### Simulating error in reading layout files

```bash
2023/09/26 21:49:48 Starting web server on port: :8080
2023/09/26 21:49:53 ERROR could not load templates from disk "search query"=./templates/*.pge.tmpl
2023/09/26 21:49:53 ERROR could not load templates from disk
2023/09/26 21:49:53 ERROR could not load cache from disk!
```

## Report bugs or suggestions

Please use the *Issues* feature in github to raise one. There is no guarantee or promise to fix it, because, well, this repo is a project after all. However, all suggestions are welcome.

## Disclaimer

Please use the code here at your own risk. This is a sample project, and may not always provide the expected outcome or result.
