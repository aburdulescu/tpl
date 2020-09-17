# tpl
Generate file based on a json data file and a template file

## Install:

### Install go
Use whatever method works for you:
- on GNU/Linux, install the official package from your package repository

OR

- follow the instructions from the golang website [here](https://golang.org/doc/install)

### Install the tool
```
git clone https://github.com/aburdulescu/tpl.git
cd tpl
make install
```

## Usage
You can find the files used here in the `example` directory.

The syntax of the template file can be found [here](https://pkg.go.dev/text/template).

- template file(`example.tpl`)

```
My name is {{.name}}.
My friends are:
{{range .friends}}
- {{.}}
{{end}}
The answer to the "Ultimate Question to Life, the Universe, and Everything" is {{.answer}}.
```

- data file(`example.json`)

```
{
    "name": "Arthur Dent",
    "friends": [
        "Ford Prefect",
        "Zaphod Beeblebrox",
        "Marvin",
        "Trillian"
    ],
    "answer": 42
}
```

- run the tool using the files above

```
tpl -t example.tpl -d example.json
```

- output

```
My name is Arthur Dent.
My friends are:

- Ford Prefect

- Zaphod Beeblebrox

- Marvin

- Trillian

The answer to the "Ultimate Question to Life, the Universe, and Everything" is 42.
```
