

# Blotion
This package is a reverse API for Notion. Write your content in Notion, and use it to read your content.

> ⚠️️ Note: `Blotion` is under development, use it carefully.

## Table of Contents

- [Installation](#install)
- [Extract Table](#table)
- [Read Page](#page)

## ⚙️ Installation
Installation is done using the [`go get`](https://golang.org/cmd/go/#hdr-Add_dependencies_to_current_module_and_install_them) command:
```
go get github.com/fj-onathan/blotion
```

## 📁 Extract Table

Get all content registered on table and organize it on JSON retrieve.

```go
package main

import (
	blotion "github.com/fj-onathan/blotion"
)

func main() {
	// ID of the Page who contains table
	PageID := "7c0af1f3ab1c4926bd11128892a174fe"
	
	// Return table list on JSON format
	TableList := blotion.ExportTable(PageID)	
	fmt.Printf("Table list in JSON %v \n", TableList)
}

```

## 📄 Read Page
Under development.