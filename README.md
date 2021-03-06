# tpl
Generate file based on a json data file and a template file

## Install:

- Install golang(e.g. on Debian based distros: `sudo apt install golang`)
- then run
```
go get github.com/aburdulescu/tpl
```

## Usage

Run the tool with `-h` to see the available options.

### Examples

You can find the files used here in the `example` directory.

The syntax of the template file can be found [here](https://pkg.go.dev/text/template).

#### Data file provided as argument

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
tpl -t example.tpl example.json
```

```
My name is Arthur Dent.
My friends are:

- Ford Prefect

- Zaphod Beeblebrox

- Marvin

- Trillian

The answer to the "Ultimate Question to Life, the Universe, and Everything" is 42.
```

#### Data file from `stdin`

- template file(`example2.tpl`)

```
Name: {{.name}}
Race: {{.race}}
```

- data file(`example2.json`)

```
{
    "persons": [
        {
            "name": "Arthur Dent",
            "race": "human"
        },
        {
            "name": "Ford Prefect",
            "race": "alien"
        }
    ]
}
```

- run the tool and pass the data using [jq](https://stedolan.github.io/jq/)

```
jq ".persons[0]" | tpl -t example.tpl
```

```
Name: Arthur Dent
Race: human
```

```
jq ".persons[1]" | tpl -t example.tpl
```

```
Name: Ford Prefect
Race: alien
```
