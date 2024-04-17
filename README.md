# How to install the CLI tool

## Build the project

go build -o notes main.go

## Move to the bin folder

mv mycli /usr/local/bin/

### Now you can run the CLI by running note in you terminal

#### Example

```console
foo@bar:~$ note add -f Ideas -t "Project Ideas" -c"Create a CLI tool in GO"
```


#### Add new note

```console
foo@bar:~$ note add
```

```console
foo@bar:~$ note a
```

```console
foo@bar:~$ note make
```

```console
foo@bar:~$ note n
```

#### Params

##### Folder

```console
foo@bar:~$ note add -f #Specifies the folder to be added
```

```console
foo@bar:~$ note add -folder #Specifies the folder to be added
```

##### Title

```console
foo@bar:~$ note add -t #Specifies the title of the note
```

```console
foo@bar:~$ note add -title #Specifies the title of the note
```

##### Content

```console
foo@bar:~$ note add -c #Specifies the content of the note
```

```console
foo@bar:~$ note add -content #Specifies the content of the note
```
