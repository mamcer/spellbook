# misc

## rip audio cd

Best quality, flac

```bash
sudo apt-get install abcde flac
```

how to rip

```bash
mkdir music
cd music
# insert audio cd
abcde -o flac
```

## cursor rules configuration

Cursor > Preferences > Cursor Settings

go-rules.mdc

```
You are an expert AI programming assistant specializing in building APIs with Go, using the standar library's net/http package and the new ServeMux introduced in Go 1.22.
Always use the latest stable version of Go (1.24 or newer) and be familiar with the RESTful API design principles, best practices, and Go idioms.
- Follow the user's requirements carefully and to the letter.
- First think step-by-step - describe yor plan for the API structure, endpoints, and data flow in pseudocode, written out in great detail.
- Confirm the plan, then write code!
- Write correct, up-to-date, bug-free, fully functional secure and efficient Go code for APIs.
- Use the standard library's net/http package for API development:
- Utilize the new ServeMux introduced in Go 1.22 for routing
- Implement proper handling of different HTTP methods (GET, POST, PUT, DELETE, etc.)
- Use method handlers with appropiate signatures (e.g.. func(w http.REsponseWriter, r *http.Request)) 
- Leverage new features like wildcard matching and regex support in routes
- Implement proper error handling, incluiding custom error types when beneficial.
- Use appropiate status codes and format JSON responses correctly.
- Implement input validation for API endpoints.
- Utilize Go's built-in concurrency features when benefical for API performance.  
```

## software best practices

### commits 

[commit messages guide](https://github.com/RomuloOliveira/commit-messages-guide/blob/master/README.md)  
[conventional commits](https://www.conventionalcommits.org/en/v1.0.0/)  
[git commit](https://chris.beams.io/posts/git-commit/)  

commit seven rules

1. Separate subject from body with a blank line
2. Limit the subject line to 50 characters
3. Capitalize the subject line
4. Do not end the subject line with a period
5. Use the imperative mood in the subject line
6. Wrap the body at 72 characters
7. Use the body to explain 'what' and 'why' vs. how

 > a commit message should continue the phrase 'If applied, this commit will'

### code reviews

[google code reviews](https://google.github.io/eng-practices/review/reviewer/)

### readme file

[make a readme](https://www.makeareadme.com)

### keep a changelog

[keep a changelog](https://keepachangelog.com/en/1.1.0/)

### hyrum's law

Represents the practical knowledge that—even with the best of intentions, the best engineers, and solid practices for code review—we cannot assume perfect adherence to published contracts or best practices. 

### andy and bill's law

Is a statement that new software will always consume any increase in computing power that new hardware can provide.  
[wikipedia](https://en.wikipedia.org/wiki/Andy_and_Bill%27s_law)

### wirth's law

Wirth's law is an adage on computer performance which states that software is getting slower more rapidly than hardware is becoming faster.  
[wikipedia](https://en.wikipedia.org/wiki/Wirth%27s_law)

## image processing

### install

```bash
sudo apt install imagemagick
```

### convert multiple heic files to jpg maintaining name

```bash
magick convert *.HEIC -set filename:base "%[basename]" "%[filename:base].jpg"
```

### resize image

```bash
convert -resize 50% IMG_1994.JPG IMG_1994.JPG
```

## tool to send files safely 

[sendsafely](https://explore.sendsafely.com/)
