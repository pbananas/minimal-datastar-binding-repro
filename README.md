## Minimal repo reproducing an issue with `data-text` binding in datastar

https://discord.com/channels/1296224603642925098/1396205754251350096

In this example we have a page which is streaming from `/forever`, every half second this will stream the following events with the follow effects, with another half second between each:

event | expected | effect
------|--------|----------
patch-signals with new `counter` | `$counter` is updated in all bound elements/attrs/etc | all good!
patch-elements with `<div data-text="$counter">` | element is updated and text content reflects `$counter` | all good!
patch-elements with `<div><span data-text="$counter">` | element is updated and text content of nested element reflects `$counter` | nested element content is <b>not</b> the value of `$counter`

![](screenshot1.png)

Once `/forever` goes around the loop and does `patch-signals` again, suddenly that last element is then updated (i.e. the binding works, but it's not being 'resolved' on update via stream).

![](screenshot2.png)

Events look like this:

```
event: datastar-patch-signals
data: signals {"counter":3}


event: datastar-patch-elements
data: elements <div id="target-element-direct" class="large-number" data-text="$counter"></div>


event: datastar-patch-elements
data: elements <div id="target-element-nested"><span data-text="$counter" class="large-number">$counter NOT RESOLVED</span></div>


event: datastar-patch-signals
data: signals {"counter":4}


event: datastar-patch-elements
data: elements <div id="target-element-direct" class="large-number" data-text="$counter"></div>


event: datastar-patch-elements
data: elements <div id="target-element-nested"><span data-text="$counter" class="large-number">$counter NOT RESOLVED</span></div>
```

So each `patch-signals` causes `$counter NOT RESOLVED` to be replaced by the value of `$counter`, but then every time we send down a new `target-element-nested` it reverts back to the placeholder.

Apologies for not using codepen, it turns out it's harder to simulate an SSE event than I thought.

### Run

```shell
go run main.go
```
