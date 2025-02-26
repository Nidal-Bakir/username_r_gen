# username_r_gen

<img src="./gopher.jpg" />

## A pseudo random username generator.

|    Name    |                             Unique Count                             |
| :--------: | :------------------------------------------------------------------: |
| Adjectives |                                 325                                  |
|   Colors   |                                 148                                  |
|  Animals   |                                 535                                  |
|   Total    | 325 x 148 x 535 = 25,733,500 Possible combination (with out postfix) |

> [!NOTE]
> Using a type of Postfix like `UseProvidedNumber` or `UseProvidedNumberAfterOverflow` would paricaly make the conflicts nearly impossible

---

# How to use

```go
nameGenerator := usernaemgen.NewUsernameGen()
uniqueName := nameGenerator.Generate(1)
fmt.Println(uniqueName) // cute-blue-fox
```

### Or use the GenerateRand() fn

```go
nameGenerator := usernaemgen.NewUsernameGen()
uniqueName := nameGenerator.GenerateRand()
fmt.Println(uniqueName) // nervous-Gold-hornbill-551536555
```

### You can customize the dictionaries, delimiter, and postfixType

```go
nameGenerator := usernaemgen.NewUsernameGenWithOptions(
	"__",
	usernaemgen.NoPostfix,
	usernaemgen.Adjectives,
	usernaemgen.Animals,
)
uniqueName := nameGenerator.GenerateRand()
fmt.Println(uniqueName) // great__canary
```
