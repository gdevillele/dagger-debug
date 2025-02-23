# Dagger Debug

This is a simple environment to try & debug Dagger things.

## dagger.json `include` and `exclude`

I don't understand how the `include` and `exclude` fields in `dagger.json` work.
My `debug` function prints out the name of all the files, no matter what I put in the `include` and `exclude` fields.

### Command to run:

```sh
dagger call debug -v --src=.
```

### Output:

```
% dagger call debug -v --src=.
✔ connect 0.4s
│ ✔ starting engine 0.2s
│ │ ✔ create 0.2s
│ │ │ ✔ exec docker ps -a --no-trunc --filter name=^/dagger-engine- --format {{.Names}} 0.1s
│ │ │ ┃ dagger-engine-v0.16.1
│ │ │ ✔ exec docker start dagger-engine-v0.16.1 0.1s
│ │ │ ┃ dagger-engine-v0.16.1
│
│ ✔ connecting to engine 0.1s
│ ✔ starting session 0.1s

✔ load module 1.5s
│ ✔ finding module configuration 0.1s
│ ✔ initializing module 1.3s
│ ✔ inspecting module metadata 0.1s
│ ✔ loading type definitions 0.1s

✔ parsing command line arguments 0.0s

✔ Host.directory(path: "/Users/gaetan/projects/gdevillele/dagger-debug"): Directory! 0.0s
│ ✔ upload /Users/gaetan/projects/gdevillele/dagger-debug from 9zh1g468cw8de612xqc6fhzzi (client id: e3ae6g15t00m8boxak0251mir, session id: k00nh2zwe9gt456fasxzaruh6) 0.0s

✔ daggerDebug: DaggerDebug! 0.0s
✔ .debug(
│ │ src: no(digest: "sha256:5e3df71e0188e1132e4cadf625171463da61b4924c66847d54eefcfc5a8ef869"): Missing
│ ): Void 0.1s
┃ 🔨 Building gameserver...
┃ ⭐️ Contents of src directory:
┃ - .DS_Store
┃ - .git
┃ - LICENSE
┃ - README.md
┃ - dagger
┃ - dagger.json
│ ✔ Missing.entries: [String!]! 0.0s

Setup tracing at https://dagger.cloud/traces/setup. To hide set DAGGER_NO_NAG=1
```

