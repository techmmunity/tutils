<div align="center">

<img src="https://github.com/techmmunity/tutils/raw/master/resources/logo.gif" width="300" height="300">

# Techmmunity - Tutils

<br>
<br>

</div>

Package of utils.

We also have **zero external dependencies**, so everything that can cause trouble to your project is inside this repository! It makes easier to check and validate if the code match your expectations and is good enough for your project.

## :warning: WARNING :warning:

The package may not be updated so often, but it **not means that it has been abandoned**. We do the utils in a way to (almost) never have to touch it again, and we only update the package to add new features, what may don't need to happen so often.

You can safely use this package, even if it don't receive updates in a year, and if you find any bug, you can report it by creating a issue in the GitHub repository, or create a PR to fix it! We will always give it special attention.

# Docs

## Miscellaneous

| function | description                          |
| -------- | ------------------------------------ |
| `Chunk`  | Chunks an array                      |
| `Nest`   | Nest objects that have been unnested |
| `Unnest` | Plain objects and arrays             |

## `get*`

| function       | description                                                            |
| -------------- | ---------------------------------------------------------------------- |
| `Unique`       | Remove duplicated values of an array (only work with primitive values) |
| `AspectRatio`  | Gets the aspect ratio (like 16:9, 4:3)                                 |
| `GCD`          | Gets the Greatest Common Divisor (GCD)                                 |
| `HexColorLuma` | Return the color luma of a hex color                                   |

## `has*`

| function      | description                        |
| ------------- | ---------------------------------- |
| `hasHtmlTags` | Return true if value has html tags |
| `hasEnvVars`  | Verify if all env vars are defined |
| `hasUrl`      | Return true if value has a url     |

## `is*`

| function           | description                                       |
| ------------------ | ------------------------------------------------- |
| `isAlphanumeric`   | Return true if value is a number or a letter      |
| `isBetween`        | Return true if value is a number between 2 values |
| `isBrazilianPhone` | Return true if value is a brazilian phone         |
| `isCnpj`           | Return true if value is a CNPJ                    |
| `isCnpjWithMask`   | Return true if value is a CNPJ with mask          |
| `isCpf`            | Return true if value is a CPF                     |
| `isCpfWithMask`    | Return true if value is a CPF with mask           |
| `isDarkHexColor`   | Return true if value is a dark hex color          |
| `isDivisibleByTen` | Return true if value is divisible by ten          |
| `isEmail`          | Return true if value is a valid email             |
| `isEmptySlice`     | Checks if is an Slice and it's empty              |
| `isEven`           | Return true if value is a even number             |
| `isHexColor`       | Return true if value is a valid hex color         |
| `isIpv4`           | Return true if value is a valid ipv4              |
| `isIpv4WithMask`   | Return true if value is a valid ipv4 with mask    |
| `isLightHexColor`  | Return true if value is a light hex color         |
| `isOdd`            | Return true if value is a odd number              |
| `isUrl`            | Return true if value is a valid url               |
