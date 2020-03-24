## ion-cli session

Manage a session within ION

### Synopsis

Allow to create, restore or delete a session file further calls would read the configs from and populate with needed data for other calls:

```
ion-cli session [flags]
```

### Options

```
  -d, --delete   Delete the current session
  -h, --help     help for session
```

### Options inherited from parent commands

```
  -c, --config string     Configs file path (default "./config/config-test.json")
      --profiles string   Profiles file path (default "./config/profiles-test.json")
  -s, --session string    Session file path (default "./config/session-test.json")
  -v, --verbose           verbose output
```

### SEE ALSO

* [ion-cli](ion-cli.md)	 - Cross-chain framework tool
* [ion-cli session init](ion-cli_session_init.md)	 - Add a session within ION

###### Auto generated by spf13/cobra on 24-Mar-2020