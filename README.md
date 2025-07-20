## Minimal repo reproducing an issue with `data-text` binding in datastar

https://discord.com/channels/1296224603642925098/1396205754251350096

In this example we have a page which is streaming from `/forever`, every half second this will stream the following events with the follow effects:

event | expected | effect
------|--------|----------
patch-signals with new `counter` | `$counter` is updated in all bound elements/attrs/etc | all good!
patch-elements with `<div id="direct-id" data-text="$counter">` | element is updated and text content reflects `$counter` | all good!
patch-elements with `<div id="nested-id"><span data-text="$counter">` | element is updated and text content of nested element reflects `$counter` | nested element content is <b>not</b> the value of `$counter`

Once `/forever` goes around the loop and does `patch-signals` again, suddenly that last element is then updated (i.e. the binding works, but it's not being 'resolved' on update via stream).

Apologies for not using codepen, it turns out it's harder to simulate an SSE event than I thought.

### Run

```shell
go run main.go
```
