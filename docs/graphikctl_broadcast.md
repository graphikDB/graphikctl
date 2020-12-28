## graphikctl broadcast

graphikDB broadcast operations

```
graphikctl broadcast [flags]
```

### Examples

```
graphikctl broadcast --channel testing --data '{"text": "testing!"}'
```

### Options

```
      --channel string   the channel to publish a message to
      --data string      json attributes of the message to send
  -h, --help             help for broadcast
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.graphikctl.yaml)
```

### SEE ALSO

* [graphikctl](graphikctl.md)	 - 

###### Auto generated by spf13/cobra on 28-Dec-2020