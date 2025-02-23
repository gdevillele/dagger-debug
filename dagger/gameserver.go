package main

// import "dagger/dagger-debug/internal/dagger"

// const (
// 	gameserverBuildEnvImage = "voxowl/ubuntu:24.04"
// )

// func (m *DaggerDebug) GameserverBuildEnv(source *dagger.Directory) *dagger.Container {
// 	// create a Dagger cache volume for dependencies
// 	// gsCache := dag.CacheVolume("gameserver")
// 	container := dag.Container().
// 		// start from a base Node.js container
// 		From("voxowl/ubuntu:24.04").
// 		// add the source code at /src
// 		WithDirectory("/src", source).
// 		// mount the cache volume at /root/.npm
// 		// WithMountedCache("/root/.npm", gsCache).
// 		// change the working directory to /src
// 		WithWorkdir("/src").
// 		// run npm install to install dependencies
// 		WithExec([]string{"npm", "install"})

// 	return container
// }

// # deps (sources)
// COPY ./cubzh/deps/bgfx /project/cubzh/deps/bgfx
// COPY ./cubzh/deps/lpng/src /project/cubzh/deps/lpng/src

// # deps (prebuilt)
// COPY ./cubzh/deps/libluau/0.661/prebuilt/linux/x86_64 /project/cubzh/deps/libluau/0.661/prebuilt/linux/x86_64
// COPY ./cubzh/deps/libssl/linux/amd64 /project/cubzh/deps/libssl/linux/amd64
// COPY ./cubzh/deps/libwebsockets/linux/amd64 /project/cubzh/deps/libwebsockets/linux/amd64
// COPY ./cubzh/deps/libz/linux-x86_64 /project/cubzh/deps/libz/linux-x86_64
// COPY ./cubzh/deps/xptools /project/cubzh/deps/xptools
