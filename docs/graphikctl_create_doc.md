## graphikctl create doc

create a document

```
graphikctl create doc [flags]
```

### Examples

```
graphikctl create doc --gtype task --attributes '{ "title": "this is a title", "description": "this is a description", "priority": "low"}'
```

### Options

```
      --attributes string   json attributes of the doc
      --gid string          the gid of the doc
      --gtype string        the gtype of the doc
  -h, --help                help for doc
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.graphikctl.yaml)
```

### SEE ALSO

* [graphikctl create](graphikctl_create.md)	 - graphikDB create operations

###### Auto generated by spf13/cobra on 27-Dec-2020