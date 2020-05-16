# ELMŰ-ÉMÁSZ

(Unofficial) [ELMŰ-ÉMÁSZ][1] API client for Go

Download it using

```go get github.com/prozsolt/elmuemasz```

## Usage

```go
package main

import (
  "time"

  "github.com/prozsolt/elmuemasz"
)

func main() {
  // Authentication
  user := elmuemasz.User{
    Username: "username",
    Password: "password",
  }
  ts, err := elmuemasz.NewTokenSource(user)
  if err != nil {
    // handle error
  }
  client := oauth2.NewClient(context.Background(), ts)
  srv := elmuemasz.NewService(client)

  // Download the last 3 months of Invoices
  filter := elmuemasz.SzamlakFilter{
    SzamlaKelteTol: time.Now().AddDate(0, -3, 0),
    SzamlaKelteIg:  time.Now(),
  }
  szamlak, err := srv.Szamlak(filter)
  if err != nil {
    // handle error
  }
  for _, szamla := range szamlak {
    err = srv.DownloadPDF(szamla, "destinationDir")
    if err != nil {
      // handle error
    }
  }

  // Report utility meter reading (currently untested)
  felhelyek, err := srv.Felhelyek()
  if err != nil {
    // handle error
  }
  merodiktalasok, err := srv.MeroDiktalasok(felhelyek[0])
  if err != nil {
    // handle error
  }
  payload := elmuemasz.MeroDiktalasPayloadFromMeroDiktalas(merodiktalasok[0], time.Now(), 1337)
  _, err := srv.MeroDiktalasPost(payload)
}
```

[1]: https://ker.elmuemasz.hu/usz(bD1odSZjPTIwMQ==)/ker/newco/index.html#/